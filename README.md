# SpecterOps coding assignment

## Background

Prime numbers have many modern day applications and a long history in mathematics. This repository implements a simple REST API allowing clients to retrieve the Nth prime number, assuming a 0th index for the prime number 2.

## Running the application

### Clone

`git clone https://github.com/gjbranham/eratosthenes.git && cd eratosthenes/`

### Build

`go build -o ./bin/nth-prime ./cmd/nth-prime `

### Run server

`./bin/nth-prime`

### Interact

Submit a GET request for the Nth prime number:

`curl http://localhost:3000/prime/1239`

### Test

Run unit tests:

`go test -v ./... -coverprofile=.cover.out && go tool cover -func=.cover.out`

Run fuzzing test (will continue until manually killed):

`go test -v -fuzz=NthPrime internal/sieve/*.go`
