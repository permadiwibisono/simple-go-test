package main

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockDo func(req *http.Request) (*http.Response, error)

type MockClient struct {
	MockDo MockDo
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestGitHubCallSuccess(t *testing.T) {
	// build our response JSON
	jsonResponse := `[{
		"full_name": "mock-repo"
	}]`

	// ceate a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	ctx := context.Background()
	result, err := GetRepos(ctx, "permadiwibisono")
	if err != nil {
		t.Error("TestGitHubCallSuccess failed.")
		return
	}
	if len(result) == 0 {
		t.Error("TestGitHubCallSuccess failed, array was empty.")
		return
	}
	if result[0]["full_name"] != "mock-repo" {
		t.Error("TestGitHubCallSuccess failed, array was not sorted correctly.")
		return
	}
}
func TestGitHubCallFail(t *testing.T) {
	Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       nil,
			}, errors.New("Mock Error")
		},
	}
	ctx := context.Background()
	_, err := GetRepos(ctx, "permadiwibisonothisusershouldnotexist")
	if err == nil {
		t.Error("TestGitHubCallFail failed.")
		return
	}
}
