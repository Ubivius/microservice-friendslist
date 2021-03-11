package data

import (
	"fmt"

	"github.com/go-playground/validator"
)

// ErrorInvalidRelationshipType : Invalid RelationshipType specific error
var ErrorInvalidRelationshipType = fmt.Errorf("Invalid RelationshipType")

// ValidateRelationship a relationship with json validation and customer SKU validator
func (relationship *Relationship) ValidateRelationship() error {
	validate := validator.New()
	err1 := validate.RegisterValidation("exist", validateExist)
	errRelationshipType := validate.RegisterValidation("isRelationshipType", validateIsRelationshipType)
	if err1 != nil {
		panic(err1)
	} else if errRelationshipType != nil {
		panic(ErrorInvalidRelationshipType)
	}
	
	return validate.Struct(relationship)
}

// validates the user exist
func validateExist(fieldLevel validator.FieldLevel) bool {
	// validation of the UserID with a call to microservice-user 
	return true
}

// validates the relationship type is valid
func validateIsRelationshipType(fieldLevel validator.FieldLevel) bool {
	relationshipType := int(fieldLevel.Field().Int())

	if relationshipType < int(None) || relationshipType > int(PendingOutgoing) {
		return false
	}

	return true
}
