FROM golang:1.19.3-alpine

WORKDIR /app

COPY . .

RUN go build -o soma

CMD ["./soma"]