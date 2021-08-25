package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http"
)


func ErrFunc(err error){
	if err != nil {
		log.Println(err)
	}
}

func ReadToken(path string, tokenType string) (string, error) {
	var data map[string]string

	file, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err

	}
	if err = json.Unmarshal(file, &data); err != nil {
		return "", err
	}

	if tokenType == "telegram" || tokenType == "weather" {
		return data[tokenType], nil
	} else{
		return "", errors.New("wrong token type")
	}
}

func MakeUrl(city string)(string, error){
	KeyWeather, err := ReadToken("./token.json", "weather")
	ErrFunc(err)

	URL := "https://api.openweathermap.org/data/2.5/weather?q=" +
		city + "&appid=" + KeyWeather

	if _, err = http.Get(URL); err != nil {
		return "", err
	}
	return URL, nil
}
