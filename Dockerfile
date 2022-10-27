FROM golang:1.19 as builder

WORKDIR /usr/src/app

COPY go.* ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /usr/local/bin/app main.go

FROM alpine:3.16

COPY --from=builder /usr/local/bin/app /app

CMD ["/app"]
