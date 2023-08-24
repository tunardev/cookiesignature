# Cookie Signature

The package offers the ability to sign and verify cookie values using HMAC and SHA-256. This feature is beneficial when you need to guarantee the accuracy of information retained in cookies.

## Installation

To install the package, you can use the `go get` command:

```bash
go get github.com/tunardev/cookiesignature
```

## Usage

Import the package in your Go code:

```go
package main

import (
	"fmt"
	"github.com/tunardev/cookiesignature"
)

func main() {
	secretKey := []byte("your-secret-key")

	// Sign a cookie value
	originalValue := "user123"
	signedValue, err := cookiesignature.Sign(originalValue, secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Signed Value:", signedValue)

	// Unsign and verify a cookie value
	unsignedValue, err := cookiesignature.Unsign(signedValue, secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Original Value:", unsignedValue)
}
```