package cookiesignature

import (
	"testing"
)

func TestSignAndUnsign(t *testing.T) {
	secret := []byte("mysecretkey")
	value := "examplevalue"

	signedValue, err := Sign(value, secret)
	if err != nil {
		t.Errorf("Error signing value: %v", err)
	}

	unsignedValue, err := Unsign(signedValue, secret)
	if err != nil {
		t.Errorf("Error unsigning value: %v", err)
	}

	if unsignedValue != value {
		t.Errorf("Unsinged value does not match original value. Expected: %s, Got: %s", value, unsignedValue)
	}
}

func TestUnsignInvalidSignature(t *testing.T) {
	secret := []byte("mysecretkey")
	value := "examplevalue"

	signedValue, _ := Sign(value, secret)
	// Modify the signed value to create an invalid signature
	modifiedValue := signedValue + "tampered"

	_, err := Unsign(modifiedValue, secret)
	if err != ErrInvalidSignature {
		t.Errorf("Expected ErrInvalidSignature, but got: %v", err)
	}
}

func TestUnsignInvalidCookieString(t *testing.T) {
	secret := []byte("mysecretkey")

	_, err := Unsign("", secret)
	if err != ErrInvalidCookieString {
		t.Errorf("Expected ErrInvalidCookieString, but got: %v", err)
	}

	_, err = Unsign("invalidvalue", secret)
	if err != ErrInvalidCookieString {
		t.Errorf("Expected ErrInvalidCookieString, but got: %v", err)
	}

	// Signed value without a dot separator
	_, err = Unsign("onlyavaluesignature", secret)
	if err != ErrInvalidCookieString {
		t.Errorf("Expected ErrInvalidCookieString, but got: %v", err)
	}
}
