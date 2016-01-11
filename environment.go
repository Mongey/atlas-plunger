package main

import (
	"errors"
	"fmt"
)

var environments = []string{"ci", "chaos", "uat", "staging", "production"}

func PreviousEnvironment(currentEnvironment string) (string, error) {
	index, err := IndexInSlice(currentEnvironment, environments)

	if err != nil {
		return "", err
	}
	if index == 0 {
		return "", errors.New("There is no stage to promote from")
	}

	return environments[index-1], nil
}

func IndexInSlice(a string, list []string) (int, error) {
	for i, b := range list {
		if b == a {
			return i, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Unable to find", a, "in", list))
}
