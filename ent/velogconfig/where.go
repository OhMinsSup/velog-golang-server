// Code generated by entc, DO NOT EDIT.

package velogconfig

import (
	"time"

	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// LogoTitle applies equality check predicate on the "logo_title" field. It's identical to LogoTitleEQ.
func LogoTitle(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogoTitle), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.VelogConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.VelogConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTitle)))
	})
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTitle)))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// LogoTitleEQ applies the EQ predicate on the "logo_title" field.
func LogoTitleEQ(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleNEQ applies the NEQ predicate on the "logo_title" field.
func LogoTitleNEQ(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleIn applies the In predicate on the "logo_title" field.
func LogoTitleIn(vs ...string) predicate.VelogConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogoTitle), v...))
	})
}

// LogoTitleNotIn applies the NotIn predicate on the "logo_title" field.
func LogoTitleNotIn(vs ...string) predicate.VelogConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogoTitle), v...))
	})
}

// LogoTitleGT applies the GT predicate on the "logo_title" field.
func LogoTitleGT(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleGTE applies the GTE predicate on the "logo_title" field.
func LogoTitleGTE(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleLT applies the LT predicate on the "logo_title" field.
func LogoTitleLT(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleLTE applies the LTE predicate on the "logo_title" field.
func LogoTitleLTE(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleContains applies the Contains predicate on the "logo_title" field.
func LogoTitleContains(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleHasPrefix applies the HasPrefix predicate on the "logo_title" field.
func LogoTitleHasPrefix(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleHasSuffix applies the HasSuffix predicate on the "logo_title" field.
func LogoTitleHasSuffix(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleIsNil applies the IsNil predicate on the "logo_title" field.
func LogoTitleIsNil() predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLogoTitle)))
	})
}

// LogoTitleNotNil applies the NotNil predicate on the "logo_title" field.
func LogoTitleNotNil() predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLogoTitle)))
	})
}

// LogoTitleEqualFold applies the EqualFold predicate on the "logo_title" field.
func LogoTitleEqualFold(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogoTitle), v))
	})
}

// LogoTitleContainsFold applies the ContainsFold predicate on the "logo_title" field.
func LogoTitleContainsFold(v string) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogoTitle), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VelogConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VelogConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VelogConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VelogConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VelogConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VelogConfig) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VelogConfig) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VelogConfig) predicate.VelogConfig {
	return predicate.VelogConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}
