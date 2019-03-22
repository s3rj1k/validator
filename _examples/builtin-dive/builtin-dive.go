package main

import (
	"fmt"
	"strings"

	v "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/validators"
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
			Field: u.names,
		},
		&vv.StringSliceDive{
			Validator: &vv.StringExclusion{ // error will be added for John
				Name:      "users.names",
				Blacklist: []string{"Donald", "Duck", "John"},
			},
			Field: []string{"Bob", "Alice", "John"},
		},
		&vv.NumberSliceDive{
			Validator: &vv.NumberInRange{ // error will be added for 0
				Name:       "users.ids",
				Min:        1,
				Max:        100,
				CheckEqual: true, // now 1 and 100 inclusive
			},
			Field: u.ids,
		},
		&vv.SliceLengthInRange{ // error will be added since ids is longer than names
			Name:  "users.ids",
			Field: u.ids,
			Min:   len(u.names),
			Max:   len(u.names),
		},
	)
	if e != nil {
		// all errors in valid JSON
		fmt.Printf("e.JSON() = %s\n", e.JSON())

		// e.Lookup() looks for errors with given prefix.
		// to get all errors for 'users' use e.Lookup("users")
		fmt.Print(`e.Lookup("users")`)
		fmt.Println()
		printmap(e.Lookup("users"))

		// to get all errors for 'users.ids' use e.Lookup("users.ids")
		fmt.Print(`e.Lookup("users.ids")`)
		fmt.Println()
		printmap(e.Lookup("users.ids"))
	}
}

func printmap(m map[string][]string) {

	for k, v := range m {
		fmt.Printf("Key: '%s'; Value: [%s]\n", k, strings.Join(v, ", "))
	}
}
