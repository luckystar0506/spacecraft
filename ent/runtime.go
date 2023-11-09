// Code generated by ent, DO NOT EDIT.

package ent

import (
	"spacecraft/ent/schema"
	"spacecraft/ent/spacecraft"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	spacecraftFields := schema.Spacecraft{}.Fields()
	_ = spacecraftFields
	// spacecraftDescUpdatedAt is the schema descriptor for updated_at field.
	spacecraftDescUpdatedAt := spacecraftFields[7].Descriptor()
	// spacecraft.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	spacecraft.UpdateDefaultUpdatedAt = spacecraftDescUpdatedAt.UpdateDefault.(func() time.Time)
}
