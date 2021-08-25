package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var responseMessage string

func main() {
	token, err := ReadToken("./token.json", "telegram")
	ErrFunc(err)

	bot, err := tgbotapi.NewBotAPI(token)
	ErrFunc(err)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	ErrFunc(err)


	for update := range updates {
		command := regexp.MustCompile("/[a-z]+").FindString(update.Message.Text)

		switch command {
		case "/start":
			message := "WeatherBot that use API `https://openweathermap.org/`\n" +
				" you can use `/help` for Weather Bot"
			response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			bot.Send(response)

		case "/help":
			message := "You can use / for commands\n" +
				"`/city` [Name] - get weather by cityName\n" +
				"`/city` [ID] - get weather by cityID \n(http://bulk.openweathermap.org/sample/)\n" +
				"`/list` - get list of requests"
			response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			bot.Send(response)

		case "/city":
			city := strings.Fields(update.Message.Text)
			if len(city) == 1{
				response := tgbotapi.NewMessage(update.Message.Chat.ID, "incorrect request")
				bot.Send(response)
				continue
			}
			url, _ := MakeUrl(city[1])
			resp, _ := http.Get(url)
			body, err := io.ReadAll(resp.Body)
			ErrFunc(err)

			var data map[string]interface{}
			json.Unmarshal(body, &data)

			/// parsing.........
			var ans string
			if data["cod"] == float64(200) {
				ans = data["name"].(string) + ", " + data["sys"].(map[string]interface{})["country"].(string) + "\n" +
					"weather: " + data["weather"].([]interface{})[0].(map[string]interface{})["main"].(string) + ", " +
					data["weather"].([]interface{})[0].(map[string]interface{})["description"].(string) + "\n" +
					"clouds: " + fmt.Sprintf("%.2f", data["clouds"].(map[string]interface{})["all"].(float64)) + "%\n"
			}else if data["cod"] == "404" {
				ans = "city not found"
			}else{
				ans = "something wrong"
			}
			response := tgbotapi.NewMessage(update.Message.Chat.ID, ans)
			bot.Send(response)

		default:
				response := tgbotapi.NewMessage(update.Message.Chat.ID, "default")
				bot.Send(response)

		}

	}

}