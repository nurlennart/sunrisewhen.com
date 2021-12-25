package main

import (
	"encoding/base64"
	"log"
	"strconv"
	"time"
)

func floatFromCookie(c string) float32 {
	float, err := strconv.ParseFloat(c, 32)
	if err != nil {
		log.Println(err)
	}
	return float32(float)
}

func base64Decode(encoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	decodedString := string(decoded)
	if err != nil {
		log.Println(err)
	}
	return decodedString
}

func Local(UTC, ZoneName string) string {
	loc, err := time.LoadLocation(ZoneName)
	if err != nil {
		log.Println(err)
	}
	parsedUTC, err := time.Parse(time.RFC3339, UTC)
	local := parsedUTC.In(loc).Format("03:04:05 PM")
	if err != nil {
		log.Println(err)
	}
	return local
}

func UTC(UTC string) string {
	parsedUTC, err := time.Parse(time.RFC3339, UTC)
	utc := parsedUTC.Format("03:04:05 PM")
	if err != nil {
		log.Println(err)
	}
	return utc
}

func Api(UTC string, ZoneName string) string {
	loc, err := time.LoadLocation(ZoneName)
	if err != nil {
		log.Println(err)
	}
	parsedUTC, err := time.Parse(time.RFC3339, UTC)
	if err != nil {
		log.Println(err)
	}
	rfc := string(parsedUTC.In(loc).Format("2006-01-02T15:04:05-07:00"))
	return rfc
}
