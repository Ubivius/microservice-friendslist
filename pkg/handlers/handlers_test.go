package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
	"github.com/gorilla/mux"
)

// Move to util package in Sprint 9, should be a testing specific logger
func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "Tests", log.LstdFlags)
}

func TestDeleteNonExistantRelationship(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/relationships/4", nil)
	response := httptest.NewRecorder()

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
	}
	request = mux.SetURLVars(request, vars)

	relationshipHandler.Delete(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Relationship not found") {
		t.Error("Expected response : Relationship not found")
	}
}

func TestAddRelationship(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		User1: 			data.User{UserID: 1, RelationshipType: data.PendingIncoming},
		User2:       	data.User{UserID: 2, RelationshipType: data.PendingOutgoing},
		ConversationID: 12,
	}

	request := httptest.NewRequest(http.MethodPost, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())
	relationshipHandler.AddRelationship(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestUpdateRelationship(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		User1: 			data.User{UserID: 1, RelationshipType: data.Friend},
		User2:       	data.User{UserID: 2, RelationshipType: data.Friend},
		ConversationID: 12,
	}

	request := httptest.NewRequest(http.MethodPut, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())
	relationshipHandler.UpdateRelationships(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingRelationship(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/relationships/1", nil)
	response := httptest.NewRecorder()

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	relationshipHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}
