FROM golang:1.23-alpine

WORKDIR /app

COPY . .
RUN go build -o stresstest cmd/stresstest/main.go

ENTRYPOINT ["./stresstest"]