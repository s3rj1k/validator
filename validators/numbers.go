package validators

// Number represents a casted integer
type Number struct {
	Value      uint64
	isNegative bool
}

// casts a into Number. Returns error if a is nil or not a integer
func cast(a interface{}) (*Number, error) {

	if a == nil {
		return nil, ErrNilValue
	}

	switch a := a.(type) {
	case int8:
		return &Number{uint64(a), a < 0}, nil
	case int16:
		return &Number{uint64(a), a < 0}, nil
	case int32:
		return &Number{uint64(a), a < 0}, nil
	case int:
		return &Number{uint64(a), a < 0}, nil
	case int64:
		return &Number{uint64(a), a < 0}, nil
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
