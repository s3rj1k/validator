package buildin

import (
	"fmt"
	"math/big"
)

// Number represents a casted integer
type Number struct {
	Value *big.Int
}

// NewNumber creates new Number object as pointer
func NewNumber(x int64) *Number {
	n := new(Number)
	n.Value = big.NewInt(x)

	return n
}

// SetUint64 sets uint64 to Number object
func (x *Number) SetUint64(u uint64) {
	x.Value.SetUint64(u)
}

// SetInt64 sets int64 to Number object
func (x *Number) SetInt64(i int64) {
	x.Value.SetInt64(i)
}

// IsGreater returns true if x > y
func (x *Number) IsGreater(y *Number) bool {
	return x.Value.Cmp(y.Value) == 1
}

// IsLess returns true if x < y
func (x *Number) IsLess(y *Number) bool {
	return x.Value.Cmp(y.Value) == -1
}

// IsEqual returns true if x == y
func (x *Number) IsEqual(y *Number) bool {
	return x.Value.Cmp(y.Value) == 0
}

// IsGreaterOrEqual returns true if x >= y
func (x *Number) IsGreaterOrEqual(y *Number) bool {
	if x.IsGreater(y) {
		return true
	}

	if x.IsEqual(y) {
		return true
	}

	return false
}

// IsLessOrEqual returns true if x <= y
func (x *Number) IsLessOrEqual(y *Number) bool {
	if x.IsLess(y) {
		return true
	}

	if x.IsEqual(y) {
		return true
	}

	return false
}

// InRange returns true if min < x < y
func (x *Number) InRange(min, max *Number) bool {
	return x.IsGreater(min) && x.IsLess(max)
}

// InRangeOrEqual returns true if min <= x <= y
func (x *Number) InRangeOrEqual(min, max *Number) bool {
	return x.IsGreaterOrEqual(min) && x.IsLessOrEqual(max)
}

// IsNegative returns true when x < 0
func (x *Number) IsNegative() bool {
	return x.Value.Sign() == -1
}

// IsPositive returns true when x >= 0
func (x *Number) IsPositive() bool {
	return !x.IsNegative()
}

// casts a into Number. Returns error if a is nil or not a integer
func cast(a interface{}) (*Number, error) {
	n := NewNumber(0)

	if a == nil {
		return n, nil
	}

	switch a := a.(type) {
	case int8:
		n.SetInt64(int64(a))

		return n, nil
	case int16:
		n.SetInt64(int64(a))

		return n, nil
	case int32:
		n.SetInt64(int64(a))

		return n, nil
	case int:
		n.SetInt64(int64(a))

		return n, nil
	case int64:
		n.SetInt64(a)

		return n, nil
	case uint8:
		n.SetUint64(uint64(a))

		return n, nil
	case uint16:
		n.SetUint64(uint64(a))

		return n, nil
	case uint32:
		n.SetUint64(uint64(a))

		return n, nil
	case uint:
		n.SetUint64(uint64(a))

		return n, nil
	case uint64:
		n.SetUint64(a)

		return n, nil
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
