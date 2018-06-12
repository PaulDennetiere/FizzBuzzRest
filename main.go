package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		fb, err := validator(values)
		if err != nil {
			w.Write([]byte("{error: unable to FizzBuzz: " + err.Error() + "}"))
			return
		}

		result, err := fizzbuzz(fb.limit, fb.int1, fb.int2, fb.string1, fb.string2)
		if err != nil {
			w.Write([]byte("{error: unable to FizzBuzz: " + err.Error() + "}"))
			return
		}
		buf, err := json.Marshal(result)
		w.Write(buf)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type fizzbuzzer struct {
	int1    int
	int2    int
	limit   int
	string1 string
	string2 string
}

func validator(values url.Values) (fizzbuzzer, error) {
	fb := fizzbuzzer{}
	var err error

	int1 := values["int1"]
	int2 := values["int2"]
	limit := values["limit"]
	string1 := values["string1"]
	string2 := values["string2"]

	if len(int1) < 1 {
		return fb, errors.New("missing parameter int1")
	}
	fb.int1, err = strconv.Atoi(int1[0])

	if err != nil {
		return fb, err
	}
	if len(int2) < 1 {
		return fb, errors.New("missing parameter int2")
	}
	fb.int2, err = strconv.Atoi(int2[0])

	if err != nil {
		return fb, err
	}
	if len(limit) < 1 {
		return fb, errors.New("missing parameter limit")
	}
	fb.limit, err = strconv.Atoi(limit[0])

	if err != nil {
		return fb, err
	}
	if len(string1) < 1 {
		return fb, errors.New("missing parameter string1")
	}
	fb.string1 = string1[0]

	if len(string2) < 1 {
		return fb, errors.New("missing parameter string2")
	}
	fb.string2 = string2[0]

	return fb, nil
}

func fizzbuzz(limit, fizzer, buzzer int, fizz, buzz string) ([]string, error) {
	result := []string{}
	if limit < 1 {
		return result, errors.New("limit must be greater or equal to 1")
	}
	if fizzer == buzzer {
		return result, errors.New("the two integers (fizzer and buzzer) must be different")
	}
	if fizz == "" || buzz == "" {
		return result, errors.New("the two strings (fizz and buzz) can't be empty")
	}
	if fizz == buzz {
		return result, errors.New("the two strings (fizz and buzz) must be different")
	}
	for i := 1; i < limit; i++ {
		if i%(fizzer*buzzer) == 0 {
			result = append(result, fizz+buzz)
		} else if i%buzzer == 0 {
			result = append(result, buzz)
		} else if i%fizzer == 0 {
			result = append(result, fizz)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result, nil
}
