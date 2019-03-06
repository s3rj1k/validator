package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// SliceIsUniqueError is a function that defines error message returned by SliceIsUnique validator.
// nolint: gochecknoglobals
var SliceIsUniqueError = func(v *SliceIsUnique) string {
	return fmt.Sprintf("%s value is not unique", v.Name)
}

// SliceIsUnique is a validator object.
type SliceIsUnique struct {
	Name  string
	Field interface{}
}

// Validate adds an error if the slice in Field has not unique values.
// Supports all Go basic types except for bool https://tour.golang.org/basics/11
// nolint: gocyclo
func (v *SliceIsUnique) Validate(e *validator.Errors) {

	var dupl = []int{} // indexes of duplicates

	m := make(map[interface{}]struct{})

	if v.Field == nil {
		e.Add(v.Name, ErrNilValue.Error())

		return
	}

	switch field := v.Field.(type) {
	case []string:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []int8:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []int16:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []int32:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []int:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []int64:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []uintptr:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []uint8:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []uint16:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []uint32:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []uint:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []uint64:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []float32:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []float64:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []complex64:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	case []complex128:
		for i, v := range field {
			if _, ok := m[v]; ok {
				dupl = append(dupl, i)
			}

			m[v] = struct{}{}
		}
	default:
		e.Add(v.Name, ErrBadSliceType.Error())
		return
	}

	if len(dupl) == 0 {
		return
	}

	// assigning error to each duplicate element
	for _, ind := range dupl {
		v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), ind)
		e.Add(v.Name, SliceIsUniqueError(v))
	}
}
