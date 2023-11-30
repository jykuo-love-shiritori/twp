package auth

import (
	"sync"

	"github.com/jykuo-love-shiritori/twp/db"
)

type challengeUser struct {
	CodeChallenge   string
	ChallengeMethod codeChallengeMethod
	Username        string
	Role            db.RoleType
}

// simple concurrent-safe in-memory store,
// could usee Redis to create zero side effects
var (
	mu                 sync.RWMutex             = sync.RWMutex{}
	codeChallengePairs map[string]challengeUser = map[string]challengeUser{}
)
