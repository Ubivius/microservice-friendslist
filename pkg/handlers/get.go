package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
)

// GetFriendsListByUserID returns all the relationships of a user from the database
func (relationshipHandler *RelationshipsHandler) GetFriendsListByUserID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getUserID(request)

	relationshipHandler.logger.Println("[DEBUG] getting friends list for userID", id)

	friends, err := data.GetFriendsListByUserID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(friends)
		if err != nil {
			relationshipHandler.logger.Println("[ERROR] serializing friends", err)
		}
	case data.ErrorRelationshipNotFound:
		relationshipHandler.logger.Println("[ERROR] fetching friends", err)
		http.Error(responseWriter, "Friends not found", http.StatusNotFound)
		return
	default:
		relationshipHandler.logger.Println("[ERROR] fetching friends", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}

// GetInvitesListByUserID returns all the invites of a user from the database
func (relationshipHandler *RelationshipsHandler) GetInvitesListByUserID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getUserID(request)

	relationshipHandler.logger.Println("[DEBUG] getting invites list for userID", id)

	invites, err := data.GetInvitesListByUserID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(invites)
		if err != nil {
			relationshipHandler.logger.Println("[ERROR] serializing invites", err)
		}
	case data.ErrorRelationshipNotFound:
		relationshipHandler.logger.Println("[ERROR] fetching invites", err)
		http.Error(responseWriter, "Invites not found", http.StatusNotFound)
		return
	default:
		relationshipHandler.logger.Println("[ERROR] fetching invites", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}
