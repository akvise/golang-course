FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get "github.com/go-telegram-bot-api/telegram-bot-api"
RUN go get "go.mongodb.org/mongo-driver/bson"
RUN go get "go.mongodb.org/mongo-driver/mongo"
RUN go get "go.mongodb.org/mongo-driver/options"

CMD ["go run ."]