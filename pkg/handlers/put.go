package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// UpdateRelationships updates the relationship with the ID specified in the received JSON relationship
func (relationshipHandler *RelationshipsHandler) UpdateRelationships(responseWriter http.ResponseWriter, request *http.Request) {
	relationship := request.Context().Value(KeyRelationship{}).(*data.Relationship)
	relationshipHandler.logger.Println("Handle PUT relationship", relationship.ID)

	// Update relationship
	err := data.UpdateRelationship(relationship)
	if err == data.ErrorUserNotFound {
		relationshipHandler.logger.Println("[ERROR} a userID doesn't exist", err)
		http.Error(responseWriter, "A UserID doesn't exist", http.StatusBadRequest)
		return
	}else if err == data.ErrorSameUserID {
		relationshipHandler.logger.Println("[ERROR} users in the relationship with same userID", err)
		http.Error(responseWriter, "Users in the relationship with same userID", http.StatusBadRequest)
		return
	}else if err == data.ErrorRelationshipExist {
		relationshipHandler.logger.Println("[ERROR} relationship already exist", err)
		http.Error(responseWriter, "Relationship already exist", http.StatusBadRequest)
		return
	}else if err == data.ErrorRelationshipNotFound {
		relationshipHandler.logger.Println("[ERROR} relationship not found", err)
		http.Error(responseWriter, "Relationship not found", http.StatusNotFound)
		return
	}

	// Returns status, no content required
	responseWriter.WriteHeader(http.StatusNoContent)
}
