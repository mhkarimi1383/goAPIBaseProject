FROM golang:1.18.2-alpine3.16 as builder

RUN apk add --no-cache git

COPY . /go/src/github.com/mhkarimi1383/goBaseAPIProject
WORKDIR /go/src/github.com/mhkarimi1383/goBaseAPIProject

## we have vendor directory in our project no need to get packages again
# RUN go get -v ./...

RUN go build -o /goBaseAPIProject

FROM alpine:3.14 as runtime

## copy and prepare binary file
COPY --from=builder /goBaseAPIProject /app/goBaseAPIProject
RUN chmod +x /app/goBaseAPIProject

## copy static files
COPY openapi.json /app/openapi.json

CMD ["/app/goBaseAPIProject"]