package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// AddRelationship creates a new relationship from the received JSON
func (relationshipHandler *RelationshipsHandler) AddRelationship(responseWriter http.ResponseWriter, request *http.Request) {
	relationshipHandler.logger.Println("Handle POST Relationship")
	relationship := request.Context().Value(KeyRelationship{}).(*data.Relationship)

	data.AddRelationship(relationship)
	responseWriter.WriteHeader(http.StatusNoContent)
}
