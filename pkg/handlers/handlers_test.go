package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
	"github.com/Ubivius/microservice-friendslist/pkg/database"
	"github.com/gorilla/mux"
)

func newRelationshipDB() database.RelationshipDB {
	return database.NewMockRelationships()
}

func TestGetExistingFriendsListByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/friends/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	productHandler := NewRelationshipsHandler(newRelationshipDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"user_id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.GetFriendsListByUserID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"user_id\":\"a2181017-5c53-422b-b6bc-036b27c04fc8\"") && !strings.Contains(response.Body.String(), "\"relationship_type\":\"Friend\"") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingFriendsListByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/friends/e2382ea2-b5fa-4506-aa9d-d338aa52af44", nil)
	response := httptest.NewRecorder()

	productHandler := NewRelationshipsHandler(newRelationshipDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"user_id": "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.GetFriendsListByUserID(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Friends not found") {
		t.Error("Expected response : Friends not found")
	}
}

func TestGetExistingInvitesListByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/invites/e2382ea2-b5fa-4506-aa9d-d338aa52af44", nil)
	response := httptest.NewRecorder()

	productHandler := NewRelationshipsHandler(newRelationshipDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"user_id": "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.GetInvitesListByUserID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"user_id\":\"e2382ea2-b5fa-4506-aa9d-d338aa52af44\"") && !strings.Contains(response.Body.String(), "\"relationship_type\":\"PendingIncoming\"") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingInvitesListByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/invites/c5825d3e-8a77-11eb-8dcd-0242ac130003", nil)
	response := httptest.NewRecorder()

	productHandler := NewRelationshipsHandler(newRelationshipDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"user_id": "c5825d3e-8a77-11eb-8dcd-0242ac130003",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.GetInvitesListByUserID(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Invites not found") {
		t.Error("Expected response : Invites not found")
	}
}

func TestAddRelationship(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		User1: 			data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.PendingIncoming},
		User2:       	data.User{UserID: "c5825d3e-8a77-11eb-8dcd-0242ac130003", RelationshipType: data.PendingOutgoing},
		ConversationID: "",
	}

	request := httptest.NewRequest(http.MethodPost, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(newRelationshipDB())
	relationshipHandler.AddRelationship(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestAddRelationshipThatAlreadyExists(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		User1: 			data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.PendingIncoming},
		User2:       	data.User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44", RelationshipType: data.PendingOutgoing},
		ConversationID: "",
	}

	request := httptest.NewRequest(http.MethodPost, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(newRelationshipDB())
	relationshipHandler.AddRelationship(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Relationship already exist") {
		t.Error("Expected response : Relationship already exist")
	}
}

func TestAddRelationshipWithSameUserID(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		User1: 			data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.PendingIncoming},
		User2:       	data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.PendingOutgoing},
		ConversationID: "",
	}

	request := httptest.NewRequest(http.MethodPost, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(newRelationshipDB())
	relationshipHandler.AddRelationship(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Users in the relationship with same userID") {
		t.Error("Expected response : Users in the relationship with same userID")
	}
}

func TestUpdateRelationship(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		ID: 			"a2181017-5c53-422b-b6bc-036b27c04fc8",		
		User1: 			data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.Friend},
		User2:       	data.User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44", RelationshipType: data.Friend},
		ConversationID: "",
	}

	request := httptest.NewRequest(http.MethodPut, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(newRelationshipDB())
	relationshipHandler.UpdateRelationships(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestUpdateRelationshipWithSameUserID(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		ID: 			"a2181017-5c53-422b-b6bc-036b27c04fc8",
		User1: 			data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.Friend},
		User2:       	data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.Friend},
		ConversationID: "",
	}

	request := httptest.NewRequest(http.MethodPut, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(newRelationshipDB())
	relationshipHandler.UpdateRelationships(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Users in the relationship with same userID") {
		t.Error("Expected response : Users in the relationship with same userID")
	}
}

func TestUpdateToARelationshipThatAlreadyExist(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		ID: 			"c5825d3e-8a77-11eb-8dcd-0242ac130003",
		User1: 			data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.Friend},
		User2:       	data.User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44", RelationshipType: data.Friend},
		ConversationID: "",
	}

	request := httptest.NewRequest(http.MethodPut, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(newRelationshipDB())
	relationshipHandler.UpdateRelationships(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Relationship already exist") {
		t.Error("Expected response : Relationship already exist")
	}
}

func TestUpdateNonExistantRelationship(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		ID: 			"",
		User1: 			data.User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: data.Friend},
		User2:       	data.User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44", RelationshipType: data.Friend},
		ConversationID: "",
	}

	request := httptest.NewRequest(http.MethodPut, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(newRelationshipDB())
	relationshipHandler.UpdateRelationships(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Relationship not found") {
		t.Error("Expected response : Relationship not found")
	}
}

func TestDeleteExistingRelationship(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/relationships/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	relationshipHandler := NewRelationshipsHandler(newRelationshipDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	relationshipHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteNonExistantRelationship(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/relationships/0", nil)
	response := httptest.NewRecorder()

	relationshipHandler := NewRelationshipsHandler(newRelationshipDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "0",
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
