package fizzbuzz

import (
	"fmt"
	"net/http"
	"testing"
)

func TestValidateDataFizzbuzz(t *testing.T) {
	testCases := []struct {
		params map[string][]string
	}{
		// Checking error if one of the ints cannot be cast to int.
		{map[string][]string{"int1": {"c"}, "int2": {"5"}, "string1": {"fizz"}, "string2": {"buzz"}, "limit": {"10"}}},
		{map[string][]string{"int1": {"3"}, "int2": {"c"}, "string1": {"fizz"}, "string2": {"buzz"}, "limit": {"10"}}},
		{map[string][]string{"int1": {"3"}, "int2": {"5"}, "string1": {"fizz"}, "string2": {"buzz"}, "limit": {"c"}}},

		// Checking error if one parameter missing.
		{map[string][]string{"int2": {"5"}, "string1": {"fizz"}, "string2": {"buzz"}, "limit": {"10"}}},
		{map[string][]string{"int1": {"3"}, "string1": {"fizz"}, "string2": {"buzz"}, "limit": {"10"}}},
		{map[string][]string{"int1": {"3"}, "int2": {"5"}, "string2": {"buzz"}, "limit": {"10"}}},
		{map[string][]string{"int1": {"3"}, "int2": {"5"}, "string1": {"fizz"}, "limit": {"10"}}},
		{map[string][]string{"int1": {"3"}, "int2": {"5"}, "string1": {"fizz"}, "string2": {"buzz"}}},
	}
	for _, test := range testCases {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			fmt.Println(err.Error())
			t.Fail()
			return
		}
		q := req.URL.Query()
		for k, v := range test.params {
			for _, value := range v {
				q.Add(k, value)
			}
		}
		req.URL.RawQuery = q.Encode()
		if _, err := ValidateDataFizzbuzz(req.URL.Query()); err == nil {
			t.Errorf("this set of parameters should not be validated %v", test.params)
		}
	}

	// Test that a set of correct paramters doesn't trigger any error.
	correctParams := map[string][]string{"int1": {"3"}, "int2": {"5"}, "string1": {"fizz"}, "string2": {"buzz"}, "limit": {"10"}}
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}
	q := req.URL.Query()
	for k, v := range correctParams {
		for _, value := range v {
			q.Add(k, value)
		}
	}
	req.URL.RawQuery = q.Encode()
	if _, err := ValidateDataFizzbuzz(req.URL.Query()); err != nil {
		t.Errorf("this set of parameters should be validated %v", correctParams)
	}
}
