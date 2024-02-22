package auth

import (
	"errors"
	"os"
)

func endpointValidation() error {
	if _, ok := os.LookupEnv("JWT_SIGNING_KEY"); !ok {
		return errors.New("environment not found: JWT_SIGNING_KEY")
	}
	return nil
}

func init() {
	if err := endpointValidation(); err != nil {
		panic(err)
	}
}
