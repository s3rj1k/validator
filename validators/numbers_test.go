package validators

import (
	"fmt"
	"math"
	"testing"
)

// slice of zero values of each integer type for testing purposes
var zeros = []interface{}{int(0), int8(0), int16(0), int32(0), int64(0),
	uint(0), uint8(0), uint16(0), uint32(0), uint64(0),
	uintptr(0), rune(0), byte(0)}

// slice of nonzero (2) values of each integer type for testing purposes
var nonzeros2 = []interface{}{
	int(2), int8(2), int16(2), int32(2), int64(2),
	uint(2), uint8(2), uint16(2), uint32(2), uint64(2), uintptr(2),
	rune(2), byte(2)}

// slice of nonzero (22) values of each integer type for testing purposes
var nonzeros10 = []interface{}{
	int(22), int8(22), int16(22), int32(22), int64(22),
	uint(22), uint8(22), uint16(22), uint32(22), uint64(22), uintptr(22),
	rune(22), byte(22)}

// slice of random types for testing purposes
var randomTypes = []interface{}{
	"string",
	true,
	[]int{1, 2, 3},
	struct{}{},
}

func Test_IsNil(t *testing.T) {

	var a interface{} // obviously nil

	if !isNil(a) {
		t.Errorf("interface is not identified as nil")
	}

	a = new(uintptr) // not nil

	if isNil(a) {
		t.Errorf("interface is mistakenly identified as nil")
	}

	a = 0 // not nil

	if isNil(a) {
		t.Errorf("interface is mistakenly identified as nil")
	}

	var b struct {
		a int
	}

	a = b.a // not nil

	if isNil(a) {
		t.Errorf("interface is mistakenly identified as nil")
	}
}

func Test_AreComparable(t *testing.T) {

	var a, b interface{}

	a = int(10) // signed integers are comparable
	b = int16(-99)

	if !areComparable(a, b) {
		t.Errorf("values are comparable but identified as not")
	}

	a = uintptr(10) // unsigned integers are comparable
	b = uint32(99)

	if !areComparable(a, b) {
		t.Errorf("values are comparable but identified as not")
	}

	a = int(10) // singed cannot be compared with unsigned
	b = uint32(99)

	if areComparable(a, b) {
		t.Errorf("values are not comparable but identified as such")
	}

	a = nil // nil is not comparable
	b = 1

	if areComparable(a, b) {
		t.Errorf("values are not comparable but identified as such")
	}

	a = 2 // string is not comparable
	b = "asd"

	if areComparable(a, b) {
		t.Errorf("values are not comparable but identified as such")
	}
}

func Test_AreSameTypeNumbers(t *testing.T) {

	var a, b interface{}

	a = 1 // same type
	b = 2

	if !areSameTypeNumbers(a, b) {
		t.Errorf("values are of same type but identified as not")
	}

	a = "asd" // same type
	b = "43666"

	if !areSameTypeNumbers(a, b) {
		t.Errorf("values are of same type but identified as not")
	}

	a = "asd"      // this function can distinct only between number types
	b = struct{}{} // this will return true

	if !areSameTypeNumbers(a, b) {
		t.Errorf("values are not of number types and expected to be identified as the same")
	}

	a = int16(1)
	b = int8(4)

	if areSameTypeNumbers(a, b) {
		t.Errorf("values are not of the same type but identifed as such")
	}

	a = 2
	b = nil

	if areSameTypeNumbers(a, b) {
		t.Errorf("values are not of the same type but identifed as such")
	}
}

func Test_Cast(t *testing.T) {

	var a, b, casteda, castedb interface{}

	a = int8(math.MinInt8)
	b = int8(math.MaxInt8)

	casteda, castedb = castBoth(a, b)

	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
		t.Errorf("values modified during casting")
	}

	a = int16(math.MinInt16)
	b = int16(math.MaxInt16)

	casteda, castedb = castBoth(a, b)

	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
		t.Errorf("values modified during casting")
	}

	a = int32(math.MinInt32)
	b = int32(math.MaxInt32)

	casteda, castedb = castBoth(a, b)

	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
		t.Errorf("values modified during casting")
	}

	a = uint8(math.MaxUint8)
	b = uint16(math.MaxUint16)

	casteda, castedb = castBoth(a, b)

	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
		t.Errorf("values modified during casting")
	}

	a = uint32(math.MaxUint32)
	b = uintptr(math.MaxUint64)

	casteda, castedb = castBoth(a, b)

	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
		t.Errorf("values modified during casting")
	}
}

func valuesDiffer(a, b interface{}) bool {
	return fmt.Sprintf("%d", a) != fmt.Sprintf("%d", b)
}
