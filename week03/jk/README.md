### Test Coverage
```
# getting coverage
go test-v -cover

# output the coverage to a file
go test -v -cover -coverprofile cover.out

# generate html file from the coverage output
go tool cover -html=cover.out -o cover.html
```

### Run Test
```
# simple test on the current folder
go test -v 

# run specific tests
go test -v -run "<Regex>"

# run test with all benchmark + memory benchmark
go test -v -bench . -benchmen

# run a specific benchmark
go test -v -bench "<Regex>" -benchmem
```