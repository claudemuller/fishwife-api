FROM golang:1.18

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /app/fishwife ./cmd/main.go

ENV PORT=8080

CMD ["/app/fishwife"]
