package buildin

import (
	"fmt"
)

// Number represents a casted integer
type Number struct {
	Value      uint64
	isNegative bool
}

// IsGreater returns true if x > y
func (x *Number) IsGreater(y *Number) bool {

	switch {
	case x.isNegative && !y.isNegative:
		return false
	case !x.isNegative && y.isNegative:
		return true
	case !x.isNegative && !y.isNegative:
		return x.Value > y.Value
	case x.isNegative && y.isNegative:
		return x.Value < y.Value
	}

	return false
}

// IsLess returns true if x < y
func (x *Number) IsLess(y *Number) bool {

	switch {
	case x.isNegative && !y.isNegative:
		return true
	case !x.isNegative && y.isNegative:
		return false
	case !x.isNegative && !y.isNegative:
		return x.Value < y.Value
	case x.isNegative && y.isNegative:
		return x.Value > y.Value
	}

	return false
}

// IsEqual returns true if x == y
func (x *Number) IsEqual(y *Number) bool {
	return x.Value == y.Value && x.isNegative == y.isNegative
}

// casts a into Number. Returns error if a is nil or not a integer
func cast(a interface{}) (*Number, error) {

	if a == nil {
		return &Number{uint64(0), false}, nil
	}

	switch a := a.(type) {
	case int8:
		if a < 0 {
			return &Number{uint64(a * -1), true}, nil
		}

		return &Number{uint64(a), false}, nil

	case int16:
		if a < 0 {
			return &Number{uint64(a * -1), true}, nil
		}

		return &Number{uint64(a), false}, nil

	case int32:
		if a < 0 {
			return &Number{uint64(a * -1), true}, nil
		}

		return &Number{uint64(a), false}, nil

	case int:
		if a < 0 {
			return &Number{uint64(a * -1), true}, nil
		}

		return &Number{uint64(a), false}, nil

	case int64:
		if a < 0 {
			return &Number{uint64(a * -1), true}, nil
		}

		return &Number{uint64(a), false}, nil

	case uintptr:
		return &Number{uint64(a), false}, nil
	case uint:
		return &Number{uint64(a), false}, nil
	case uint8:
		return &Number{uint64(a), false}, nil
	case uint16:
		return &Number{uint64(a), false}, nil
	case uint32:
		return &Number{uint64(a), false}, nil
	case uint64:
		return &Number{a, false}, nil
	}

	return nil, ErrBadNumType
}

// NumFieldToString returns string representation of number field
func NumFieldToString(field interface{}) string {

	if field == nil {
		return "0"
	}

	return fmt.Sprintf("%d", field)
}
