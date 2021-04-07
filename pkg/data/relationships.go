package data

import (
	"fmt"
)

// ErrorRelationshipNotFound : Relationship specific errors
var ErrorRelationshipNotFound = fmt.Errorf("Relationship not found")

// ErrorSameUserID : Invalid Relationship specific error
var ErrorSameUserID = fmt.Errorf("can't create a relationship with two users with the same userID")

// ErrorRelationshipExist : Invalid Relationship specific error
var ErrorRelationshipExist = fmt.Errorf("a relationship with these two users already exists")

// ErrorUserNotFound : User specific errors
var ErrorUserNotFound = fmt.Errorf("UserID doesn't exist")

// RelationshipType of a relationship
type RelationshipType string

// relationship type of a friend request
const (
	None            RelationshipType = "None"            // user has no intrinsic relationship
	Friend          RelationshipType = "Friend"          // user is a friend
	Blocked         RelationshipType = "Blocked"         // user is blocked
	PendingIncoming	RelationshipType = "PendingIncoming" // user has a pending incoming friend request to connected user
	PendingOutgoing	RelationshipType = "PendingOutgoing" // current user has a pending outgoing friend request to user
)

// Relationship defines the structure for an API relationship.
type Relationship struct {
	ID             string `json:"id" bson:"_id"`
	User1          User   `json:"user_1" bson:"user_1"`
	User2          User   `json:"user_2" bson:"user_2"`
	ConversationID string `json:"conversation_id" bson:"conversation_id"`
	CreatedOn      string `json:"-"`
	UpdatedOn      string `json:"-"`
}

// User in a relationship
type User struct {
	UserID      		string  	     	`json:"user_id" bson:"user_id" validate:"required"`
	RelationshipType	RelationshipType	`json:"relationship_type" bson:"relationship_type" validate:"required,isRelationshipType"`
}

// Relationships is a collection of Relationship
type Relationships []*Relationship

const MicroserviceUserPath = "http://microservice-user:9090"
const MicroserviceTextChatPath = "http://microservice-text-chat:9090"
