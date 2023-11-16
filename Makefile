.PHONY: build run test

build:
	docker build -t specter-ops-takehome .

run:
	docker run -p 127.0.0.1:3000:3000 specter-ops-takehome

build_manual:
	go build -o ./bin/eratosthenes ./cmd/eratosthenes

run_manual: build_manual
	./bin/eratosthenes -host localhost

test_manual:
	go test -v ./... -coverprofile=.cover.out && go tool cover -func=.cover.out