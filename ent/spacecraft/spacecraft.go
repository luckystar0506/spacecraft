// Code generated by ent, DO NOT EDIT.

package spacecraft

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the spacecraft type in the database.
	Label = "spacecraft"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldClass holds the string denoting the class field in the database.
	FieldClass = "class"
	// FieldCrew holds the string denoting the crew field in the database.
	FieldCrew = "crew"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeArmaments holds the string denoting the armaments edge name in mutations.
	EdgeArmaments = "armaments"
	// Table holds the table name of the spacecraft in the database.
	Table = "spacecrafts"
	// ArmamentsTable is the table that holds the armaments relation/edge.
	ArmamentsTable = "spacecraft_armaments"
	// ArmamentsInverseTable is the table name for the SpacecraftArmament entity.
	// It exists in this package in order to avoid circular dependency with the "spacecraftarmament" package.
	ArmamentsInverseTable = "spacecraft_armaments"
	// ArmamentsColumn is the table column denoting the armaments relation/edge.
	ArmamentsColumn = "spacecraft_id"
)

// Columns holds all SQL columns for spacecraft fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldClass,
	FieldCrew,
	FieldImage,
	FieldValue,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Spacecraft queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByClass orders the results by the class field.
func ByClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClass, opts...).ToFunc()
}

// ByCrew orders the results by the crew field.
func ByCrew(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCrew, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByArmamentsCount orders the results by armaments count.
func ByArmamentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newArmamentsStep(), opts...)
	}
}

// ByArmaments orders the results by armaments terms.
func ByArmaments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArmamentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newArmamentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArmamentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ArmamentsTable, ArmamentsColumn),
	)
}
