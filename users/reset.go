package users

import (
	"crypto/rand"
	"encoding/base64"
	"net/smtp"
	"os"
)

// validTokens is just an array of currently valid tokens.
var validTokens []string

const passwordForm = `
<h1>Enter you new password</h1>
<form action="/reset" method="POST">
	<input type="hidden" name="email" value="{{ . }}" required/>
	<label for="password">Password</label>
	<input type="password" name="password" required/>
	<input type="submit" value="Submit" />
</form>
`

// SendPasswordResetEmail sends a password reset email to the given user.
func SendPasswordResetEmail(email string) error {
	token := string(genRandBytes())
	validTokens = append(validTokens, token)
	resetLink := "http://localhost:3000/reset?user=" + email + "&token=" + token

	username := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_PASSWORD")
	auth := smtp.PlainAuth("smpt.gmail.com:587", auth, username, "smpt.gmail.com")

	return smtp.SendMail("smpt.gmail.com:587", auth, username, []string{email}, []byte(resetLink))
}

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
