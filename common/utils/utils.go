package utils

import "os"

func ReadInputFromFile(fileName string) (string, error) {
	testInput, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(testInput), nil
}
