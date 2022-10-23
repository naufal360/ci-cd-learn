FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /docker-go-praktikum

EXPOSE 8000

CMD ["/docker-go-praktikum"]