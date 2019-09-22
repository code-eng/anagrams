package handlers

import (
	"anagrams/server/state"
	"anagrams/utils"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	contentType = "application/x-www-form-urlencoded"

	wordsDict  = `["foobar"]`
	wordsDict1 = `["test"]`

	emptyWordsDict   = `[]`
	emptyPayload     = ``
	incorrectPayload = `{ "foo": "bar" }`
)

var loadDictTests = []struct {
	method         string
	data           string
	expectedStatus int
	bodyIsEmpty    bool
}{
	{
		http.MethodGet,
		"",
		http.StatusMethodNotAllowed,
		false,
	},
	{
		http.MethodPost,
		emptyPayload,
		http.StatusBadRequest,
		false,
	},
	{
		http.MethodPost,
		incorrectPayload,
		http.StatusBadRequest,
		false,
	},
	{
		http.MethodPost,
		wordsDict,
		http.StatusOK,
		true,
	},
	{
		http.MethodPost,
		emptyWordsDict,
		http.StatusOK,
		true,
	},
}

func isEmpty(s string) bool {
	return len(s) == 0
}

func TestLoadDict_Requests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(LoadDict))
	defer server.Close()

	url := server.URL
	client := server.Client()

	for _, testCase := range loadDictTests {
		req, err := http.NewRequest(testCase.method, url, strings.NewReader(testCase.data))
		if err != nil {
			t.Errorf("Cannot create request, err: %v", err)
		}
		req.Header.Set("Content-Type", contentType)

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

		if testCase.bodyIsEmpty && !isEmpty(string(body)) {
			t.Errorf("Body should be empty, actual body: %v", body)
		} else if !testCase.bodyIsEmpty && isEmpty(string(body)) {
			t.Error("Body should not be empty")
		}
	}

}

func TestLoadDict_DictIsLoaded(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(LoadDict))
	defer server.Close()

	url := server.URL
	client := server.Client()

	// suppressing errors, because we have other test suite for this purpose
	_, _ = client.Post(url, contentType, strings.NewReader(wordsDict))
	expected := []string{"foobar"}
	result := state.GetAnagrams("foobar")
	if !utils.CompareStringSlices(expected, result) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, result)
	}

	_, _ = client.Post(url, contentType, strings.NewReader(wordsDict1))
	expected = []string{}
	result = state.GetAnagrams("foobar")
	if !utils.CompareStringSlices(expected, result) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, result)
	}

	expected = []string{"test"}
	result = state.GetAnagrams("test")
	if !utils.CompareStringSlices(expected, result) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, result)
	}
}

type badReader int

func (badReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestLoadDict_ReadBodyError(t *testing.T) {
	badRequest := httptest.NewRequest(http.MethodPost, "/load", badReader(0))
	w := httptest.NewRecorder()
	LoadDict(w, badRequest)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %v, got: %v", http.StatusInternalServerError, res.Status)
	}

	if isEmpty(string(body)) {
		t.Error("Body should not be empty")
	}
}
