from golang:1.16

ENV GO111MODULE=off

WORKDIR /var/www/html/app

COPY main.go .

run go build -o api .

EXPOSE 8081

CMD ["./api"]