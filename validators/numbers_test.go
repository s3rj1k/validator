package validators

// func Test_cast(t *testing.T) {
// 	var f1, f2 interface{}
// 	var oldf1, oldf2 interface{}

// 	// cast is OK
// 	f1 = int8(math.MaxInt8)
// 	f2 = int16(math.MaxInt16)
// 	oldf1 = f1
// 	oldf2 = f2

// 	isSigned, err := castOld(&f1, &f2)
// 	if err != nil {
// 		t.Errorf("unexpected error")
// 	}
// 	if isSigned != true {
// 		t.Errorf("sign identification error")
// 	}
// 	if valuesDiffer(f1, oldf1) || valuesDiffer(f2, oldf2) {
// 		t.Errorf("values changed during casting")
// 	}

// 	// cast is OK
// 	f1 = uint32(math.MaxUint32)
// 	f2 = uint16(math.MaxUint16)
// 	oldf1 = f1
// 	oldf2 = f2

// 	isSigned, err = castOld(&f1, &f2)
// 	if err != nil {
// 		t.Errorf("unexpected error")
// 	}
// 	if isSigned != false {
// 		t.Errorf("sign identification error")
// 	}
// 	if valuesDiffer(f1, oldf1) || valuesDiffer(f2, oldf2) {
// 		t.Errorf("values changed during casting")
// 	}

// 	// cast error: different signs
// 	f1 = int8(math.MinInt8)
// 	f2 = uint16(math.MaxUint16)
// 	isSigned, err = castOld(&f1, &f2)
// 	if err == nil {
// 		t.Errorf("expecting ErrDiffSigns not returned")
// 	}

// 	// cast OK
// 	f1 = []int{11, 24}
// 	f2 = int64(math.MaxInt64)
// 	isSigned, err = castOld(&f1, &f2)
// 	if err != nil {
// 		t.Errorf("unexpected error")
// 	}
// 	if isSigned != true {
// 		t.Errorf("sign identification error")
// 	}

// 	// cast error: different signs
// 	f1 = []uint{11, 24}
// 	f2 = int64(math.MaxInt64)
// 	isSigned, err = castOld(&f1, &f2)
// 	if err == nil {
// 		t.Errorf("expecting ErrDiffSigns not returned")
// 	}

// }

// func Test_casted(t *testing.T) {

// 	var a, b, casteda, castedb interface{}
// 	var err error

// 	a = int8(math.MinInt8)
// 	b = int8(math.MaxInt8)

// 	casteda, _, err = castedOld(a)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	castedb, _, err = castedOld(b)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
// 		t.Errorf("values modified during casting")
// 	}

// 	a = int16(math.MinInt16)
// 	b = int16(math.MaxInt16)

// 	casteda, _, err = castedOld(a)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	castedb, _, err = castedOld(b)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
// 		t.Errorf("values modified during casting")
// 	}

// 	a = int32(math.MinInt32)
// 	b = int32(math.MaxInt32)

// 	casteda, _, err = castedOld(a)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	castedb, _, err = castedOld(b)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
// 		t.Errorf("values modified during casting")
// 	}

// 	a = uint8(math.MaxUint8)
// 	b = uint16(math.MaxUint16)

// 	casteda, _, err = castedOld(a)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	castedb, _, err = castedOld(b)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
// 		t.Errorf("values modified during casting")
// 	}

// 	a = uint32(math.MaxUint32)
// 	b = uintptr(math.MaxUint64)

// 	casteda, _, err = castedOld(a)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	castedb, _, err = castedOld(b)
// 	if err != nil {
// 		t.Errorf("unexpected error %s", err)
// 	}

// 	if valuesDiffer(a, casteda) || valuesDiffer(b, castedb) {
// 		t.Errorf("values modified during casting")
// 	}

// 	a = nil
// 	b = "bad type"

// 	casteda, _, err = castedOld(a)
// 	if err == nil {
// 		t.Errorf("expected error for nil value, not returned")
// 	}

// 	castedb, _, err = castedOld(b)
// 	if err == nil {
// 		t.Errorf("expected error for wrong type, not returned")
// 	}

// }

// func valuesDiffer(a, b interface{}) bool {
// 	return fmt.Sprintf("%d", a) != fmt.Sprintf("%d", b)
// }
