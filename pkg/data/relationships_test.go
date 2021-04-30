package data

import "testing"

func TestChecksValidation(t *testing.T) {
	relationship := &Relationship{
		User1:      	User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: Friend},
		User2: 			User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44", RelationshipType: Friend},
		ConversationID:	"",
	}

	err := relationship.ValidateRelationship()

	if err != nil {
		t.Fatal(err)
	}
}

func TestInvalidRelationshipType(t *testing.T) {
	relationship := &Relationship{
		User1:      	User{UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8", RelationshipType: "Deleted"},
		User2: 			User{UserID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44", RelationshipType: Friend},
		ConversationID:	"",
	}

	err := relationship.ValidateRelationship()

	if !(err != nil && err.Error() == "Key: 'Relationship.User1.RelationshipType' Error:Field validation for 'RelationshipType' failed on the 'isRelationshipType' tag") {
		t.Errorf("A relationship type of value %s passed but RelationshipType need to be between %s and %s", relationship.User1.RelationshipType, None, PendingOutgoing)
	}
}
