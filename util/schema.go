package util

import (
	"io/ioutil"
)

func GetSchema(path string) (string, error) {
	schema, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(schema), nil
}
