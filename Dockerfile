FROM golang:latest

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v main
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-arm64.tar.gz | tar xvz
RUN mv migrate.linux-amd64 /usr/bin/migrate
EXPOSE 8080 8080
CMD ["/usr/src/app/main"]