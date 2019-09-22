package handlers

import (
	"anagrams/server/state"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAnagrams_BadRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(GetAnagrams))
	defer server.Close()

	client := server.Client()

	tests := []struct {
		method         string
		query          string
		expectedStatus int
	}{
		{
			http.MethodPost,
			"",
			http.StatusMethodNotAllowed,
		},
		{
			http.MethodGet,
			"a=b",
			http.StatusBadRequest,
		},
		{
			http.MethodGet,
			"",
			http.StatusBadRequest,
		},
	}

	for _, testCase := range tests {
		url := fmt.Sprintf("%v/get=?%v", server.URL, testCase.query)
		req, err := http.NewRequest(testCase.method, url, strings.NewReader(""))
		if err != nil {
			t.Errorf("Cannot create request, err: %v", err)
		}

		res, err := client.Do(req)
		if err != nil {
			t.Errorf("Error sending request: %v", err)
		}

		if res.StatusCode != testCase.expectedStatus {
			t.Errorf("Expected status code %v, got: %v", testCase.expectedStatus, res.Status)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}

		if isEmpty(string(body)) {
			t.Error("Body should not be empty")
		}
	}
}

var sampleDict = []string{"foobar", "aabb", "baba", "boofar", "test"}

func TestGetAnagrams_CorrectRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(GetAnagrams))
	defer server.Close()

	client := server.Client()

	tests := []struct {
		words        []string
		word         string
		expectedBody string
	}{
		{
			[]string{},
			"test",
			`null`,
		},
		{
			sampleDict,
			"test",
			`["test"]`,
		},
		{
			sampleDict,
			"foobar",
			`["foobar","boofar"]`,
		},
	}

	for _, testCase := range tests {
		state.LoadDictionary(testCase.words)

		url := fmt.Sprintf("%v/get?word=%v", server.URL, testCase.word)
		res, err := client.Get(url)
		if err != nil {
			t.Errorf("Cannot send request, err: %v", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("Wrong status code, got: %v", res.Status)
		}

		if res.Header.Get("Content-Type") != "application/json; charset=utf-8" {
			t.Errorf("Wrong Content-Type header, got %v", res.Header.Get("Content-Type"))
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}

		body := strings.Trim(string(b), "\n")

		if strings.Compare(testCase.expectedBody, body) != 0 {
			t.Errorf("\nExpected: \"%v\"\nGot: \"%v\"", testCase.expectedBody, body)
		}
	}
}
