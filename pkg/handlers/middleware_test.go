package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
	"github.com/gorilla/mux"
)

func TestValidationMiddlewareWithValidBody(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		User1: 			data.User{UserID: 10, RelationshipType: data.PendingOutgoing},
		User2:       	data.User{UserID: 11, RelationshipType: data.PendingIncoming},
		ConversationID: 1,
	}
	bodyBytes, _ := json.Marshal(body)

	request := httptest.NewRequest(http.MethodPost, "/relationships", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())

	// Create a router for middleware because function attachment is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/relationships", relationshipHandler.AddRelationship)
	router.Use(relationshipHandler.MiddlewareRelationshipValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestValidationMiddlewareWithNoUserID(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		User1:       	data.User{RelationshipType: data.PendingIncoming},
		User2:       	data.User{UserID: 2, RelationshipType: data.PendingIncoming},
		ConversationID: 1,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Error("Body passing to test is not a valid json struct : ", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/relationships", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())

	// Create a router for middleware because linking is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/relationships", relationshipHandler.AddRelationship)
	router.Use(relationshipHandler.MiddlewareRelationshipValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Field validation for 'UserID' failed on the 'required' tag") {
		t.Error("Expected error on field validation for UserID but got : ", response.Body.String())
	}
}

func TestValidationMiddlewareWithNoUser1(t *testing.T) {
	// Creating request body
	body := &data.Relationship{
		User2:       	data.User{UserID: 2, RelationshipType: data.PendingIncoming},
		ConversationID: 1,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Error("Body passing to test is not a valid json struct : ", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/relationships", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	relationshipHandler := NewRelationshipsHandler(NewTestLogger())

	// Create a router for middleware because linking is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/relationships", relationshipHandler.AddRelationship)
	router.Use(relationshipHandler.MiddlewareRelationshipValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Field validation for 'UserID' failed on the 'required' tag") {
		t.Error("Expected error on field validation for UserID but got : ", response.Body.String())
	}
}
