package users

import (
	"crypto/rand"
	"encoding/base64"
)

// genRandBytes generates a 32 bytes long string of random bytes.
func genRandBytes() []byte {
	b := make([]byte, 24)
	if _, err := rand.Read(b); err != nil {
		// Panic if we can't generate a random number, since this typically
		// indicates an operating system failure
		panic(err)
	}
	return []byte(base64.URLEncoding.EncodeToString(b))
}
