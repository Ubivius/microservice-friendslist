package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// Delete a relationship with specified id from the database
func (relationshipHandler *RelationshipsHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getRelationshipID(request)
	relationshipHandler.logger.Println("Handle DELETE relationship", id)

	err := relationshipHandler.db.DeleteRelationship(id)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorRelationshipNotFound:
		relationshipHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Relationship not found", http.StatusNotFound)
		return
	default:
		relationshipHandler.logger.Println("[ERROR] deleting relationship", err)
		http.Error(responseWriter, "Error deleting relationship", http.StatusInternalServerError)
		return
	}
}
