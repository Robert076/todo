FROM golang:1.23

WORKDIR /app
COPY go.mod .
COPY main/main.go .

RUN go build -o bin .
ENTRYPOINT [ "/app/bin" ]