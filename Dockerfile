FROM golang:1.18.2-alpine3.16 as builder

RUN apk add --no-cache git

COPY . /go/src/github.com/mhkarimi1383/goBaseAPIProject
WORKDIR /go/src/github.com/mhkarimi1383/goBaseAPIProject

RUN go build -o /goBaseAPIProject

FROM alpine:3.14 as runtime

COPY --from=builder /goBaseAPIProject /goBaseAPIProject
RUN chmod +x /goBaseAPIProject

CMD ["/goBaseAPIProject"]