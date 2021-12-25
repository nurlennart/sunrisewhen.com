package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type timezoneapi struct {
	Status           string      `json:"status"`
	Message          string      `json:"message"`
	CountryCode      string      `json:"countryCode"`
	CountryName      string      `json:"countryName"`
	ZoneName         string      `json:"zoneName"`
	Abbreviation     string      `json:"abbreviation"`
	GmtOffset        int         `json:"gmtOffset"`
	Dst              string      `json:"dst"`
	ZoneStart        int         `json:"zoneStart"`
	ZoneEnd          interface{} `json:"zoneEnd"`
	NextAbbreviation interface{} `json:"nextAbbreviation"`
	Timestamp        int         `json:"timestamp"`
	Formatted        string      `json:"formatted"`
}

func timezone(lat, lon float32) timezoneapi {
	response, err := http.Get(fmt.Sprintf("http://api.timezonedb.com/v2.1/get-time-zone?key=&by=position&lat=%f&lng=%f&format=json", lat, lon))
	if err != nil {
		log.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	var responseObject timezoneapi
	json.Unmarshal(responseData, &responseObject)
	return (responseObject)
}
