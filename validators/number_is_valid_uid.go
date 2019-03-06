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

// nolint: gochecknoglobals
var (
	// DefaultMinUID is a default value for MinUID, used then parsing of 'login.defs' fails.
	DefaultMinUID uint64 = 1000

	// DefaultMaxUID is a default value for MaxUID, used then parsing of 'login.defs' fails.
	DefaultMaxUID uint64 = 60000
)

// NumberIsValidUIDError is a function that defines error message returned by NumberIsValidUID validator.
// nolint: gochecknoglobals
var NumberIsValidUIDError = func(v *NumberIsValidUID) string {
	return fmt.Sprintf("%d is not valid UID", v.Field)
}

// NumberIsValidUID is a validator object.
type NumberIsValidUID struct {
	Name  string
	Field interface{}
}

// Validate adds an error if the Field is in range of UID_MIN, UID_MAX from '/etc/login.defs'.
func (v *NumberIsValidUID) Validate(e *validator.Errors) {

	fNum, err := cast(v.Field)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	minUID, maxUID := readUIDRange(LoginDefsPath)

	//  for os.Chown func a uid or gid of -1 means to not change that value
	if fNum.isNegative && fNum.Value == 1 {
		return
	}

	if fNum.Value >= minUID &&
		fNum.Value <= maxUID &&
		!fNum.isNegative {

		return
	}

	e.Add(v.Name, NumberIsValidUIDError(v))
}

// SetField sets validator field.
func (v *NumberIsValidUID) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberIsValidUID) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberIsValidUID) GetName() string {
	return v.Name
}

// readUIDRange parses 'login.defs' file.
func readUIDRange(path string) (uint64, uint64) {

	var (
		minUID = DefaultMinUID
		maxUID = DefaultMaxUID
	)

	fd, err := os.Open(path)
	if err != nil {
		return minUID, maxUID
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

		if fields[0] == "UID_MIN" {
			if i, err := strconv.ParseUint(fields[1], 10, 64); err == nil {
				minUID = i
			}
		}

		if fields[0] == "UID_MAX" {
			if i, err := strconv.ParseUint(fields[1], 10, 64); err == nil {
				maxUID = i
			}
		}

	}

	return minUID, maxUID
}
