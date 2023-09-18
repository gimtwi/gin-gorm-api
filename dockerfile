FROM golang:1.20

ENV GOPROXY=https://proxy.golang.org/
ENV GOCACHE=/go/cache

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o main .

EXPOSE 8888

CMD ["./main"]