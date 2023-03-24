package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func convertFile(inputFile string, outputFile string) error {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// preprocess to clean the input
	in := Preprocess(string(input))

	converted, err := convertYaml([]byte(in))
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, converted, os.ModePerm)
}

// converts a yaml filler to a python test
func convertYaml(input []byte) ([]byte, error) {
	tests := make(map[string]Test)
	if err := yaml.Unmarshal(input, &tests); err != nil {
		return nil, err
	}
	if len(tests) != 1 {
		return nil, errors.New("can only do one test per file atm")
	}

	for name, test := range tests {
		return fillTest(test, name)
	}

	return nil, nil
}
