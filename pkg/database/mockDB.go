package database

import (
	"time"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
	"github.com/google/uuid"
)

type MockRelationships struct {
}

func NewMockRelationships() RelationshipDB {
	log.Info("Connecting to mock database")
	return &MockRelationships{}
}

func (mp *MockRelationships) Connect() error {
	return nil
}

func (mp *MockRelationships) PingDB() error {
	return nil
}

func (mp *MockRelationships) CloseDB() {
	log.Info("Mocked DB connection closed")
}

func (mp *MockRelationships) GetFriendsListByUserID(userID string) (*data.Relationships, error) {
	friendsList := findFriendsListByUserID(userID)
	if len(friendsList) == 0 {
		return nil, data.ErrorRelationshipNotFound
	}
	return &friendsList, nil
}

func (mp *MockRelationships) GetInvitesListByUserID(userID string) (*data.Relationships, error) {
	invitesList := findInvitesListByUserID(userID)
	if len(invitesList) == 0 {
		return nil, data.ErrorRelationshipNotFound
	}
	return &invitesList, nil
}

func (mp *MockRelationships) UpdateRelationship(relationship *data.Relationship) error {
	index := findIndexByRelationshipID(relationship.ID)
	if index == -1 {
		return data.ErrorRelationshipNotFound
	}

	err := mp.validateRelationship(relationship)
	if err != nil {
		return err
	}

	relationshipList[index] = relationship
	return nil
}

func (mp *MockRelationships) AddRelationship(relationship *data.Relationship) error {
	err := mp.validateRelationship(relationship)
	if err == nil {
		relationship.ID = uuid.NewString()
		relationship.ConversationID, err = mp.getConversationID([]string{relationship.User1.UserID, relationship.User2.UserID})
		relationshipList = append(relationshipList, relationship)
	}
	return err
}

func (mp *MockRelationships) DeleteRelationship(id string) error {
	index := findIndexByRelationshipID(id)
	if index == -1 {
		return data.ErrorRelationshipNotFound
	}

	relationshipList = append(relationshipList[:index], relationshipList[index+1:]...)

	return nil
}

// Returns an array of friends in the database
// Returns -1 when no relationship is found
func findFriendsListByUserID(id string) data.Relationships {
	var friendsList data.Relationships
	for _ , relationship := range relationshipList {
		if relationship.User1.UserID == id && relationship.User1.RelationshipType == data.Friend {
			friendsList = append(friendsList, relationship)
		} else if relationship.User2.UserID == id && relationship.User2.RelationshipType == data.Friend{
			friendsList = append(friendsList, relationship)
		}
	}
	return friendsList
}

// Returns an array of friends invites in the database
// Returns -1 when no relationship is found
func findInvitesListByUserID(id string) data.Relationships {
	var invitesList data.Relationships
	for _ , relationship := range relationshipList {
		if relationship.User1.UserID == id && relationship.User1.RelationshipType == data.PendingIncoming {
			invitesList = append(invitesList, relationship)
		} else if relationship.User2.UserID == id && relationship.User2.RelationshipType == data.PendingIncoming{
			invitesList = append(invitesList, relationship)
		}
	}
	return invitesList
}

// Returns a relationship in the database
// Returns -1 when no relationship is found
func findIndexByRelationshipID(id string) int {
	for index, relationship := range relationshipList {
		if relationship.ID == id {
			return index
		}
	}
	return -1
}

func (mp *MockRelationships) validateRelationship(relationship *data.Relationship) error {
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

func (mp *MockRelationships) validateUserExist(userID string) bool {
	return true
}

// Returns an bool when a relationship with the two users is found
func (mp *MockRelationships) relationshipExist(id string, userID1 string, userID2 string) (bool, error) {
	for _ , relationship := range relationshipList {
		if relationship.ID != id &&
		(relationship.User1.UserID == userID1 || relationship.User1.UserID == userID2) &&
		(relationship.User2.UserID == userID1 || relationship.User2.UserID == userID2){
			return true, nil
		}
	}
	return false, nil
}

func (mp *MockRelationships) getConversationID(userID []string) (string, error) {
	return uuid.NewString(), nil
}

////////////////////////////////////////////////////////////////////////////////
/////////////////////////// Mocked database ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

var relationshipList = []*data.Relationship{
	{
		ID:          	"a2181017-5c53-422b-b6bc-036b27c04fc8",
		User1:        	data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.PendingOutgoing},
		User2: 			data.User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44", RelationshipType: data.PendingIncoming},
		ConversationID:	"a2181017-5c53-422b-b6bc-036b27c04fc8",
		CreatedOn:   	time.Now().UTC().String(),
		UpdatedOn:   	time.Now().UTC().String(),
	},
	{
		ID:          	"e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		User1:        	data.User{UserID: "c5825d3e-8a77-11eb-8dcd-0242ac130003", RelationshipType: data.PendingOutgoing},
		User2: 			data.User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44", RelationshipType: data.PendingIncoming},
		ConversationID:	"e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		CreatedOn:   	time.Now().UTC().String(),
		UpdatedOn:   	time.Now().UTC().String(),
	},
	{
		ID:          	"c5825d3e-8a77-11eb-8dcd-0242ac130003",
		User1:        	data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.Friend},
		User2: 			data.User{UserID: "f171ea04-8a77-11eb-8dcd-0242ac130003", RelationshipType: data.Friend},
		ConversationID:	"c5825d3e-8a77-11eb-8dcd-0242ac130003",
		CreatedOn:   	time.Now().UTC().String(),
		UpdatedOn:   	time.Now().UTC().String(),
	},
	{
		ID:          	"f171ea04-8a77-11eb-8dcd-0242ac130003",
		User1:        	data.User{UserID: "0af831ea-8a78-11eb-8dcd-0242ac130003", RelationshipType: data.Friend},
		User2: 			data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.Friend},
		ConversationID:	"f171ea04-8a77-11eb-8dcd-0242ac130003",
		CreatedOn:   	time.Now().UTC().String(),
		UpdatedOn:   	time.Now().UTC().String(),
	},
}
