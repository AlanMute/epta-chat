FROM golang:1.22.4

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get -y install postgresql-client && rm -rf /var/lib/apt/lists/* \
    && chmod +x wait-for-postgres.sh \
    && go build -o eptanit ./cmd/main.go

EXPOSE 8080

CMD ["./eptanit"]