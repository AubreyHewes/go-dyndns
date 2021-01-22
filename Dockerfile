FROM golang:1.13-alpine3.10 as builder

RUN apk --no-cache --no-progress add make git

WORKDIR /go/src/github.com/AubreyHewes/go-dyndns
COPY . .
RUN make dist

FROM alpine:3.10
RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates

COPY --from=builder /go/src/github.com/AubreyHewes/go-dyndns/dist/dyndns /usr/bin/dyndns

ENTRYPOINT [ "/usr/bin/dyndns" ]
