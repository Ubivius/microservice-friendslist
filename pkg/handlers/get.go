package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-friendslist/pkg/data"
	"go.opentelemetry.io/otel"
)

// GetFriendsListByUserID returns all the relationships of a user from the database
func (relationshipHandler *RelationshipsHandler) GetFriendsListByUserID(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("friendslist").Start(request.Context(), "getFriendsListByUserID")
	defer span.End()
	id := getUserID(request)

	log.Info("GetFriendsListByUserID request for userID", "id", id)

	friends, err := relationshipHandler.db.GetFriendsListByUserID(request.Context(), id)
	switch err {
	case nil:
		detailedFriends, err := GetUserDetails(id, *friends)
		if err != nil {
			log.Error(err, "Error fetching users details")
		}

		err = json.NewEncoder(responseWriter).Encode(detailedFriends)
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
	_, span := otel.Tracer("friendslist").Start(request.Context(), "getInvitesListByUserId")
	defer span.End()
	id := getUserID(request)

	log.Info("GetInvitesListByUserID request for userID", "id", id)

	invites, err := relationshipHandler.db.GetInvitesListByUserID(request.Context(), id)
	switch err {
	case nil:
		detailedInvites, err := GetUserDetails(id, *invites)
		if err != nil {
			log.Error(err, "Error fetching users details")
		}

		err = json.NewEncoder(responseWriter).Encode(detailedInvites)
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

func GetUserDetails(userID string, relations data.Relationships) (*data.DetailedRelationships, error){
	detailedRelationsList := data.DetailedRelationships{}

	for _, relation := range relations{
		userIDToFetch := relation.User1.UserID
		relationshipType := relation.User1.RelationshipType

		if(userID == relation.User1.UserID){
			userIDToFetch = relation.User2.UserID
			relationshipType = relation.User2.RelationshipType
		}

		detailedUser, err := GetUserByID(userIDToFetch)
		if (err != nil){
			return nil, err
		}
		detailedUser.RelationshipType = relationshipType

		detailedRelationship := data.DetailedRelationship{
			ID: relation.ID,
			User: *detailedUser,
			ConversationID: relation.ConversationID,
			CreatedOn: relation.CreatedOn,
			UpdatedOn: relation.UpdatedOn,
		}
		detailedRelationsList = append(detailedRelationsList, &detailedRelationship)
	}
	return &detailedRelationsList, nil
}

func GetUserByID(userID string) (*data.DetailedUser, error) {
	getUserByIDPath := data.MicroserviceUserPath + "/users/" + userID
	resp, err := http.Get(getUserByIDPath)
	if err != nil {
		return nil, err
	}

	detailedUser := &data.DetailedUser{}
	err = json.NewDecoder(resp.Body).Decode(detailedUser)
	if err != nil {
		return nil, err
	}

	return detailedUser, nil
}
