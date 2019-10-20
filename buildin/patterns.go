package buildin

import (
	"regexp"
)

// Basic regular expressions for validating strings
// nolint: lll
const (
	Email              = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	Alpha              = "^[a-zA-Z]+$"
	Alphanumeric       = "^[a-zA-Z0-9]+$"
	Numeric            = "^[0-9]+$"
	Integer            = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Float              = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	Hexadecimal        = "^[0-9a-fA-F]+$"
	Hexcolor           = "^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"
	RGBcolor           = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	ASCII              = "^[\x00-\x7F]+$"
	Base64             = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	PrintableASCII     = "^[\x20-\x7E]+$"
	hasLowerCase       = ".*[[:lower:]]"
	hasUpperCase       = ".*[[:upper:]]"
	Passwd             = "^[a-zA-Z0-9\\.\\/]+$" // nolint: gosec
	UserGroupName      = "^[a-z_][a-z0-9_-]*[$]?$"
	UserGroupNameI     = "^[a-zA-z_][a-zA-z0-9_-]*[$]?$"
	HasAlpha           = "[a-zA-Z]+"
	HasNumber          = "[0-9]+"
	UnixFilePermission = "^0[1-7][0-7]{2,3}$"
	SetNameIndex       = `\[[0-9]+\]$`
	UUID               = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	UUIDv4             = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[8,9,a,b][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`
)

// nolint: gochecknoglobals
var (
	rxEmail              = regexp.MustCompile(Email)
	rxAlpha              = regexp.MustCompile(Alpha)
	rxAlphanumeric       = regexp.MustCompile(Alphanumeric)
	rxNumeric            = regexp.MustCompile(Numeric)
	rxInteger            = regexp.MustCompile(Integer)
	rxFloat              = regexp.MustCompile(Float)
	rxHexadecimal        = regexp.MustCompile(Hexadecimal)
	rxHexcolor           = regexp.MustCompile(Hexcolor)
	rxRGBcolor           = regexp.MustCompile(RGBcolor)
	rxASCII              = regexp.MustCompile(ASCII)
	rxPrintableASCII     = regexp.MustCompile(PrintableASCII)
	rxBase64             = regexp.MustCompile(Base64)
	rxHasLowerCase       = regexp.MustCompile(hasLowerCase)
	rxHasUpperCase       = regexp.MustCompile(hasUpperCase)
	rxPasswd             = regexp.MustCompile(Passwd)
	rxUserGroupName      = regexp.MustCompile(UserGroupName)
	rxUserGroupNameI     = regexp.MustCompile(UserGroupNameI)
	rxHasAlpha           = regexp.MustCompile(HasAlpha)
	rxHasNumber          = regexp.MustCompile(HasNumber)
	rxUnixFilePermission = regexp.MustCompile(UnixFilePermission)
	rxUUID               = regexp.MustCompile(UUID)
	rxUUIDv4             = regexp.MustCompile(UUIDv4)
	RxSetNameIndex       = regexp.MustCompile(SetNameIndex)
)
