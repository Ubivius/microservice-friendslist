package router

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/handlers"
	"github.com/Ubivius/pkg-telemetry/metrics"
	tokenValidation "github.com/Ubivius/shared-authentication/pkg/auth"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

// New : Mux route handling with gorilla/mux
func New(relationshipHandler *handlers.RelationshipsHandler) *mux.Router {
	log.Info("Starting router")
	router := mux.NewRouter()
	router.Use(otelmux.Middleware("friendslist"))
	router.Use(metrics.RequestCountMiddleware)

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.Use(tokenValidation.Middleware)
	getRouter.HandleFunc("/friends/{user_id:[0-9a-z-]+}", relationshipHandler.GetFriendsListByUserID)
	getRouter.HandleFunc("/invites/{user_id:[0-9a-z-]+}", relationshipHandler.GetInvitesListByUserID)

	//Health Check
	healthRouter := router.Methods(http.MethodGet).Subrouter()
	healthRouter.HandleFunc("/health/live", relationshipHandler.LivenessCheck)
	healthRouter.HandleFunc("/health/ready", relationshipHandler.ReadinessCheck)

	// Put router
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.Use(tokenValidation.Middleware)
	putRouter.HandleFunc("/relationships", relationshipHandler.UpdateRelationships)
	putRouter.Use(relationshipHandler.MiddlewareRelationshipValidation)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.Use(tokenValidation.Middleware)
	postRouter.HandleFunc("/relationships", relationshipHandler.AddRelationship)
	postRouter.Use(relationshipHandler.MiddlewareRelationshipValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.Use(tokenValidation.Middleware)
	deleteRouter.HandleFunc("/relationships/{id:[0-9a-z-]+}", relationshipHandler.Delete)

	return router
}
