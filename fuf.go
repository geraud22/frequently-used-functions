package fuf

import (
	"encoding/json"
	"os"
)

func LoadJSONFIle(filename string, output interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &output); err != nil {
		return err
	}
	return nil
}

func WriteJSONFile(filename string, input interface{}) error {
	data, err := json.MarshalIndent(input, " ", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return err
	}
	return nil
}
