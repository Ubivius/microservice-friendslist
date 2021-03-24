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
	err := relationshipHandler.db.UpdateRelationship(relationship)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorUserNotFound:
		relationshipHandler.logger.Println("[ERROR} a userID doesn't exist", err)
		http.Error(responseWriter, "A UserID doesn't exist", http.StatusBadRequest)
		return
	case data.ErrorSameUserID:
		relationshipHandler.logger.Println("[ERROR} users in the relationship with same userID", err)
		http.Error(responseWriter, "Users in the relationship with same userID", http.StatusBadRequest)
		return
	case data.ErrorRelationshipExist:
		relationshipHandler.logger.Println("[ERROR} relationship already exist", err)
		http.Error(responseWriter, "Relationship already exist", http.StatusBadRequest)
		return
	case data.ErrorRelationshipNotFound:
		relationshipHandler.logger.Println("[ERROR} relationship not found", err)
		http.Error(responseWriter, "Relationship not found", http.StatusNotFound)
		return
	default:
		relationshipHandler.logger.Println("[ERROR] updating relationship", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
