package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// AddRelationship creates a new relationship from the received JSON
func (relationshipHandler *RelationshipsHandler) AddRelationship(responseWriter http.ResponseWriter, request *http.Request) {
	relationshipHandler.logger.Println("Handle POST Relationship")
	relationship := request.Context().Value(KeyRelationship{}).(*data.Relationship)

	err := data.AddRelationship(relationship)
	if err == data.ErrorSameUserID {
		relationshipHandler.logger.Println("[ERROR} users with same userID", err)
		http.Error(responseWriter, "Users with same userID", http.StatusNotFound)
		return
	}else if err == data.ErrorRelationshipExist {
		relationshipHandler.logger.Println("[ERROR} relationship already exist", err)
		http.Error(responseWriter, "Relationship already exist", http.StatusNotFound)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
