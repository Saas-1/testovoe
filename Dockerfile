FROM golang:1.23.8

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client



RUN go mod download
RUN go build -o main ./cmd/main.go

EXPOSE 8000

CMD ["./main"]