package auth

import "sync"

// simple concurrent-safe in-memory store,
// could usee Redis to create zero side effects
var (
	mu                 sync.RWMutex      = sync.RWMutex{}
	codeChallengePairs map[string]string = map[string]string{}
)
