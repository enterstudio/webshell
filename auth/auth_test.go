package auth

import (
	"bitbucket.org/taruti/pbkdf2"
	"fmt"
	"testing"
)

const test_password = "this is a test password"

func TestSetup(t *testing.T) {
	LookupCredentials = testAuthProvider
}

func testAuthProvider(user interface{}) (salt, hash []byte) {
	hash = []byte{}
	salt = []byte{}
	switch user.(type) {
	default:
		return
	case string:
		break
	}

	if user != "user" {
		return
	}
	ph := pbkdf2.HashPassword(test_password)
	hash = ph.Hash
	salt = ph.Salt
	return
}

func TestAuthentication(t *testing.T) {
	fmt.Printf("[+] testing user authentication: ")

	if !Authenticate("user", test_password) {
		fmt.Println("failed")
		t.FailNow()
	} else if Authenticate("user", "bad password") {
		fmt.Println("failed")
		t.FailNow()
	} else if Authenticate("user", "") {
		fmt.Println("failed")
		t.FailNow()
	} else if Authenticate("", "") {
		fmt.Println("failed")
		t.FailNow()
	} else if Authenticate("eve", test_password) {
		fmt.Println("failed")
		t.FailNow()
	}
}

func BenchmarkAuthenticateSuccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Authenticate("user", test_password)
	}
}

func BenchmarkAuthenticateFailure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Authenticate("user", "bad pass")
	}
}

