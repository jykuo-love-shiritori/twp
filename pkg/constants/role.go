package constants

type Role int

const (
	ADMIN Role = iota
	USER
)

func (r Role) String() string {
	switch r {
	case ADMIN:
		return "admin"
	case USER:
		return "user"
	default:
		return "unknown"
	}
}
