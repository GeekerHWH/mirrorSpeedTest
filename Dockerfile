# First stage: build MST
FROM golang:1.22.0-alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN go build -o mirrorSpeedTest main/main.go

# Final stage: cp executable file into image
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/mirrorSpeedTest .

COPY ./urls.json .

CMD ["./mirrorSpeedTest", "-h"]
