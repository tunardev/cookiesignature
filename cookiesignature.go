package cookiesignature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

// Definition of errors
var (
	ErrCookieValueRequired = errors.New("Cookie value must be provided")
	ErrSecretKeyRequired   = errors.New("Secret key must be provided")
	ErrInvalidCookieString = errors.New("Invalid signed cookie string")
	ErrInvalidSignature    = errors.New("Invalid signature")
)

// Signs a given value using the provided secret key and returns the signed value.
func Sign(val string, secret []byte) (string, error) {
	if len(val) == 0 {
		return "", ErrCookieValueRequired
	}
	if len(secret) == 0 {
		return "", ErrSecretKeyRequired
	}

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(val))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return val + "." + signature, nil
}

// Unsigns a signed input using the provided secret key, verifies the signature,
// and returns the original value.
func Unsign(input string, secret []byte) (string, error) {
	if len(input) == 0 {
		return "", ErrInvalidCookieString
	}
	if len(secret) == 0 {
		return "", ErrSecretKeyRequired
	}

	parts := strings.Split(input, ".")
	if len(parts) != 2 {
		return "", ErrInvalidCookieString
	}

	val := parts[0]
	signature := parts[1]

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(val))
	expectedSignature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if signature != expectedSignature {
		return "", ErrInvalidSignature
	}
	return val, nil
}
