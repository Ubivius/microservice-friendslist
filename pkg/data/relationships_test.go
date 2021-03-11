package data

import "testing"

func TestChecksValidation(t *testing.T) {
	relationship := &Relationship{
		User1:      	User{UserID: 1, RelationshipType: Friend},
		User2: 			User{UserID: 2, RelationshipType: Friend},
		ConversationID:	1,
	}

	err := relationship.ValidateRelationship()

	if err != nil {
		t.Fatal(err)
	}
}
