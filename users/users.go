// Store is very simple in memory Database, that we'll use to store our users.
// It is protected by a read-wrote mutex, so that two goroutines can't modify
// the underlying map at the same time (since maps are not safe for concurrent
// use in Go)

package users

import (
	"errors"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	rwm *sync.RWMutex
	m   map[string]string
}

var ErrUserAlreadyExists = errors.New("User already exists")

var DB = newDB()

// newDB is a convenience method to initalize our in memory DB when our program
// starts.
func newDB() *Store {
	return &Store{
		rwm: &sync.RWMutex{},
		m:   make(map[string]string),
	}
}

// NewUser accepts a username and password and creates a new user in our DB
func NewUser(username, password string) error {
	err := exists(username)
	if err != nil {
		return err
	}

	DB.rwm.Lock()
	defer DB.rwm.Unlock()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	DB.m[username] = string(hashedPassword)
	return nil
}

// AuthenticateUser accepts a username and password, and checks that the given
// password matches the hashed password. It returns nil on success, and an error
// on failuer.
func AuthenticateUser(username, password string) error {
	DB.rwm.RLock()
	defer DB.rwm.RUnlock()

	hashedPassword := DB.m[username]
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}

// exists is an internal utility function for ensuring the username are unique.
func exists(username string) error {
	DB.rwm.RLock()
	defer DB.rwm.RUnlock()

	if DB.m[username] != "" {
		return ErrUserAlreadyExists
	}
	return nil
}
