package fizzbuzz

import (
	"errors"
	"net/url"
	"strconv"
)

// ValidateDataFizzbuzz is the data valdation function for Fizzbuzzer.
// It returns a valid Fizzbuzzer if no error occures, an incomplete one with the error otherwise.
func ValidateDataFizzbuzz(values url.Values) (Fizzbuzzer, error) {
	fb := Fizzbuzzer{}
	var err error

	// Validating int1
	int1 := values["int1"]
	if len(int1) < 1 {
		return fb, errors.New("missing parameter int1")
	}
	fb.int1, err = strconv.Atoi(int1[0])
	if err != nil {
		return fb, errors.New("unable to convert " + int1[0] + " to int")
	}

	// Validating int2
	int2 := values["int2"]
	if len(int2) < 1 {
		return fb, errors.New("missing parameter int2")
	}
	fb.int2, err = strconv.Atoi(int2[0])
	if err != nil {
		return fb, errors.New("unable to convert " + int2[0] + " to int")
	}

	// Validating limit
	limit := values["limit"]
	if len(limit) < 1 {
		return fb, errors.New("missing parameter limit")
	}
	fb.limit, err = strconv.Atoi(limit[0])
	if err != nil {
		return fb, errors.New("unable to convert " + limit[0] + " to int")
	}

	// Validating string1
	string1 := values["string1"]
	if len(string1) < 1 {
		return fb, errors.New("missing parameter string1")
	}
	fb.string1 = string1[0]

	// Validating string2
	string2 := values["string2"]
	if len(string2) < 1 {
		return fb, errors.New("missing parameter string2")
	}
	fb.string2 = string2[0]

	return fb, nil
}
