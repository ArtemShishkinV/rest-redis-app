FROM golang:1.18-alpine

WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o /main ./cmd/apiserver/main.go

EXPOSE 8080

ENV host localhost
ENV port :6379

ENTRYPOINT ["/main"]