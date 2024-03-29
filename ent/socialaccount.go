// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/OhMinsSup/story-server/ent/socialaccount"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// SocialAccount is the model entity for the SocialAccount schema.
type SocialAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SocialID holds the value of the "social_id" field.
	SocialID string `json:"social_id,omitempty"`
	// AccessToken holds the value of the "access_token" field.
	AccessToken string `json:"access_token,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// FkUserID holds the value of the "fk_user_id" field.
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SocialAccount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case socialaccount.FieldSocialID, socialaccount.FieldAccessToken, socialaccount.FieldProvider:
			values[i] = &sql.NullString{}
		case socialaccount.FieldCreatedAt, socialaccount.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		case socialaccount.FieldID, socialaccount.FieldFkUserID:
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type SocialAccount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SocialAccount fields.
func (sa *SocialAccount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case socialaccount.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sa.ID = *value
			}
		case socialaccount.FieldSocialID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field social_id", values[i])
			} else if value.Valid {
				sa.SocialID = value.String
			}
		case socialaccount.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				sa.AccessToken = value.String
			}
		case socialaccount.FieldProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider", values[i])
			} else if value.Valid {
				sa.Provider = value.String
			}
		case socialaccount.FieldFkUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fk_user_id", values[i])
			} else if value != nil {
				sa.FkUserID = *value
			}
		case socialaccount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sa.CreatedAt = value.Time
			}
		case socialaccount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sa.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SocialAccount.
// Note that you need to call SocialAccount.Unwrap() before calling this method if this SocialAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (sa *SocialAccount) Update() *SocialAccountUpdateOne {
	return (&SocialAccountClient{config: sa.config}).UpdateOne(sa)
}

// Unwrap unwraps the SocialAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sa *SocialAccount) Unwrap() *SocialAccount {
	tx, ok := sa.config.driver.(*txDriver)
	if !ok {
		panic("ent: SocialAccount is not a transactional entity")
	}
	sa.config.driver = tx.drv
	return sa
}

// String implements the fmt.Stringer.
func (sa *SocialAccount) String() string {
	var builder strings.Builder
	builder.WriteString("SocialAccount(")
	builder.WriteString(fmt.Sprintf("id=%v", sa.ID))
	builder.WriteString(", social_id=")
	builder.WriteString(sa.SocialID)
	builder.WriteString(", access_token=")
	builder.WriteString(sa.AccessToken)
	builder.WriteString(", provider=")
	builder.WriteString(sa.Provider)
	builder.WriteString(", fk_user_id=")
	builder.WriteString(fmt.Sprintf("%v", sa.FkUserID))
	builder.WriteString(", created_at=")
	builder.WriteString(sa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(sa.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SocialAccounts is a parsable slice of SocialAccount.
type SocialAccounts []*SocialAccount

func (sa SocialAccounts) config(cfg config) {
	for _i := range sa {
		sa[_i].config = cfg
	}
}
