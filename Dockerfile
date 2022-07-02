FROM golang:1.18-alpine

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . .

RUN apk add git build-base
RUN go get -d
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o grpc-example .

EXPOSE 8080
CMD ["./grpc-example"]