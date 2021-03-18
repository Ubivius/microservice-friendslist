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

func TestGetExistingFriendsListByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/friends/1", nil)
	response := httptest.NewRecorder()

	productHandler := NewRelationshipsHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"userid": "1",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.GetFriendsListByUserID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"userid\":1") && !strings.Contains(response.Body.String(), "\"relationshiptype\":1") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingFriendsListByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/friends/2", nil)
	response := httptest.NewRecorder()

	productHandler := NewRelationshipsHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"userid": "2",
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
	request := httptest.NewRequest(http.MethodGet, "/invites/2", nil)
	response := httptest.NewRecorder()

	productHandler := NewRelationshipsHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"userid": "2",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.GetInvitesListByUserID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"userid\":2") && !strings.Contains(response.Body.String(), "\"relationshiptype\":3") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingInvitesListByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/invites/3", nil)
	response := httptest.NewRecorder()

	productHandler := NewRelationshipsHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"userid": "3",
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
		User1: 			data.User{UserID: 1, RelationshipType: data.PendingIncoming},
		User2:       	data.User{UserID: 3, RelationshipType: data.PendingOutgoing},
		ConversationID: 1,
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

func TestAddRelationshipThatAlreadyExists(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		User1: 			data.User{UserID: 1, RelationshipType: data.PendingIncoming},
		User2:       	data.User{UserID: 2, RelationshipType: data.PendingOutgoing},
		ConversationID: 1,
	}

	request := httptest.NewRequest(http.MethodPost, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())
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
		User1: 			data.User{UserID: 1, RelationshipType: data.PendingIncoming},
		User2:       	data.User{UserID: 1, RelationshipType: data.PendingOutgoing},
		ConversationID: 1,
	}

	request := httptest.NewRequest(http.MethodPost, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())
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
		ID: 			1,		
		User1: 			data.User{UserID: 1, RelationshipType: data.Friend},
		User2:       	data.User{UserID: 2, RelationshipType: data.Friend},
		ConversationID: 2,
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

func TestUpdateRelationshipWithSameUserID(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		ID: 			1,
		User1: 			data.User{UserID: 1, RelationshipType: data.Friend},
		User2:       	data.User{UserID: 1, RelationshipType: data.Friend},
		ConversationID: 2,
	}

	request := httptest.NewRequest(http.MethodPut, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())
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
		ID: 			3,
		User1: 			data.User{UserID: 1, RelationshipType: data.Friend},
		User2:       	data.User{UserID: 2, RelationshipType: data.Friend},
		ConversationID: 2,
	}

	request := httptest.NewRequest(http.MethodPut, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())
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
		ID: 			8,
		User1: 			data.User{UserID: 1, RelationshipType: data.Friend},
		User2:       	data.User{UserID: 2, RelationshipType: data.Friend},
		ConversationID: 2,
	}

	request := httptest.NewRequest(http.MethodPut, "/relationships", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyRelationship{}, body)
	request = request.WithContext(ctx)

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())
	relationshipHandler.UpdateRelationships(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Relationship not found") {
		t.Error("Expected response : Relationship not found")
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

func TestDeleteNonExistantRelationship(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/relationships/0", nil)
	response := httptest.NewRecorder()

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())

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
