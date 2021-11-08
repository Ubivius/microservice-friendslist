package database

import (
	"context"
	"os"
	"testing"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

func integrationTestSetup(t *testing.T) {
	t.Log("Test setup")

	if os.Getenv("DB_USERNAME") == "" {
		os.Setenv("DB_USERNAME", "admin")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "pass")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "27888")
	}
	if os.Getenv("DB_HOSTNAME") == "" {
		os.Setenv("DB_HOSTNAME", "localhost")
	}

	err := deleteAllRelationshipsFromMongoDB()
	if err != nil {
		t.Errorf("Failed to delete existing relationships from database during setup")
	}
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoRelationships()
	if mp == nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBAddRelationshipIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	relationship := &data.Relationship{
		User1:          data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.PendingOutgoing},
		User2:          data.User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af45", RelationshipType: data.PendingIncoming},
		ConversationID: "",
	}

	mp := NewMongoRelationships()
	err := mp.AddRelationship(context.Background(), relationship)
	if err != nil {
		t.Errorf("Failed to add relationship to database")
	}
	mp.CloseDB()
}

func TestMongoDBUpdateRelationshipIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	relationship := &data.Relationship{
		ID:             "eb9aff9f-8c4e-47c3-9f6d-bd9aac3d9f31",
		User1:          data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.Friend},
		User2:          data.User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44", RelationshipType: data.Friend},
		ConversationID: "",
	}

	mp := NewMongoRelationships()
	err := mp.UpdateRelationship(context.Background(), relationship)
	if err != nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBGetFriendsListByUserIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoRelationships()
	_, err := mp.GetFriendsListByUserID(context.Background(), "a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetInvitesListByUserIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoRelationships()
	_, err := mp.GetInvitesListByUserID(context.Background(), "e2382ea2-b5fa-4506-aa9d-d338aa52af84")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}
