package utils

import (
	"errors"
	"fmt"
	"os"
)

func GetEnvVar(varName string) (string, error) {
	value, exists := os.LookupEnv(varName)

	if !exists {
		return "", errors.New(fmt.Sprintf("Environment variable '%s' is not set.", varName))
	}

	return value, nil
}
