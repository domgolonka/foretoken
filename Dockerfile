FROM golang:1.15-alpine AS builder

RUN apk add bash ca-certificates git libxml2-dev pkgconfig

RUN mkdir /app
WORKDIR /app
ENV GO111MODULE=on

COPY . .
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o foretoken cmd/main.go

# Run container
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/foretoken .
COPY --from=builder /app/infra/config/config.stag.yml .

ENTRYPOINT ["./foretoken", "--conf=config.stag.yml"]

