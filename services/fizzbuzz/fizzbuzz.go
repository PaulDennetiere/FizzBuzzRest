package fizzbuzz

import (
	"errors"
	"strconv"
)

// Fizzbuzzer holds parameters and performs the fizzbuzz.
type Fizzbuzzer struct {
	int1    int
	int2    int
	limit   int
	string1 string
	string2 string
}

// Fizzbuzz is the Fizzbuzzer's method that performs the fizzbuzz.
// In order to avoid "strange" fizzbuzz, some rules are added :
//  - limit must be greater or equal to 1
//  - int1 and int2 must be different
//  - string1 and string2 can't be empty
//  - string1 and string2 must be different
// Those rules are subjectives decision.
func (f *Fizzbuzzer) Fizzbuzz() ([]string, error) {
	result := []string{}
	if f.limit < 1 {
		return result, errors.New("limit must be greater or equal to 1")
	}
	if f.int1 <= 0 || f.int2 <= 0 {
		return result, errors.New("int1 and int2 must be strictly positive")
	}
	if f.int1 == f.int2 {
		return result, errors.New("int1 and int2 must be different")
	}
	if f.string1 == "" || f.string2 == "" {
		return result, errors.New("string1 and string2 can't be empty")
	}
	if f.string1 == f.string2 {
		return result, errors.New("string1 and string2 must be different")
	}
	for i := 1; i <= f.limit; i++ {
		if i%(f.int1*f.int2) == 0 {
			result = append(result, f.string1+f.string2)
		} else if i%f.int2 == 0 {
			result = append(result, f.string2)
		} else if i%f.int1 == 0 {
			result = append(result, f.string1)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result, nil
}
