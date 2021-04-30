# microservice-friendslist
Friends list microservice for our online game framework.

## Friends list endpoints

`GET` `/friends/{user_id}` Returns all friend relationships of the specific user. `user_id=[string]`

`GET` `/invites/{user_id}` Resends all friend invitations for the specific user. `user_id=[string]`

`GET` `/health/live` Returns a Status OK when live.

`GET` `/health/ready` Returns a Status OK when ready or an error when dependencies are not available.

`POST` `/relationships` Add new relationship with specific data. Creates a conversation with [microservice-text-chat](https://github.com/Ubivius/microservice-text-chat) and add the conversation ID to the relationship. </br>
__Data Params__
```json
{
  "user_1": {
    "user_id":           "string, required",
    "relationship_type": "string, required",
  },
  "user_2": {
    "user_id":           "string, required",
    "relationship_type": "string, required",
  },
}
```
__Relationshipe type__
```
None            // user has no intrinsic relationship
Friend          // user is a friend
Blocked         // user is blocked
PendingIncoming	// user has a pending incoming friend request to connected user
PendingOutgoing	// current user has a pending outgoing friend request to user
```

`PUT` `/relationships` Update relationship data</br>
__Data Params__
```json
{
  "id":                  "string, required",
  "user_1": {
    "user_id":           "string",
    "relationship_type": "string",
  },
  "user_2": {
    "user_id":           "string",
    "relationship_type": "string",
  },
  "conversation_id":     "string",
}
```

`DELETE` `/relationships/{id}` Delete a relationship.  `id=[string]`
