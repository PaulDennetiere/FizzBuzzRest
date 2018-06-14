package fizzbuzz

import (
	"testing"
)

func TestFizzbuzzMethod(t *testing.T) {
	fizzbuzzer := Fizzbuzzer{
		int1:    3,
		int2:    5,
		string1: "fizz",
		string2: "buzz",
		limit:   100,
	}
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz", "31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz", "41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz", "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz", "61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz", "71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz", "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz", "91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"}

	got, err := fizzbuzzer.Fizzbuzz()
	if err != nil {
		t.Error(err.Error())
	}

	// Testing sizes
	if len(got) != len(expected) {
		t.Error("The result does not have the right length.")
		return
	}

	// Testing content
	if !checkStringSliceContentEquality(got, expected) {
		t.Error("The result does not have the expected content.")
		return
	}
}

func TestFizzbuzzMethodRules(t *testing.T) {
	testTable := []struct {
		int1    int
		int2    int
		string1 string
		string2 string
		limit   int
	}{
		// Checking same int1 and int2
		{2, 2, "fizz", "buzz", 10},
		// Checking same string1 and string2
		{3, 5, "fizz", "fizz", 10},
		// Checking one string is empty
		{3, 5, "", "buzz", 10},
		{3, 5, "fizz", "", 10},
		// Checking limit at 0
		{3, 5, "fizz", "buzz", 0},
		// Checking int1 or int2 equal to 0
		{0, 0, "fizz", "buzz", 10},
	}
	// Testing that those set of parameters trigger errors.
	for _, test := range testTable {
		fb := Fizzbuzzer{
			int1:    test.int1,
			int2:    test.int2,
			string1: test.string1,
			string2: test.string2,
			limit:   test.limit,
		}
		_, err := fb.Fizzbuzz()
		if err == nil {
			t.Errorf("Those values should not be acceptable: %d int1, %d int2, %s string1, %s string2, %d limit.", test.int1, test.int2, test.string1, test.string2, test.limit)
		}
	}
}

func checkStringSliceContentEquality(firstSlice, secondSlice []string) bool {
	for i := range firstSlice {
		if firstSlice[i] != secondSlice[i] {
			return false
		}
	}
	return true
}
