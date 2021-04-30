package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// AddRelationship creates a new relationship from the received JSON
func (relationshipHandler *RelationshipsHandler) AddRelationship(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("AddRelationship request")
	relationship := request.Context().Value(KeyRelationship{}).(*data.Relationship)

	err := relationshipHandler.db.AddRelationship(relationship)
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
	default:
		log.Error(err, "Error adding relationship")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
