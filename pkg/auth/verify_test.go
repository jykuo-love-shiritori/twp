package auth

import "testing"

func TestVerifyCodeChallenge(t *testing.T) {
	verifier := "Mjg0MDQ1NjU2MjI3NzY4NTg5MTkzNDg1Nzc0MjYzMzM4NjgxMjYyOA"
	challenge := "Li76VznZu2Eh5kpxJyz8qjdh0Fr78lxPtCBU3tnHDq4"

	success := verifyCodeChallenge(verifier, challengeUser{
		CodeChallenge:   challenge,
		ChallengeMethod: s256,
	})

	if !success {
		t.Fatal("Code challenge failed")
	}
}

func TestValidUsername(t *testing.T) {
	username := "Test1234"

	valid := isValidUsername(username)

	if !valid {
		t.Fatal("Valid username failed")
	}
}

func TestInvalidUsername(t *testing.T) {
	username := "123%_wer"

	valid := isValidUsername(username)

	if valid {
		t.Fatal("Invalid username failed")
	}
}

func TestValidPassword(t *testing.T) {
	password := "Secret_p@ssw0rd"

	valid := isValidPassword(password)

	if !valid {
		t.Fatal("Valid password failed")
	}
}

func TestInvalidPasswordNoUpper(t *testing.T) {
	password := "secret_p@ssw0rd"

	valid := isValidPassword(password)

	if valid {
		t.Fatal("Invalid password failed")
	}
}

func TestInvalidPasswordNoSpecial(t *testing.T) {
	password := "Secretpassw0rd"

	valid := isValidPassword(password)

	if valid {
		t.Fatal("Invalid password failed")
	}
}

func TestInvalidPasswordNoNumber(t *testing.T) {
	password := "Secret_password"

	valid := isValidPassword(password)

	if valid {
		t.Fatal("Invalid password failed")
	}
}

func TestInvalidPasswordNoLower(t *testing.T) {
	password := "SECRET_PASSWORD"

	valid := isValidPassword(password)

	if valid {
		t.Fatal("Invalid password failed")
	}
}
