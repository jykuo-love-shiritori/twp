package auth

import "testing"

func TestVerifyCodeChallenge(t *testing.T) {
	verifier := "Mjg0MDQ1NjU2MjI3NzY4NTg5MTkzNDg1Nzc0MjYzMzM4NjgxMjYyOA"
	success := verifyCodeChallenge(verifier, challengeUser{
		CodeChallenge:   "Li76VznZu2Eh5kpxJyz8qjdh0Fr78lxPtCBU3tnHDq4",
		ChallengeMethod: s256,
	})

	if !success {
		t.Fatal("Code challenge failed")
	}
}
