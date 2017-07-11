package main

import (
	"fmt"

	"github.com/miguellgt/goprojects/users"
)

func main() {
	username, password := "miguellgt", "01_miguellgt_10"

	if err := users.NewUser(username, password); err != nil {
		fmt.Printf("Couldn't create user: %s\n", err.Error())
		return
	}

	if err := users.AuthenticateUser(username, password); err != nil {
		fmt.Printf("Couldn't authenticate user: %s", err.Error())
		return
	}
	fmt.Printf("Successfully created  and authenticated user %s", username)
}
