// Code generated by entc, DO NOT EDIT.

package ent

import (
	"gremlin-orm-sample/ent/car"
	"gremlin-orm-sample/ent/schema"
	"gremlin-orm-sample/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescID is the schema descriptor for id field.
	carDescID := carFields[0].Descriptor()
	// car.IDValidator is a validator for the "id" field. It is called by the builders before save.
	car.IDValidator = carDescID.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[1].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
}
