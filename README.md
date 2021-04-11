# simple-go-test

Call a github api (get repos by username) and try make unit tests using mock interfaces.
#### go test

```bash
go test -v ./... -coverpkg ./... -coverprofile cover.out
```

#### results

```bash
=== RUN   TestGitHubCallSuccess
--- PASS: TestGitHubCallSuccess (0.00s)
=== RUN   TestGitHubCallFail
--- PASS: TestGitHubCallFail (0.00s)
=== RUN   TestGitHubCallBadJsonFail
--- PASS: TestGitHubCallBadJsonFail (0.00s)
PASS
coverage: 65.0% of statements in ./...
ok      simple-go-test  0.669s  coverage: 65.0% of statements in ./...
```
