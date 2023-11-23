# syntax=docker/dockerfile:1
FROM golang:1.21-alpine as builder

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY *.go .
RUN go build -o /bin/server

FROM scratch
LABEL org.opencontainers.image.authors="Yann Duval <yann.duval@posteo.de>"
LABEL org.opencontainers.image.source=https://github.com/devndive/docker-helm-github
LABEL org.opencontainers.image.revision=v1.3.0
LABEL org.opencontainers.image.title="Awesome webserver"
LABEL org.opencontainers.image.description="A webserver to demo helm chart deploymentes"

WORKDIR /app

COPY --from=builder /bin/server /bin/server
COPY --from=builder /etc/passwd /etc/passwd

USER appuser

EXPOSE 8080

ENTRYPOINT ["/bin/server"]
