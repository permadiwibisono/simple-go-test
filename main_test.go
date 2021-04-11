package main

import (
	"context"
	"testing"
)

func TestGitHubCallSuccess(t *testing.T) {
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
	if result[0]["full_name"] != "permadiwibisono/votey-uppy" {
		t.Error("TestGitHubCallSuccess failed, array was not sorted correctly.")
		return
	}
}
func TestGitHubCallFail(t *testing.T) {
	ctx := context.Background()
	_, err := GetRepos(ctx, "permadiwibisonothisusershouldnotexist")
	if err == nil {
		t.Error("TestGitHubCallFail failed.")
		return
	}
}
