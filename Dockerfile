FROM golang:1.13-alpine3.10 as builder

ARG UPX_VERSION=3.96

RUN apk --no-cache --no-progress add make git curl

RUN curl -Ls https://github.com/upx/upx/releases/download/v${UPX_VERSION}/upx-${UPX_VERSION}-amd64_linux.tar.xz -o upx.tar.xz && \
    tar xf upx.tar.xz && \
    mv upx*/upx /usr/local/bin && \
    rm -rf upx*

WORKDIR /go/src/github.com/AubreyHewes/go-dyndns
COPY . .

RUN make dist-cli

FROM alpine:3.13
RUN apk update \
    && apk add --no-cache --no-progress ca-certificates tzdata \
    && update-ca-certificates

COPY --from=builder /go/src/github.com/AubreyHewes/go-dyndns/dist/go-dyndns /usr/bin/go-dyndns

RUN adduser -h /go-dyndns -D go-dyndns go-dyndns
USER go-dyndns
WORKDIR /go-dyndns

ENTRYPOINT [ "/usr/bin/go-dyndns" ]
