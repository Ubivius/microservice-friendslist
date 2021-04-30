package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/database"
	"github.com/gorilla/mux"
)

// KeyRelationship is a key used for the Relationship object inside context
type KeyRelationship struct{}

// RelationshipsHandler contains the items common to all relationship handler functions
type RelationshipsHandler struct {
	db database.RelationshipDB
}

// NewRelationshipsHandler returns a pointer to a RelationshipsHandler with the logger passed as a parameter
func NewRelationshipsHandler(db database.RelationshipDB) *RelationshipsHandler {
	return &RelationshipsHandler{db}
}

// getRelationshipID extracts the relationship ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getRelationshipID(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["id"]
	
	return id
}

// getUserID extracts the user ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getUserID(request *http.Request) string {
	vars := mux.Vars(request)
	userID := vars["user_id"]

	return userID
}
