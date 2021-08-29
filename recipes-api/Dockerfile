# Dockerfile-golang
FROM golang:1.16

WORKDIR /app

COPY go.* .

# airの設定ファイルをコピー
COPY *air.toml .

RUN go mod download

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]
