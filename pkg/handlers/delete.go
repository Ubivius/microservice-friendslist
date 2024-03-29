package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
	"go.opentelemetry.io/otel"
)

// Delete a relationship with specified id from the database
func (relationshipHandler *RelationshipsHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("friendslist").Start(request.Context(), "deleteRelationship")
	defer span.End()
	id := getRelationshipID(request)
	log.Info("Delete relationship by ID request", "id", id)

	err := relationshipHandler.db.DeleteRelationship(request.Context(), id)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorRelationshipNotFound:
		log.Error(err, "Error deleting relationship, id does not exist")
		http.Error(responseWriter, "Relationship not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error deleting relationship")
		http.Error(responseWriter, "Error deleting relationship", http.StatusInternalServerError)
		return
	}
}
