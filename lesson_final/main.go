package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

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
		InitDB()
		command := regexp.MustCompile("/[a-z]+").FindString(update.Message.Text)
		ID := update.Message.Chat.ID
		err = AddDBRequest(ID, update.Message.Text)
		ErrFunc(err)
		log.Println("ID:", ID, "  message:", update.Message.Text)

		switch command {
		case "/start":
			message := "WeatherBot that use API `https://openweathermap.org/`\n" +
				" you can use `/help` for Weather Bot"
			response := tgbotapi.NewMessage(ID, message)

			bot.Send(response)

		case "/help":
			message := "You can use / for commands\n" +
				"`/city` [Name] - get weather by cityName\n" +
				"`/city` [ID] - get weather by cityID \n(http://bulk.openweathermap.org/sample/)\n" +
				"`/list` - get list of requests"
			response := tgbotapi.NewMessage(ID, message)
			bot.Send(response)

		case "/city":
			text := update.Message.Text
			city := strings.Fields(text)
			if len(city) == 1{
				response := tgbotapi.NewMessage(ID, "incorrect request")
				bot.Send(response)
				continue
			}

			cityStr := strings.Replace(text, "/city ", "", 1)
			url, _ := MakeUrl(cityStr)
			resp, _ := http.Get(url)
			body, err := io.ReadAll(resp.Body)
			ErrFunc(err)

			var data map[string]interface{}
			json.Unmarshal(body, &data)

			/// parsing.........
			var ans string
			if data["cod"] == float64(200) {
				ans = "🏙️ " + data["name"].(string) + ", " + data["sys"].(map[string]interface{})["country"].(string) + "\n" +
					"⛅ weather: " + data["weather"].([]interface{})[0].(map[string]interface{})["main"].(string) + ", " +
					data["weather"].([]interface{})[0].(map[string]interface{})["description"].(string) + "\n" +
					"☁ clouds: " + fmt.Sprintf("%.1f", data["clouds"].(map[string]interface{})["all"].(float64)) + "%\n" +
					"🌡 temp: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["temp"].(float64)-273.15) + "°C, " +
					"feels like: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["feels_like"].(float64)-273.15) + "°C\n" +
					"💦 humidity: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["humidity"].(float64)) + "%\n" +
					"🔨 pressure: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["pressure"].(float64)) + "hPa\n" +
					"💨 wind speed: " + fmt.Sprintf("%.1f",data["wind"].(map[string]interface{})["speed"].(float64)) + "m/s"
			}else if data["cod"] == "404" {
				ans = "city not found"
			}else{
				ans = "something wrong"
			}
			response := tgbotapi.NewMessage(ID, ans)
			bot.Send(response)

		case "/list":
			message, _ := GetList()
			response := tgbotapi.NewMessage(ID, fmt.Sprint(message))
			bot.Send(response)
		default:
			response := tgbotapi.NewMessage(ID, "unrecognized message")
			bot.Send(response)
		}

	}

	Client.Disconnect(ctx)
}