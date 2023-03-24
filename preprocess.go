package main

import (
	"regexp"
)

var (
	replacements = map[*regexp.Regexp]string{
		regexp.MustCompile(`data:\s*-\s''`): ``, // if data = - `` in transaction, remove it
	}
)

// Preprocess does some replacements on an input code.
// This method is needed since some of the tests are not valid yaml.
func Preprocess(test string) string {
	res := test
	for regexp, replace := range replacements {
		res = regexp.ReplaceAllString(test, replace)
	}
	return res
}
