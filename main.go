package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

// GetRepos takes a username and return their repos
func GetRepos(ctx context.Context, username string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=created&direction=desc", username)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	m := []map[string]interface{}{}
	err = json.NewDecoder(response.Body).Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func init() {
	Client = &http.Client{}
}

func main() {
	ctx := context.Background()
	result, err := GetRepos(ctx, "permadiwibisono")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("result: ", result)
}
