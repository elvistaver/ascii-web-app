FROM golang:1.22.2

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

CMD ["./myapp"]

