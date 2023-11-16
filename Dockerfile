FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go test -v ./... -coverprofile=.cover.out && go tool cover -func=.cover.out

RUN go build -o ./bin/eratosthenes ./cmd/eratosthenes

CMD ["./bin/eratosthenes", "-host", "0.0.0.0", "-port", "3000"]

EXPOSE 3000