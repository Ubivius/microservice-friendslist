package database

import (
	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// The interface that any kind of database must implement
type RelationshipDB interface {
	GetFriendsListByUserID(userID string) (*data.Relationships, error)
	GetInvitesListByUserID(userID string) (*data.Relationships, error)
	UpdateRelationship(relationship *data.Relationship) error
	AddRelationship(relationship *data.Relationship) error
	DeleteRelationship(id string) error
	validateRelationship(relationship *data.Relationship) error
	relationshipExist(id string, userID1 string, userID2 string) (bool, error)
	validateUserExist(userID string) bool
	getConversationID() string
	Connect() error
	CloseDB()
}
