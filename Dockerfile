FROM golang:1.18 as builder
RUN mkdir /resources
ADD . /resources/
WORKDIR /resources

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go mod download

RUN go build -a -installsuffix cgo -o service cmd/api/main.go

FROM alpine:latest as cacerts
RUN apk --no-cache add ca-certificates

FROM scratch
COPY --from=cacerts  /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /resources/config /config
COPY --from=builder /resources/service .
EXPOSE 8080
CMD ["./service"]