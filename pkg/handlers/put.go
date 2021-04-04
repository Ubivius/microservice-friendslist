package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// UpdateRelationships updates the relationship with the ID specified in the received JSON relationship
func (relationshipHandler *RelationshipsHandler) UpdateRelationships(responseWriter http.ResponseWriter, request *http.Request) {
	relationship := request.Context().Value(KeyRelationship{}).(*data.Relationship)
	log.Info("UpdateRelationships request", "id", relationship.ID)

	// Update relationship
	err := relationshipHandler.db.UpdateRelationship(relationship)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorUserNotFound:
		log.Error(err, "A UserID doesn't exist")
		http.Error(responseWriter, "A UserID doesn't exist", http.StatusBadRequest)
		return
	case data.ErrorSameUserID:
		log.Error(err, "Users in the relationship with same userID")
		http.Error(responseWriter, "Users in the relationship with same userID", http.StatusBadRequest)
		return
	case data.ErrorRelationshipExist:
		log.Error(err, "Relationship already exist")
		http.Error(responseWriter, "Relationship already exist", http.StatusBadRequest)
		return
	case data.ErrorRelationshipNotFound:
		log.Error(err, "Relationship not found")
		http.Error(responseWriter, "Relationship not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error updating relationship")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
