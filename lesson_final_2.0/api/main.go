package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	var conn net.Conn

	/// connect with telegram bot
	token, err := ReadToken("./telegram_token.json")
	ErrFunc(err)

	bot, err := tgbotapi.NewBotAPI(token)
	ErrFunc(err)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	ErrFunc(err)

	for update := range updates {
		if update.Message != nil {
			conn, _ = net.Dial("tcp", "127.0.0.1:8080")

			var request string

			id := update.Message.Chat.ID

			if len(update.Message.Text) > 100 {
				request = update.Message.Text[:100]
			} else {
				request = update.Message.Text
			}

			/// send request to server
			log.Println(id, "request:", request)
			fmt.Fprintf(conn, request+"\n")

			/// get response and show to user
			message, _ := bufio.NewReader(conn).ReadString('\v')
			response := tgbotapi.NewMessage(id, message)
			bot.Send(response)
		}
	}
	conn.Close()
}

func ReadToken(Path string) (string, error) {

	var data map[string]string
	file, err := ioutil.ReadFile(Path)

	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(file, &data); err != nil {
		return "", err
	}

	return data["token"], nil

}

func ErrFunc(err error){
	if err != nil {
		log.Println(err)
	}
}