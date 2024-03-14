FROM golang:latest

WORKDIR /app

RUN curl -L -o pkl https://github.com/apple/pkl/releases/download/0.25.2/pkl-linux-aarch64 && chmod +x pkl && mv pkl /usr/local/bin/

COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go build -o thunder-api ./cmd/main.go

EXPOSE 4040

CMD ["./thunder-api"]