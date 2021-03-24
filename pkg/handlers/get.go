package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// GetFriendsListByUserID returns all the relationships of a user from the database
func (relationshipHandler *RelationshipsHandler) GetFriendsListByUserID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getUserID(request)

	log.Info("GetFriendsListByUserID request for userID","id", id)

	friends, err := relationshipHandler.db.GetFriendsListByUserID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(friends)
		if err != nil {
			log.Error(err, "Error serializing friends")
		}
		return
	case data.ErrorRelationshipNotFound:
		log.Error(err, "Friends not found")
		http.Error(responseWriter, "Friends not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error fetching friends")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}

// GetInvitesListByUserID returns all the invites of a user from the database
func (relationshipHandler *RelationshipsHandler) GetInvitesListByUserID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getUserID(request)

	log.Info("GetInvitesListByUserID request for userID","id", id)

	invites, err := relationshipHandler.db.GetInvitesListByUserID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(invites)
		if err != nil {
			log.Error(err, "Error serializing invites")
		}
		return
	case data.ErrorRelationshipNotFound:
		log.Error(err, "Invites not found")
		http.Error(responseWriter, "Invites not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error fetching invites")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}
