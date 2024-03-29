// Code generated by entc, DO NOT EDIT.

package emailauth

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the emailauth type in the database.
	Label = "email_auth"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldLogged holds the string denoting the logged field in the database.
	FieldLogged = "logged"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// Table holds the table name of the emailauth in the database.
	Table = "email_auths"
)

// Columns holds all SQL columns for emailauth fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldEmail,
	FieldLogged,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultLogged holds the default value on creation for the "logged" field.
	DefaultLogged bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
