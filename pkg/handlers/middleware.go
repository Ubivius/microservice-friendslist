package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// MiddlewareRelationshipValidation is used to validate incoming relationship JSONS
func (relationshipHandler *RelationshipsHandler) MiddlewareRelationshipValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		relationship := &data.Relationship{}

		err := json.NewDecoder(request.Body).Decode(relationship)
		if err != nil {
			relationshipHandler.logger.Println("[ERROR] deserializing relationship", err)
			http.Error(responseWriter, "Error reading relationship", http.StatusBadRequest)
			return
		}

		// validate the relationship
		err = relationship.ValidateRelationship()
		if err != nil {
			relationshipHandler.logger.Println("[ERROR] validating relationship", err)
			http.Error(responseWriter, fmt.Sprintf("Error validating relationship: %s", err), http.StatusBadRequest)
			return
		}

		// Add the relationship to the context
		ctx := context.WithValue(request.Context(), KeyRelationship{}, relationship)
		request = request.WithContext(ctx)

		// Call the next handler, which can be another middleware or the final handler
		next.ServeHTTP(responseWriter, request)
	})
}
