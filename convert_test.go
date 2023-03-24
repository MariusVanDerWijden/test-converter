package main

import "testing"

func TestConvertFile(t *testing.T) {
	if err := convertFile("testdata/add11_ymlFiller.yml", "testdata/add11_yml.py"); err != nil {
		panic(err)
	}
}

func TestConvertFile2(t *testing.T) {
	if err := convertFile("testdata/baseFeeExampleFiller.yml", "testdata/baseFeeExample.py"); err != nil {
		panic(err)
	}
}
