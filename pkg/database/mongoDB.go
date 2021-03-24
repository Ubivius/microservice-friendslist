package database

import (
	"context"
	"os"
	"time"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRelationships struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoRelationships() RelationshipDB {
	mp := &MongoRelationships{}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		log.Error(err, "MongoDB setup failed")
		os.Exit(1)
	}
	return mp
}

func (mp *MongoRelationships) Connect() error {
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:pass@localhost:27888/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err, "Failed to ping database. Shutting down service")
		os.Exit(1)
	}

	log.Info("Connection to MongoDB established")

	collection := client.Database("test").Collection("relationships")

	// Assign client and collection to the MongoRelationships struct
	mp.collection = collection
	mp.client = client
	return nil
}

func (mp *MongoRelationships) CloseDB() {
	err := mp.client.Disconnect(context.TODO())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoRelationships) GetFriendsListByUserID(userID string) (*data.Relationships, error) {
	// MongoDB search filter
	filter := bson.D{{ 
		Key:"$or", 
		Value: bson.A{
			bson.D{{
				Key: "$and",
				Value: bson.A{ 
					bson.D{{Key: "user_1.user_id", Value: userID}},
					bson.D{{Key: "user_1.relationship_type", Value: data.Friend}},
				},
			}},
			bson.D{{
				Key: "$and",
				Value: bson.A{ 
					bson.D{{Key: "user_2.user_id", Value: userID}},
					bson.D{{Key: "user_2.relationship_type", Value: data.Friend}},
				},
			}},
		}, 
	}}

	// friends will hold the array of Relationships
	var friends data.Relationships

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error getting friends from database")
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
		var result data.Relationship
		err := cursor.Decode(&result)
		if err != nil {
			log.Error(err, "Error decoding friends from database")
		}
		friends = append(friends, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err, "Error in cursor after iteration")
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return &friends, err
}

func (mp *MongoRelationships) GetInvitesListByUserID(userID string) (*data.Relationships, error) {
	// MongoDB search filter
	filter := bson.D{{ 
		Key:"$or", 
		Value: bson.A{
			bson.D{{
				Key: "$and",
				Value: bson.A{ 
					bson.D{{Key: "user_1.user_id", Value: userID}},
					bson.D{{Key: "user_1.relationship_type", Value: data.PendingIncoming}},
				},
			}},
			bson.D{{
				Key: "$and",
				Value: bson.A{ 
					bson.D{{Key: "user_2.user_id", Value: userID}},
					bson.D{{Key: "user_2.relationship_type", Value: data.PendingIncoming}},
				},
			}},
		}, 
	}}

	// friends will hold the array of Relationships
	var invites data.Relationships

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error getting invites from database")
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
		var result data.Relationship
		err := cursor.Decode(&result)
		if err != nil {
			log.Error(err, "Error decoding invites from database")
		}
		invites = append(invites, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err, "Error in cursor after iteration")
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return &invites, err
}

func (mp *MongoRelationships) UpdateRelationship(relationship *data.Relationship) error {
	err := mp.validateRelationship(relationship)
	if err != nil {
		return err
	}

	// Set updated timestamp in relationship
	relationship.UpdatedOn = time.Now().UTC().String()

	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: relationship.ID}}

	// Update sets the matched relationships in the database to relationship
	update := bson.M{"$set": relationship}

	// Update a single item in the database with the values in update that match the filter
	_, err = mp.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error(err, "Error updating relationship")
	}

	return err
}

func (mp *MongoRelationships) AddRelationship(relationship *data.Relationship) error {
	err := mp.validateRelationship(relationship)
	if err != nil {
		return err
	}

	relationship.ID = uuid.NewString()
	relationship.ConversationID = mp.getConversationID()
	// Adding time information to new relationship
	relationship.CreatedOn = time.Now().UTC().String()
	relationship.UpdatedOn = time.Now().UTC().String()

	// Inserting the new relationship into the database
	insertResult, err := mp.collection.InsertOne(context.TODO(), relationship)
	if err != nil {
		return err
	}

	log.Info("Inserting relationship", "Inserted ID", insertResult.InsertedID)
	return nil
}

func (mp *MongoRelationships) DeleteRelationship(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error deleting relationship")
	}

	log.Info("Deleted documents in relationships collection", "delete_count", result.DeletedCount)
	return nil
}

func (mp *MongoRelationships) validateRelationship(relationship *data.Relationship) error {
	if !mp.validateUserExist(relationship.User1.UserID) || !mp.validateUserExist(relationship.User2.UserID){
		return data.ErrorUserNotFound
	}
	if relationship.User1.UserID == relationship.User2.UserID {
		return data.ErrorSameUserID
	}

	exist, err := mp.relationshipExist(relationship.ID, relationship.User1.UserID, relationship.User2.UserID)
	if err != nil {
		return err
	}

	if exist {
		return data.ErrorRelationshipExist
	}

	return nil
}

func (mp *MongoRelationships) validateUserExist(userID string) bool {
	// validation of the UserID with a call to microservice-user 
	return true
}

func (mp *MongoRelationships) relationshipExist(id string, userID1 string, userID2 string) (bool, error) {
	// MongoDB search filter
	filter := bson.D{
		{ 
			Key:"$or", 
			Value: bson.A{
				bson.D{{Key: "user_1.user_id", Value: userID1}},
				bson.D{{Key: "user_1.user_id", Value: userID2}},
			}, 
		},
		{
			Key: "$or", 
			Value: bson.A{
				bson.D{{Key: "user_2.user_id", Value: userID1}},
				bson.D{{Key: "user_2.user_id", Value: userID2}},
			},
		},
	}

	// Holds search result
	var result data.Relationship

	// Find a single matching item from the database
	err := mp.collection.FindOne(context.TODO(), filter).Decode(&result)

	if err == mongo.ErrNoDocuments || result.ID == id {
		return false, nil
	}
	return true, err
}

func (mp *MongoRelationships) getConversationID() string {
	// Call to the text-chat microservice to create a conversation and get the ID
	return ""
}
