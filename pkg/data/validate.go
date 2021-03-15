package data

import (
	"fmt"

	"github.com/go-playground/validator"
)

// ErrorInvalidRelationshipType : Invalid RelationshipType specific error
var ErrorInvalidRelationshipType = fmt.Errorf("Invalid RelationshipType")

// ErrorSameUserID : Invalid Relationship specific error
var ErrorSameUserID = fmt.Errorf("Can't build a relationship with the same UserID")

// ErrorRelationshipExist : Invalid Relationship specific error
var ErrorRelationshipExist = fmt.Errorf("A relationship with these two users already exists")

// ValidateRelationship a relationship with json validation and customer SKU validator
func (relationship *Relationship) ValidateRelationship() error {
	validate := validator.New()

	errRelationship := validateRelationship(relationship);
	if errRelationship != nil {
		return errRelationship
	}

	errUserExist := validate.RegisterValidation("userExist", validateUserExist)
	errRelationshipType := validate.RegisterValidation("isRelationshipType", validateIsRelationshipType)
	if errUserExist != nil {
		panic(errUserExist)
	} else if errRelationshipType != nil {
		panic(ErrorInvalidRelationshipType)
	}
	
	return validate.Struct(relationship)
}

// validates a relation
func validateRelationship(relationship *Relationship) error {
	if relationship.User1.UserID == relationship.User2.UserID {
		return ErrorSameUserID
	}else if relationshipExist(relationship.User1.UserID, relationship.User2.UserID){
		return ErrorRelationshipExist
	}
	return nil
}

// validates the user exist
func validateUserExist(fieldLevel validator.FieldLevel) bool {
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
