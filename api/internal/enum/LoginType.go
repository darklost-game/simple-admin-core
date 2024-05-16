package enum

// LoginType is a custom type for login methods.
type LoginType int

const (
	LOGINTYPE_UNKNOWN LoginType = iota
	LOGINTYPE_USERNAME
	LOGINTYPE_SMS
	LOGINTYPE_EMAIL
	LOGINTYPE_OAUTH
)

// String returns the string representation of the LoginType.
func (lt LoginType) String() string {
	return [...]string{"UNKNOWN", "USERNAME", "SMS", "EMAIL", "OAUTH"}[lt]
}

// IsValid checks if the LoginType is valid.
func (lt LoginType) IsValid() bool {
	switch lt {
	case LOGINTYPE_USERNAME, LOGINTYPE_SMS, LOGINTYPE_EMAIL, LOGINTYPE_OAUTH:
		return true
	default:
		return false
	}
}

// FromString converts a string to a LoginType.
func FromString(s string) LoginType {
	switch s {
	case "USERNAME":
		return LOGINTYPE_USERNAME
	case "SMS":
		return LOGINTYPE_SMS
	case "EMAIL":
		return LOGINTYPE_EMAIL
	case "OAUTH":
		return LOGINTYPE_OAUTH
	default:
		return LOGINTYPE_UNKNOWN
	}
}
