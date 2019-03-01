package validators

import "fmt"

const (
	// ErrUnsupportedNumberType is an error message for unsupported number type
	ErrUnsupportedNumberType = "unsupported number type"

	// ErrNumberTypesNotMatchError is an error message for different number types provided
	ErrNumberTypesNotMatchError = "types not match"

	// ErrNumbersNotComparable is an error message when number types provided cannot be correctly compared.
	// (signed and unsigned ints cannot be correctly compared)
	ErrNumbersNotComparable = "types cannot be compared"

	// ErrNilFields is an error message if nil interfaces are provided
	ErrNilFields = "nil fields are forbidden"
)

// NumberType is used to define number types used by number validators
type NumberType int

// These number types are supported by number validators.
const (
	Int NumberType = iota
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Invalid
)

// cast all signed integers to int64 and unsigned integers to uint64
// returnes caster interface, NumberType and signed bool
func cast(n interface{}) (interface{}, NumberType, bool) {
	switch n := n.(type) {
	case int:
		return int64(n), Int, true
	case int8:
		return int64(n), Int8, true
	case int16:
		return int64(n), Int16, true
	case int32:
		return int64(n), Int32, true
	case int64:
		return int64(n), Int64, true
	case uint:
		return uint64(n), Uint, false
	case uintptr:
		return uint64(n), Uintptr, false
	case uint8:
		return uint64(n), Uint8, false
	case uint16:
		return uint64(n), Uint16, false
	case uint32:
		return uint64(n), Uint32, false
	case uint64:
		return uint64(n), Uint64, false
	}
	return nil, Invalid, false
}

// castBoth casts 2 values of signed integers to int64 and unsigned integers to uint64
func castBoth(a, b interface{}) (interface{}, interface{}) {
	ca, _, _ := cast(a)
	cb, _, _ := cast(b)

	return ca, cb
}

// isNil returns true if at least one interface has nil underlying  value
func isNil(a ...interface{}) bool {
	for _, i := range a {

		if i == nil {
			return true
		}
	}
	return false
}

// areSameTypeNumbers returns true if both interfaces have the same dynamic types
func areSameTypeNumbers(a, b interface{}) bool {
	_, atype, _ := cast(a)
	_, btype, _ := cast(b)

	return atype == btype
}

// areComparable returns true if both a and b are signed or both are unsigned
func areComparable(a, b interface{}) bool {
	_, atype, asigned := cast(a)
	_, btype, bsigned := cast(b)

	// invalid types are not comparable
	if atype == Invalid || btype == Invalid {
		return false
	}

	return asigned == bsigned
}

// checkNums is a wrapper for various interfaces-numbers validations
func checkNums(a, b interface{}, needSameType bool) error {

	// nil fields are always bad - cannot happen if user provided a number
	if isNil(a, b) {
		return fmt.Errorf(ErrNilFields)
	}

	// sametypes are not essential, can be of different types if user allows
	if needSameType && !areSameTypeNumbers(a, b) {
		return fmt.Errorf(ErrNumberTypesNotMatchError)
	}

	// cannot be uncomparable - signed and unsigned integers are not correctly compared
	if !areComparable(a, b) {
		return fmt.Errorf(ErrNumbersNotComparable)
	}

	return nil
}
