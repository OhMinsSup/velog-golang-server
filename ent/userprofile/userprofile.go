// Code generated by entc, DO NOT EDIT.

package userprofile

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the userprofile type in the database.
	Label = "user_profile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldShortBio holds the string denoting the short_bio field in the database.
	FieldShortBio = "short_bio"
	// FieldAbout holds the string denoting the about field in the database.
	FieldAbout = "about"
	// FieldProfileLinks holds the string denoting the profile_links field in the database.
	FieldProfileLinks = "profile_links"
	// FieldThumbnail holds the string denoting the thumbnail field in the database.
	FieldThumbnail = "thumbnail"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the userprofile in the database.
	Table = "user_profiles"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "user_profiles"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_user_profile"
)

// Columns holds all SQL columns for userprofile fields.
var Columns = []string{
	FieldID,
	FieldDisplayName,
	FieldShortBio,
	FieldAbout,
	FieldProfileLinks,
	FieldThumbnail,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the UserProfile type.
var ForeignKeys = []string{
	"user_user_profile",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// ShortBioValidator is a validator for the "short_bio" field. It is called by the builders before save.
	ShortBioValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
