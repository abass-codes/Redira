package security

import "testing"

func TestValidateURL(t *testing.T) {

	valid := ValidateURL(
		"https://google.com",
	)

	if !valid {
		t.Fatal(
			"expected valid URL",
		)
	}
}

func TestRejectInvalidURL(t *testing.T) {

	valid := ValidateURL(
		"not-a-url",
	)

	if valid {
		t.Fatal(
			"expected invalid URL",
		)
	}
}
