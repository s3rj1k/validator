package main

import (
	"fmt"
	"strings"

	v "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/buildin"
)

type users struct {
	names []string
	ids   []uint32
}

func main() {
	u := &users{
		names: []string{"Bob", "Alice", "John"},
		ids:   []uint32{12, 13, 14, 0},
	}

	e := v.Validate(
		&vv.StringSliceDive{
			Validator: &vv.StringIsAlpha{
				Name: "users.names",
			},
			Field: u.names, // all names are alpha - no errors
		},
		&vv.StringSliceDive{
			Validator: &vv.StringExclusion{
				Name:      "users.names",
				Blacklist: []string{"Donald", "Duck", "John"},
			},
			Field: []string{"Bob", "Alice", "John"}, // John is in the blacklist
		},
		&vv.NumberSliceDive{
			Validator: &vv.NumberInRange{
				Name:       "users.ids",
				Min:        1,
				Max:        100,
				CheckEqual: true, // now 1 and 100 inclusive
			},
			Field: u.ids, // 0 is not in min-max range
		},
		&vv.SliceLengthInRange{
			Name:  "users.ids",
			Field: u.ids,
			Min:   len(u.names),
			Max:   len(u.names), // error will be added since len(ids) > len(names)
		},
	)
	if e != nil {
		// all errors in valid JSON
		fmt.Println(e.JSON())
		// Output:
		// {"users.ids":["[12 13 14 0] length=4 not in range(3, 3)"],"users.ids[3]":["'0' not in range(1, 100) (inclusive)"],
		// "users.names[2]":["users.names[2] is in the blacklist [Donald Duck John]"]}

		// e.Lookup(key) looks for errors with given prefix and return a map.
		// to get all errors for 'users' use e.Lookup("users")
		printmap(e.Lookup("users"))
		// Output:
		// Key: 'users.ids[3]'; Value: ['0' not in range(1, 100) (inclusive)]
		// Key: 'users.ids'; Value: [[12 13 14 0] length=4 not in range(3, 3)]
		// Key: 'users.names[2]'; Value: [users.names[2] is in the blacklist [Donald Duck John]]

		// to get all errors for 'users.ids' use e.Lookup("users.ids")
		printmap(e.Lookup("users.ids"))
		// Output:
		// Key: 'users.ids[3]'; Value: ['0' not in range(1, 100) (inclusive)]
		// Key: 'users.ids'; Value: [[12 13 14 0] length=4 not in range(3, 3)]
	}
}

func printmap(m map[string][]string) {
	for k, v := range m {
		fmt.Printf("Key: '%s'; Value: [%s]\n", k, strings.Join(v, ", "))
	}
}
