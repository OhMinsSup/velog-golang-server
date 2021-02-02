// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// EmailAuthsColumns holds the columns for the "email_auths" table.
	EmailAuthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "code", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "logged", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// EmailAuthsTable holds the schema information for the "email_auths" table.
	EmailAuthsTable = &schema.Table{
		Name:        "email_auths",
		Columns:     EmailAuthsColumns,
		PrimaryKey:  []*schema.Column{EmailAuthsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "emailauth_code",
				Unique:  false,
				Columns: []*schema.Column{EmailAuthsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true, Size: 255},
		{Name: "is_certified", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[2]},
			},
		},
	}
	// UserProfilesColumns holds the columns for the "user_profiles" table.
	UserProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "display_name", Type: field.TypeString, Size: 255},
		{Name: "short_bio", Type: field.TypeString, Size: 255},
		{Name: "about", Type: field.TypeString, Size: 2147483647},
		{Name: "profile_links", Type: field.TypeJSON},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_user_profile", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// UserProfilesTable holds the schema information for the "user_profiles" table.
	UserProfilesTable = &schema.Table{
		Name:       "user_profiles",
		Columns:    UserProfilesColumns,
		PrimaryKey: []*schema.Column{UserProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_profiles_users_user_profile",
				Columns: []*schema.Column{UserProfilesColumns[8]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EmailAuthsTable,
		UsersTable,
		UserProfilesTable,
	}
)

func init() {
	UserProfilesTable.ForeignKeys[0].RefTable = UsersTable
}
