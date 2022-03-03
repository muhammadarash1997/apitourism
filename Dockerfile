FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o binary-apitourism

EXPOSE 8080

CMD ./binary-apitourism