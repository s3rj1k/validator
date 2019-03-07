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
	// DefaultMinUserUID is a default value for MinUserUID, used then parsing of 'login.defs' fails.
	DefaultMinUserUID uint64 = 1000

	// DefaultMaxUserUID is a default value for MaxUserUID, used then parsing of 'login.defs' fails.
	DefaultMaxUserUID uint64 = 60000
)

// NumberIsValidUserUIDError is a function that defines error message returned by NumberIsValidUserUID validator.
// nolint: gochecknoglobals
var NumberIsValidUserUIDError = func(v *NumberIsValidUserUID) string {
	return fmt.Sprintf("%d is not valid user UID", v.Field)
}

// NumberIsValidUserUID is a validator object.
type NumberIsValidUserUID struct {
	Name  string
	Field interface{}
}

// Validate adds an error if the Field is in range of UID_MIN, UID_MAX from '/etc/login.defs'.
func (v *NumberIsValidUserUID) Validate(e *validator.Errors) {

	fNum, err := cast(v.Field)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	minUserUID, maxUserUID := readUserUIDRange(LoginDefsPath)

	//  for os.Chown func a uid or gid of -1 means to not change that value
	if fNum.isNegative && fNum.Value == 1 {
		return
	}

	if fNum.Value >= minUserUID &&
		fNum.Value <= maxUserUID &&
		!fNum.isNegative {

		return
	}

	e.Add(v.Name, NumberIsValidUserUIDError(v))
}

// SetField sets validator field.
func (v *NumberIsValidUserUID) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberIsValidUserUID) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberIsValidUserUID) GetName() string {
	return v.Name
}

// readUserUIDRange parses 'login.defs' file.
func readUserUIDRange(path string) (uint64, uint64) {

	var (
		minUserUID = DefaultMinUserUID
		maxUserUID = DefaultMaxUserUID
	)

	fd, err := os.Open(path)
	if err != nil {
		return minUserUID, maxUserUID
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
				minUserUID = i
			}
		}

		if fields[0] == "UID_MAX" {
			if i, err := strconv.ParseUint(fields[1], 10, 64); err == nil {
				maxUserUID = i
			}
		}

	}

	return minUserUID, maxUserUID
}
