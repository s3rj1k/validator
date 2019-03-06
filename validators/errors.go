package validators

import (
	"errors"
)

var (
	// ErrBadNumType is an error for unsupported type provided to number validators
	ErrBadNumType = errors.New("unsupported type provided to number validator")

	// ErrBadSliceType is an error for unsupported type provided to slice validators
	ErrBadSliceType = errors.New("unsupported type provided to slice validator")

	// ErrNilValue is an error for encountering nil in validator
	ErrNilValue = errors.New("nil value must not be provided to validator")
)
