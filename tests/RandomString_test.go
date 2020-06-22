package tests

import "testing"
import "password-generator/packages"

func TestRandomString(t *testing.T) {
	randomString16 := packages.RandomString(10)
	if len(randomString16) != 10 {
		t.Error("The length of generated string is not 10")
	}

	randomString10 := packages.RandomString(16)
	if len(randomString10) != 16 {
		t.Error("The length of generated string is not 16")
	}
}