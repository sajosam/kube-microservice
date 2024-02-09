FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o student .

EXPOSE 8080

CMD ["./student"]
