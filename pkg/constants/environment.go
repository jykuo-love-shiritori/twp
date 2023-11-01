package constants

type Environment int

const (
	DEV Environment = iota
	PROD
)

func (e Environment) String() string {
	switch e {
	case DEV:
		return "dev"
	case PROD:
		return "prod"
	default:
		return "unknown"
	}
}
