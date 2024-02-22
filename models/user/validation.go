package user

import (
	"errors"
	"os"
)

func userModelValidation() error {
	if _, ok := os.LookupEnv("DATABASE_ENGINE"); !ok {
		return errors.New("there is no DATABASE_ENGINE environment")
	}
	return nil
}
