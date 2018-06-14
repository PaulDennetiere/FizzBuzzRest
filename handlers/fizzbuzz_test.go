package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFizzbuzzHandler(t *testing.T) {
	fb := Fizzbuzz{}
	testCases := []struct {
		params             map[string][]string
		expected           string
		expectedStatusCode int
	}{
		// Checking error at validateData and status code 400
		{map[string][]string{"int1": {"c"}, "int2": {"5"}, "string1": {"fizz"}, "string2": {"buzz"}, "limit": {"10"}}, "{\"error\":\"unable to convert c to int\"}", 400},

		// Checking error for fizzbuzz's rules and status code 400
		{map[string][]string{"int1": {"3"}, "int2": {"5"}, "string1": {"fizz"}, "string2": {"fizz"}, "limit": {"10"}}, "{\"error\":\"string1 and string2 must be different\"}", 400},

		// Checking correct parameters
		{map[string][]string{"int1": {"3"}, "int2": {"5"}, "string1": {"fizz"}, "string2": {"buzz"}, "limit": {"10"}}, `{"data":["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz"]}`, 200},
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
		resp := httptest.NewRecorder()
		fb.ServeHTTP(resp, req)
		if resp.Code != test.expectedStatusCode {
			t.Error("wrong status code")
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			t.Fail()
			return
		}

		if string(body) != test.expected {
			t.Error("wrong content")
		}
	}

}
