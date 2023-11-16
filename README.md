# Background

Prime numbers have many modern day applications and a long history in mathematics. This repository implements a simple REST API allowing clients to retrieve the Nth prime number, assuming a 0th index for the prime number 2.

## Usage

### Clone the repository

`git clone https://github.com/gjbranham/eratosthenes.git && cd eratosthenes/`

### Notes on Docker

Docker is the recommended way to interact with the application. The `build` Makefile target will run all unit tests during Docker image build.

[If you do not have Docker installed](#manual-commands)

#### Build with Docker

`make build`

#### Run server with Docker

`make run`

#### Interact with the API

Submit a GET request for the 1239th prime number:

`curl http://localhost:3000/nthPrime/1239`

#### Fuzz testing

Run fuzzing test (will continue until manually killed with ctrl+c):

`go test -v -fuzz=NthPrime internal/sieve/*.go`

### Manual commands

If Docker is not installed, one may choose to build, run, and test the application manually:

`make test_manual`

`make build_manual`

`make run_manual`