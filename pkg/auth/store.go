package auth

import "sync"

type challenge struct {
	challengeString string
	challengeMethod codeChallengeMethod
}

// simple concurrent-safe in-memory store,
// could usee Redis to create zero side effects
var (
	mu                 sync.RWMutex         = sync.RWMutex{}
	codeChallengePairs map[string]challenge = map[string]challenge{}
)
