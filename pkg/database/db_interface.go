package database

import (
	"context"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// The interface that any kind of database must implement
type RelationshipDB interface {
	GetFriendsListByUserID(ctx context.Context, userID string) (*data.DetailedRelationships, error)
	GetInvitesListByUserID(ctx context.Context, userID string) (*data.DetailedRelationships, error)
	UpdateRelationship(ctx context.Context, relationship *data.Relationship) error
	AddRelationship(ctx context.Context, relationship *data.Relationship) error
	DeleteRelationship(ctx context.Context, id string) error
	GetUserDetails(userID string, relations data.Relationships) (*data.DetailedRelationships, error)
	GetUserByID(userID string) (*data.DetailedUser, error)
	Connect() error
	PingDB() error
	CloseDB()
}
