FROM golang:latest AS builder

RUN go version

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go build -o iamvisible iamvisible.go

FROM debian:buster-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/iamvisible /

RUN groupadd -g 999 appuser && \
    useradd -r -u 999 -g appuser appuser

EXPOSE 8080

USER 999:999

CMD ["/iamvisible"]