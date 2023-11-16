FROM golang:1.21

ENV DOCKERIZED Yes

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go test -v ./... -coverprofile=.cover.out && go tool cover -func=.cover.out

RUN go build -o ./bin/eratosthenes ./cmd/eratosthenes

CMD ["./bin/eratosthenes"]

EXPOSE 3000