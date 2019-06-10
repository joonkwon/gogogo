### Test Coverage
```
# getting coverage
go test-v -cover

# output the coverage to a file
go test -v -cover -coverprofile cover.out

# generate html file from the coverage output
go tool cover -html=cover.out -o cover.html
```