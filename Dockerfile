FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o main ./src

RUN ls

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

RUN ls -l /root/

RUN chmod +x ./main

CMD ["/root/main"]
