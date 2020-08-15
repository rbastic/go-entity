package entity

import (
	"errors"
)

var (
	// ErrKeyIsMissing is returned when a key is missing from the entity.
	ErrKeyIsMissing = errors.New("key is missing")

	// ErrValueIsNil is returned when a value is casted to a default
	// result (i.e. 0, ""), but is originally nil in the Entity.
	ErrValueIsNil = errors.New("value is nil")

	ErrRefusingCast = errors.New("type conversion is illogical - refusing to cast")
)
