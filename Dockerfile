FROM golang:latest

WORKDIR /app/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

EXPOSE 3000

CMD ["./main"]
