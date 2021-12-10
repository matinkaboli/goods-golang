package utils

import (
	"crypto/sha256"
	"fmt"
)

func HashString(str string) string {
	hashed := sha256.Sum256([]byte(str))
	hashed2 := fmt.Sprintf("%x", hashed)

	return hashed2
}
