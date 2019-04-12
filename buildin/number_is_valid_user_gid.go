package buildin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/s3rj1k/validator"
)

// nolint: gochecknoglobals
var (
	// DefaultMinUserGID is a default value for MinUserGID, used then parsing of 'login.defs' fails.
	DefaultMinUserGID uint64 = 1000

	// DefaultMaxUserGID is a default value for MaxUserGID, used then parsing of 'login.defs' fails.
	DefaultMaxUserGID uint64 = 60000
)

// NumberIsValidUserGIDError is a function that defines error message returned by NumberIsValidUserGID validator.
// nolint: gochecknoglobals
var NumberIsValidUserGIDError = func(v *NumberIsValidUserGID) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a valid user GID", NumFieldToString(v.Field))
}

// NumberIsValidUserGID is a validator object.
// Validate adds an error if the Field is in range of GID_MIN, GID_MAX from '/etc/login.defs'.
type NumberIsValidUserGID struct {
	Name    string
	Field   interface{}
	Message string
}

// Validate adds an error if the Field is in range of GID_MIN, GID_MAX from '/etc/login.defs'.
func (v *NumberIsValidUserGID) Validate(e *validator.Errors) {

	fNum, err := cast(v.Field)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	minUserGID, maxUserGID := ReadUserGIDRange(LoginDefsPath)

	//  for os.Chown func a uid or gid of -1 means to not change that value
	if fNum.isNegative && fNum.Value == 1 {
		return
	}

	if fNum.Value >= minUserGID &&
		fNum.Value <= maxUserGID &&
		!fNum.isNegative {

		return
	}

	e.Add(v.Name, NumberIsValidUserGIDError(v))
}

// SetField sets validator field.
func (v *NumberIsValidUserGID) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberIsValidUserGID) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberIsValidUserGID) GetName() string {
	return v.Name
}

// ReadUserGIDRange parses 'login.defs' file.
func ReadUserGIDRange(path string) (uint64, uint64) {

	var (
		minUserGID = DefaultMinUserGID
		maxUserGID = DefaultMaxUserGID
	)

	fd, err := os.Open(path)
	if err != nil {
		return minUserGID, maxUserGID
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
				minUserGID = i
			}
		}

		if fields[0] == "GID_MAX" {
			if i, err := strconv.ParseUint(fields[1], 10, 64); err == nil {
				maxUserGID = i
			}
		}

	}

	return minUserGID, maxUserGID
}
