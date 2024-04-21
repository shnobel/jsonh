package jsonh

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

func GetDataFromJson[T any](fileName string) (*T, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var t T
	err = json.Unmarshal(byteValue, &t)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling data into %T: %w", t, err)
	}
	return &t, nil
}
