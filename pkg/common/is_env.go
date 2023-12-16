package common

import (
	"os"

	"github.com/jykuo-love-shiritori/twp/pkg/constants"
)

func IsEnv(env constants.Environment) bool {
	return os.Getenv("TWP_ENV") == env.String()
}
