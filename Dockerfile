FROM golang:1.17

ENV Environment "Deployment"

RUN mkdir server
WORKDIR /server

COPY main.go /server/
COPY go.mod /server/

RUN go mod download

EXPOSE 8080

CMD go run main.go
