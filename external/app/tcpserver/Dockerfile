FROM golang:1.18-alpine as tcpserver

WORKDIR /server/

COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go build -o /main ./tcpserver.go

ENTRYPOINT ["/main"]

EXPOSE 4545