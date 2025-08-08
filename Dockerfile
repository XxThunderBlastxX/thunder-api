FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go build -o thunder-api ./cmd/main.go

EXPOSE 4040

CMD ["./thunder-api"]
