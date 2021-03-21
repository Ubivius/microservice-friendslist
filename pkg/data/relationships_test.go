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

func TestInvalidRelationshipType(t *testing.T) {
	relationship := &Relationship{
		User1:      	User{UserID: 1, RelationshipType: "Deleted"},
		User2: 			User{UserID: 2, RelationshipType: Friend},
		ConversationID:	1,
	}

	err := relationship.ValidateRelationship()

	if !(err != nil && err.Error() == "Key: 'Relationship.User1.RelationshipType' Error:Field validation for 'RelationshipType' failed on the 'isRelationshipType' tag") {
		t.Errorf("A relationship type of value %s passed but RelationshipType need to be between %s and %s", relationship.User1.RelationshipType, None, PendingOutgoing)
	}
}
