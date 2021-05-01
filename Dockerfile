FROM golang:1.16

WORKDIR /go/src/app

EXPOSE 8090

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go", "session.go", "database.go"]