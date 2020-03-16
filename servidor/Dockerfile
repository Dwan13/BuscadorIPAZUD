FROM golang:1.13

EXPOSE 3000

WORKDIR /app

RUN go get -u github.com/go-sql-driver/mysql && go get github.com/rs/cors && go get -u github.com/gorilla/mux


COPY main.go /app
CMD go run main.go
