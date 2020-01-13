FROM golang:1.8

WORKDIR /go/src/app/api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["api"]