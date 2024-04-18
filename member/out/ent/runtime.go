// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"todopoint/member/out/ent/member"
	"todopoint/member/out/ent/schema"
)

// The init function reads all schema descriptors with runtime codes
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescEmail is the schema descriptor for email field.
	memberDescEmail := memberFields[1].Descriptor()
	// member.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	member.EmailValidator = memberDescEmail.Validators[0].(func(string) error)
	// memberDescUsername is the schema descriptor for username field.
	memberDescUsername := memberFields[2].Descriptor()
	// member.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	member.UsernameValidator = memberDescUsername.Validators[0].(func(string) error)
	// memberDescPassword is the schema descriptor for password field.
	memberDescPassword := memberFields[3].Descriptor()
	// member.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	member.PasswordValidator = memberDescPassword.Validators[0].(func(string) error)
	// memberDescCreatedAt is the schema descriptor for created_at field.
	memberDescCreatedAt := memberFields[4].Descriptor()
	// member.DefaultCreatedAt holds the default value on creation for the created_at field.
	member.DefaultCreatedAt = memberDescCreatedAt.Default.(time.Time)
}