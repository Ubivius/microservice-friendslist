package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// AddRelationship creates a new relationship from the received JSON
func (relationshipHandler *RelationshipsHandler) AddRelationship(responseWriter http.ResponseWriter, request *http.Request) {
	relationshipHandler.logger.Println("Handle POST Relationship")
	relationship := request.Context().Value(KeyRelationship{}).(*data.Relationship)

	err := relationshipHandler.db.AddRelationship(relationship)
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
	default:
		relationshipHandler.logger.Println("[ERROR] adding relationship", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
