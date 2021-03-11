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
	if err == data.ErrorRelationshipNotFound {
		relationshipHandler.logger.Println("[ERROR} relationship not found", err)
		http.Error(responseWriter, "Relationship not found", http.StatusNotFound)
		return
	}

	// Returns status, no content required
	responseWriter.WriteHeader(http.StatusNoContent)
}
