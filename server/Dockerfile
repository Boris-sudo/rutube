FROM golang:1.22-bullseye as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN go build -o app ./cmd/ColdStart/main.go

FROM ubuntu:22.04
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app .

EXPOSE 8080

CMD ["./app"]