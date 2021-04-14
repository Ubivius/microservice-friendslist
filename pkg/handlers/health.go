package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// LivenessCheck determine when the application needs to be restarted
func (relationshipHandler *RelationshipsHandler) LivenessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("LivenessCheck")
	responseWriter.WriteHeader(http.StatusOK)
}

//ReadinessCheck verifies that the application is ready to accept requests
func (relationshipHandler *RelationshipsHandler) ReadinessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("ReadinessCheck")

	err := relationshipHandler.db.PingDB()

	if err != nil {
		log.Error(err, "DB unavailable")
		http.Error(responseWriter, "DB unavailable", http.StatusServiceUnavailable)
		return
	}

	readinessProbeMicroserviceUser := data.MicroserviceUserPath + "/health/ready"
	_, err = http.Get(readinessProbeMicroserviceUser)

	if err != nil {
		log.Error(err, "Microservice-user unavailable")
		http.Error(responseWriter, "Microservice-user unavailable", http.StatusServiceUnavailable)
		return
	}

	readinessProbeMicroserviceTextChat := data.MicroserviceTextChatPath + "/health/ready"
	_, err = http.Get(readinessProbeMicroserviceTextChat)

	if err != nil {
		log.Error(err, "Microservice-text-chat unavailable")
		http.Error(responseWriter, "Microservice-text-chat unavailable", http.StatusServiceUnavailable)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
