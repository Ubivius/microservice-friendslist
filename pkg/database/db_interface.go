package database

import (
	"context"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// The interface that any kind of database must implement
type RelationshipDB interface {
	GetFriendsListByUserID(ctx context.Context, userID string) (*data.Relationships, error)
	GetInvitesListByUserID(ctx context.Context, userID string) (*data.Relationships, error)
	UpdateRelationship(ctx context.Context, relationship *data.Relationship) error
	AddRelationship(ctx context.Context, relationship *data.Relationship) error
	DeleteRelationship(ctx context.Context, id string) error
	validateRelationship(ctx context.Context, relationship *data.Relationship) error
	relationshipExist(ctx context.Context, id string, userID1 string, userID2 string) (bool, error)
	validateUserExist(ctx context.Context, userID string) bool
	getConversationID(ctx context.Context, userID []string) (string, error)
	Connect() error
	PingDB() error
	CloseDB()
}
