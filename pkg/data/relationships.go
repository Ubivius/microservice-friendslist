package data

import (
	"fmt"
	"time"
)

// ErrorRelationshipNotFound : Relationship specific errors
var ErrorRelationshipNotFound = fmt.Errorf("Relationship not found")

// RelationshipType of a relationship
type RelationshipType int

// relationship type of a friend request
const (
    None RelationshipType = iota  // user has no intrinsic relationship
    Friend						  // user is a friend
    Blocked						  // user is blocked
	PendingIncoming				  // user has a pending incoming friend request to connected user
	PendingOutgoing				  // current user has a pending outgoing friend request to user
)

// Relationship defines the structure for an API relationship.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Relationship struct {
	ID          	int     `json:"id"`
	User1      		User  	`json:"user1"`
	User2 			User	`json:"user2"`
	ConversationID 	int 	`json:"conversationid"`
	CreatedOn   	string  `json:"-"`
	UpdatedOn   	string  `json:"-"`
	DeletedOn   	string  `json:"-"`
}

// User in a relationship
type User struct {
	UserID      		int  	     		`json:"userid" validate:"required,exist"`
	RelationshipType	RelationshipType	`json:"relationshiptype" validate:"required,isRelationshipType"`
}

// Relationships is a collection of Relationship
type Relationships []*Relationship

// GetFriendsListByUserID returns a list of friends for the given user id
func GetFriendsListByUserID(id int) ([]*int, error) {
	friendsList := findFriendsListByUserID(id)
	if len(friendsList) == 0 {
		return nil, ErrorRelationshipNotFound
	}
	return friendsList, nil
}

// GetInvitesListByUserID returns returns a list of friends invites for the given user id
func GetInvitesListByUserID(id int) ([]*int, error) {
	invitesList := findInvitesListByUserID(id)
	if len(invitesList) == 0 {
		return nil, ErrorRelationshipNotFound
	}
	return invitesList, nil
}

// UpdateRelationship updates the relationship specified in received JSON
func UpdateRelationship(relationship *Relationship) error {
	index := findIndexByRelationshipID(relationship.ID)
	if index == -1 {
		return ErrorRelationshipNotFound
	}
	relationshipList[index] = relationship
	return nil
}

// AddRelationship creates a new relationship
func AddRelationship(relationship *Relationship) {
	relationship.ID = getNextID()
	relationshipList = append(relationshipList, relationship)
}

// DeleteRelationship deletes the relationship with the given id
func DeleteRelationship(id int) error {
	index := findIndexByRelationshipID(id)
	if index == -1 {
		return ErrorRelationshipNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	relationshipList = append(relationshipList[:index], relationshipList[index+1:]...)

	return nil
}

// Returns an array of friends in the database
// Returns -1 when no relationship is found
func findFriendsListByUserID(id int) []*int {
	var friendsList []*int
	for _ , relationship := range relationshipList {
		if relationship.User1.UserID == id && relationship.User1.RelationshipType == Friend {
			friendsList = append(friendsList, &relationship.User2.UserID)
		}else if relationship.User2.UserID == id && relationship.User2.RelationshipType == Friend{
			friendsList = append(friendsList, &relationship.User1.UserID)
		}
	}
	return friendsList
}

// Returns an array of friends invites in the database
// Returns -1 when no relationship is found
func findInvitesListByUserID(id int) []*int {
	var invitesList []*int
	for _ , relationship := range relationshipList {
		if relationship.User1.UserID == id && relationship.User1.RelationshipType == PendingIncoming {
			invitesList = append(invitesList, &relationship.User2.UserID)
		}else if relationship.User2.UserID == id && relationship.User2.RelationshipType == PendingIncoming{
			invitesList = append(invitesList, &relationship.User1.UserID)
		}
	}
	return invitesList
}

// Returns a relationship in the database
// Returns -1 when no relationship is found
func findIndexByRelationshipID(id int) int {
	for index, relationship := range relationshipList {
		if relationship.ID == id {
			return index
		}
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////

// Finds the maximum index of our fake database and adds 1
func getNextID() int {
	lastRelationship := relationshipList[len(relationshipList)-1]
	return lastRelationship.ID + 1
}

// relationshipList is a hard coded list of relationships for this
// example data source. Should be replaced by database connection
var relationshipList = []*Relationship{
	{
		ID:          	1,
		User1:        	User{UserID: 1, RelationshipType: PendingOutgoing},
		User2: 			User{UserID: 2, RelationshipType: PendingIncoming},
		ConversationID:	1,
		CreatedOn:   	time.Now().UTC().String(),
		UpdatedOn:   	time.Now().UTC().String(),
	},
	{
		ID:          	2,
		User1:        	User{UserID: 3, RelationshipType: PendingOutgoing},
		User2: 			User{UserID: 2, RelationshipType: PendingIncoming},
		ConversationID:	2,
		CreatedOn:   	time.Now().UTC().String(),
		UpdatedOn:   	time.Now().UTC().String(),
	},
	{
		ID:          	3,
		User1:        	User{UserID: 1, RelationshipType: Friend},
		User2: 			User{UserID: 4, RelationshipType: Friend},
		ConversationID:	3,
		CreatedOn:   	time.Now().UTC().String(),
		UpdatedOn:   	time.Now().UTC().String(),
	},
	{
		ID:          	4,
		User1:        	User{UserID: 5, RelationshipType: Friend},
		User2: 			User{UserID: 1, RelationshipType: Friend},
		ConversationID:	4,
		CreatedOn:   	time.Now().UTC().String(),
		UpdatedOn:   	time.Now().UTC().String(),
	},
}
