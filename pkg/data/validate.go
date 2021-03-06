package data

import (
	"fmt"

	"github.com/go-playground/validator"
)

// ErrorInvalidRelationshipType : Invalid RelationshipType specific error
var ErrorInvalidRelationshipType = fmt.Errorf("invalid RelationshipType")

// ValidateRelationship a relationship with json validation
func (relationship *Relationship) ValidateRelationship() error {
	validate := validator.New()

	err := validate.RegisterValidation("isRelationshipType", validateIsRelationshipType)
	if err != nil {
		panic(ErrorInvalidRelationshipType)
	}
	
	return validate.Struct(relationship)
}

// validates the relationship type is valid
func validateIsRelationshipType(fieldLevel validator.FieldLevel) bool {
	relationshipType := fieldLevel.Field().String()

	switch relationshipType {
    case string(None), string(Friend), string(Blocked), string(PendingIncoming), string(PendingOutgoing):
        return true
    }
	return false
}
