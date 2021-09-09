package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
)

func main() {

	/// Launching server
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8080")

	for {

		conn, _ := ln.Accept()

		var message string
		/// get request
		request, _ := bufio.NewReader(conn).ReadString('\n')
		request = strings.TrimSpace(request)
		log.Println(request)

		command := regexp.MustCompile("/[a-z]+").FindString(request)

		switch command {
		case "/start":
			message = "WeatherBot that use API `https://openweathermap.org/`\n" +
				" you can use `/help` for Weather Bot"

			/// send response (message)

		case "/help":
			message = "You can use / for commands\n" +
				"`/city` [Name] - get weather by cityName\n" +
				"`/city` [ID] - get weather by cityID \n(http://bulk.openweathermap.org/sample/)\n" +
				"`/list` - get list of requests"

			/// send response (message)

		case "/city":
			city := strings.Fields(request)

			if len(city) == 1 {
				message = "incorrect request: command '/city without city name'"
				/// send response (message)
				break
			}

			cityStr := strings.Replace(request, "/city ", "", 1)
			url, _ := MakeUrl(cityStr)
			resp, _ := http.Get(url)

			body, err := io.ReadAll(resp.Body)
			ErrFunc(err)

			var data map[string]interface{}
			json.Unmarshal(body, &data)

			/// parsing.........
			if data["cod"] == float64(200) {
				message = "🏙️ " + data["name"].(string) + ", " + data["sys"].(map[string]interface{})["country"].(string) + "\n" +
					"⛅ weather: " + data["weather"].([]interface{})[0].(map[string]interface{})["main"].(string) + ", " +
					data["weather"].([]interface{})[0].(map[string]interface{})["description"].(string) + "\n" +
					"☁ clouds: " + fmt.Sprintf("%.1f", data["clouds"].(map[string]interface{})["all"].(float64)) + "%\n" +
					"🌡 temp: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["temp"].(float64)-273.15) + "°C, " +
					"feels like: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["feels_like"].(float64)-273.15) + "°C\n" +
					"💦 humidity: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["humidity"].(float64)) + "%\n" +
					"🔨 pressure: " + fmt.Sprintf("%.1f", data["main"].(map[string]interface{})["pressure"].(float64)) + "hPa\n" +
					"💨 wind speed: " + fmt.Sprintf("%.1f", data["wind"].(map[string]interface{})["speed"].(float64)) + "m/s"
			} else if data["cod"] == "404" {
				message = "city not found"
			} else {
				message = "something wrong with site (weather APi)"
			}

			InitDB()
			err = AddDBRequest(cityStr, message)
			ErrFunc(err)

		case "/list":
			InitDB()
			message, _ = GetList()
		default:
			message = "incorrect request"
		}
		conn.Write([]byte(message + "\v"))

	}

	err := Client.Disconnect(ctx)
	ErrFunc(err)
}


func MakeUrl(city string)(string, error){
	KeyWeather, err := ReadToken("./weather_token.json")
	if err != nil {
		return "", err
	}

	URL := "https://api.openweathermap.org/data/2.5/weather?q=" +
		city + "&appid=" + KeyWeather

	if _, err = http.Get(URL); err != nil {
		return "", err
	}
	return URL, nil
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