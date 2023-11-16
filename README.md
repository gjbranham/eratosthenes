# SpecterOps coding assignment

## Background

Prime numbers have many modern day applications and a long history in mathematics. This repository implements a simple RESTful API allowing clients to retrieve the Nth prime number, assuming a 0th index for the prime number 2.

## Running the application

### Clone

`git clone https://github.com/gjbranham/eratosthenes.git && cd eratosthenes/`

### Build

`go build ./...`

### Test

Run unit tests:

`go test -v ./... -coverprofile=.cover.out && go tool cover -func=.cover.out`

Run fuzzing test (will continue until manually killed):

`go test -v -fuzz=NthPrime internal/sieve/*.go`
