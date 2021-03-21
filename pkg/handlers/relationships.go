package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyRelationship is a key used for the Relationship object inside context
type KeyRelationship struct{}

// RelationshipsHandler contains the items common to all relationship handler functions
type RelationshipsHandler struct {
	logger *log.Logger
}

// NewRelationshipsHandler returns a pointer to a RelationshipsHandler with the logger passed as a parameter
func NewRelationshipsHandler(logger *log.Logger) *RelationshipsHandler {
	return &RelationshipsHandler{logger}
}

// getRelationshipID extracts the relationship ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getRelationshipID(request *http.Request) int {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}

// getUserID extracts the user ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getUserID(request *http.Request) int {
	vars := mux.Vars(request)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		panic(err)
	}
	return userID
}
