FROM golang:1.20-alpine as builder
RUN apk update && apk add --no-cache make git tzdata
ENV USER=containeruser
ENV UID=10001

WORKDIR $GOPATH/src/echo

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o ".bin/release/linux/monitor" main.go;

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /go/src/echo/.bin/release/linux/monitor /go/bin/monitor

USER containeruser:containeruser

EXPOSE 8181
ENTRYPOINT ["/go/bin/monitor"]
