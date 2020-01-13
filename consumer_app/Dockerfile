FROM golang:1.8

WORKDIR /go/src/app/consumer_app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["consumer_app"]