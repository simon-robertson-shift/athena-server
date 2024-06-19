package security

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/alexedwards/argon2id"
)

// Generates a random 256-bit token, as a HEX string.
func CreateToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// Hashes a password using Argon2id.
func HashPassword(password string) string {
	hash, _ := argon2id.CreateHash(password, argon2id.DefaultParams)
	return hash
}

// Checks if the provided password matches the hash. The hash must have been
// created using the HashPassword() function.
func VerifyPassword(password string, hash string) bool {
	ok, _ := argon2id.ComparePasswordAndHash(password, hash)
	return ok
}
