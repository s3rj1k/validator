package validators

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/s3rj1k/validator"
)

// NumberIsValidGIDError is a function that defines error message returned by NumberIsValidGID validator.
// nolint: gochecknoglobals
var NumberIsValidGIDError = func(v *NumberIsValidGID) string {
	return fmt.Sprintf("%d is not valid GID", v.Field)
}

// NumberIsValidGID is a validator object.
type NumberIsValidGID struct {
	Name  string
	Field interface{}
}

// Validate adds an error if the Field is in range of GID_MIN, GID_MAX from '/etc/login.defs'.
func (v *NumberIsValidGID) Validate(e *validator.Errors) {

	fNum, err := cast(v.Field)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	minGID, maxGID := readGIDRange(LoginDefsPath)

	if fNum.Value >= minGID &&
		fNum.Value <= maxGID &&
		!fNum.isNegative {

		return
	}

	e.Add(v.Name, NumberIsValidGIDError(v))
}

// SetField sets validator field.
func (v *NumberIsValidGID) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberIsValidGID) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberIsValidGID) GetName() string {
	return v.Name
}

// readGIDRange parses 'login.defs' file.
func readGIDRange(path string) (uint64, uint64) {

	var (
		minGID uint64 = DefaultMinGID
		maxGID uint64 = DefaultMaxGID
	)

	fd, err := os.Open(path)
	if err != nil {
		return minGID, maxGID
	}

	defer func(fd *os.File) {
		_ = fd.Close()
	}(fd)

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {

		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			continue
		}

		if fields[0] == "GID_MIN" {
			if i, err := strconv.ParseUint(fields[1], 10, 64); err == nil {
				minGID = i
			}
		}

		if fields[0] == "GID_MAX" {
			if i, err := strconv.ParseUint(fields[1], 10, 64); err == nil {
				maxGID = i
			}
		}

	}

	return minGID, maxGID
}
