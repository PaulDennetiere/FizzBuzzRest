package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	nf := NotFound{}
	testCases := []struct {
		params             map[string][]string
		expected           string
		expectedStatusCode int
	}{
		// Checking response from bad method
		{map[string][]string{"int1": {"3"}, "int2": {"5"}, "string1": {"fizz"}, "string2": {"buzz"}, "limit": {"10"}}, "{\"error\":\"Bad route\"}", 404},
	}
	for _, test := range testCases {
		req, err := http.NewRequest("GET", "/unknown", nil)
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
		nf.ServeHTTP(resp, req)
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
