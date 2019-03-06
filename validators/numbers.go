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
