FROM golang:1.13-alpine3.10 as builder

RUN apk --no-cache --no-progress add make git

WORKDIR /go/src/github.com/AubreyHewes/update-dynamic-host
COPY . .
RUN make build

FROM alpine:3.10
RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates

COPY --from=builder /go/src/github.com/AubreyHewes/update-dynamic-host/dist/update-dynamic-host /usr/bin/update-dynamic-host
ENTRYPOINT [ "/usr/bin/update-dynamic-host" ]
