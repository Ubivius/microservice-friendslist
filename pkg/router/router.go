package router

import (
	"log"
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/handlers"
	"github.com/gorilla/mux"
)

// New : Mux route handling with gorilla/mux
func New(relationshipHandler *handlers.RelationshipsHandler, logger *log.Logger) *mux.Router {
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/friends/{userid:[0-9]+}", relationshipHandler.GetFriendsListByUserID)
	getRouter.HandleFunc("/invites/{userid:[0-9]+}", relationshipHandler.GetInvitesListByUserID)

	// Put router
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/relationships", relationshipHandler.UpdateRelationships)
	putRouter.Use(relationshipHandler.MiddlewareRelationshipValidation)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/relationships", relationshipHandler.AddRelationship)
	postRouter.Use(relationshipHandler.MiddlewareRelationshipValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/relationships/{id:[0-9]+}", relationshipHandler.Delete)

	return router
}
