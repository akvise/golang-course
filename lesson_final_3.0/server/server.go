package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"io"
	"log"
	"main/database"
	"main/entity"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type App struct {
	repo database.Repository
}

func NewApp() *App {
	log.Println("Init application")
	repository := database.InitDb()

	return &App{
		repo: *repository,
	}
}

func (app *App) Run() {
	log.Println("Launching application")

	token := viper.GetString("telegram_token")
	bot, err := tgbotapi.NewBotAPI(token)
	ErrFunc(err)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	ErrFunc(err)

	for update := range updates {
		text := update.Message.Text
		ID := update.Message.Chat.ID
		command := regexp.MustCompile("/[a-z]+").FindString(text)

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
			if len(strings.Fields(text)) == 1 {
				response := tgbotapi.NewMessage(ID, "incorrect request")
				bot.Send(response)
				continue
			}

			city := strings.ReplaceAll(text, "/city ", "")
			resp, _ := http.Get(
				"https://api.openweathermap.org/data/2.5/weather?q=" +
					city + "&appid=" + viper.GetString("weather_token"),
			)
			body, err := io.ReadAll(resp.Body)
			ErrFunc(err)

			var data map[string]interface{}
			json.Unmarshal(body, &data)

			/// parsing.........
			var value string
			if data["cod"] == float64(200) {
				value = "🏙️ " + data["name"].(string) + ", " + data["sys"].(map[string]interface{})["country"].(string) + "\n" +
					"⛅ weather: " + data["weather"].([]interface{})[0].(map[string]interface{})["main"].(string) + ", " +
					data["weather"].([]interface{})[0].(map[string]interface{})["description"].(string) + "\n" +
					"☁ clouds: " + fmt.Sprintf("%.1f", data["clouds"].(map[string]interface{})["all"].(float64)) + "%\n" +
					"🌡 temp: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["temp"].(float64)-273.15) + "°C, " +
					"feels like: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["feels_like"].(float64)-273.15) + "°C\n" +
					"💦 humidity: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["humidity"].(float64)) + "%\n" +
					"🔨 pressure: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["pressure"].(float64)) + "hPa\n" +
					"💨 wind speed: " + fmt.Sprintf("%.1f", data["wind"].(map[string]interface{})["speed"].(float64)) + "m/s"
			} else if data["cod"] == "404" {
				value = "city not found"
			} else {
				value = "something wrong"
			}

			err = database.Set(
				app.repo,
				entity.Response{City: city, Value: value, TimeRequested: time.Now().Format("2006-01-02 15:04:05")},
			)
			ErrFunc(err)

			response := tgbotapi.NewMessage(ID, value)
			bot.Send(response)

		case "/list":
			message, err := database.Get(app.repo)
			ErrFunc(err)
			response := tgbotapi.NewMessage(ID, fmt.Sprint(message))
			bot.Send(response)

		default:
			response := tgbotapi.NewMessage(ID, "unrecognized message")
			bot.Send(response)

		}
	}
}

func ErrFunc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
