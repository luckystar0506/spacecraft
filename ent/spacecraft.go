// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"spacecraft/ent/spacecraft"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Spacecraft is the model entity for the Spacecraft schema.
type Spacecraft struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Class holds the value of the "class" field.
	Class string `json:"class,omitempty"`
	// Crew holds the value of the "crew" field.
	Crew int `json:"crew,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Value holds the value of the "value" field.
	Value float64 `json:"value,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SpacecraftQuery when eager-loading is set.
	Edges        SpacecraftEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SpacecraftEdges holds the relations/edges for other nodes in the graph.
type SpacecraftEdges struct {
	// Armaments holds the value of the armaments edge.
	Armaments []*SpacecraftArmament `json:"armaments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ArmamentsOrErr returns the Armaments value or an error if the edge
// was not loaded in eager-loading.
func (e SpacecraftEdges) ArmamentsOrErr() ([]*SpacecraftArmament, error) {
	if e.loadedTypes[0] {
		return e.Armaments, nil
	}
	return nil, &NotLoadedError{edge: "armaments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Spacecraft) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case spacecraft.FieldValue:
			values[i] = new(sql.NullFloat64)
		case spacecraft.FieldID, spacecraft.FieldCrew:
			values[i] = new(sql.NullInt64)
		case spacecraft.FieldName, spacecraft.FieldClass, spacecraft.FieldImage, spacecraft.FieldStatus:
			values[i] = new(sql.NullString)
		case spacecraft.FieldCreatedAt, spacecraft.FieldUpdatedAt, spacecraft.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Spacecraft fields.
func (s *Spacecraft) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case spacecraft.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case spacecraft.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case spacecraft.FieldClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class", values[i])
			} else if value.Valid {
				s.Class = value.String
			}
		case spacecraft.FieldCrew:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field crew", values[i])
			} else if value.Valid {
				s.Crew = int(value.Int64)
			}
		case spacecraft.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				s.Image = value.String
			}
		case spacecraft.FieldValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				s.Value = value.Float64
			}
		case spacecraft.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = value.String
			}
		case spacecraft.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case spacecraft.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case spacecraft.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = new(time.Time)
				*s.DeletedAt = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Spacecraft.
// This includes values selected through modifiers, order, etc.
func (s *Spacecraft) GetValue(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryArmaments queries the "armaments" edge of the Spacecraft entity.
func (s *Spacecraft) QueryArmaments() *SpacecraftArmamentQuery {
	return NewSpacecraftClient(s.config).QueryArmaments(s)
}

// Update returns a builder for updating this Spacecraft.
// Note that you need to call Spacecraft.Unwrap() before calling this method if this Spacecraft
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Spacecraft) Update() *SpacecraftUpdateOne {
	return NewSpacecraftClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Spacecraft entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Spacecraft) Unwrap() *Spacecraft {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Spacecraft is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Spacecraft) String() string {
	var builder strings.Builder
	builder.WriteString("Spacecraft(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("class=")
	builder.WriteString(s.Class)
	builder.WriteString(", ")
	builder.WriteString("crew=")
	builder.WriteString(fmt.Sprintf("%v", s.Crew))
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(s.Image)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", s.Value))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(s.Status)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := s.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Spacecrafts is a parsable slice of Spacecraft.
type Spacecrafts []*Spacecraft
