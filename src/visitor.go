package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func addHeaders(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "https://sunrisewhen.com")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Content-Type", "application/json")
}

func cachedTimeVisitor(w http.ResponseWriter, r *http.Request) {
	loc := userLocation(r.Header.Get("X-Forwarded-For"))
	timesc, err := r.Cookie("times")
	if err != nil {
		log.Println(err)
	}
	decoded := base64Decode(timesc.Value)
	var timesObject pagedata
	json.Unmarshal([]byte(decoded), &timesObject)
	data := pagedata{
		Sunrise:            timesObject.Sunrise,
		Sunset:             timesObject.Sunset,
		UTC_Sunrise:        timesObject.UTC_Sunrise,
		UTC_Sunset:         timesObject.UTC_Sunset,
		CivilTwilightBegin: timesObject.CivilTwilightBegin,
		CivilTwilightEnd:   timesObject.CivilTwilightEnd,
		Lat:                loc.lat,
		Lon:                loc.lon,
	}
	marshalled, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	addHeaders(w)
	w.Write(marshalled)
}

func cachedLocationVisitor(w http.ResponseWriter, r *http.Request) {
	latc, err := r.Cookie("lat")
	if err != nil {
		log.Println(err)
	}
	lonc, err := r.Cookie("lon")
	if err != nil {
		log.Println(err)
	}
	loc := Location{
		lat: floatFromCookie(latc.Value),
		lon: floatFromCookie(lonc.Value),
	}
	marshalled, err := json.Marshal(preparedPagedata(loc))
	if err != nil {
		log.Println(err)
	}
	expiration := time.Now().Add(6 * 60 * time.Minute)
	decodedTimes := base64.StdEncoding.EncodeToString(marshalled)
	timesCookie := http.Cookie{Name: "times", Value: decodedTimes, Expires: expiration, Domain: "sunrisewhen.com", Secure: true}
	http.SetCookie(w, &timesCookie)
	addHeaders(w)
	w.Write(marshalled)
}

func emptyVisitor(w http.ResponseWriter, r *http.Request) {
	loc := userLocation(r.Header.Get("X-Forwarded-For"))
	expiration := time.Now().Add(6 * 60 * time.Minute)
	data := preparedPagedata(loc)
	marshalled, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	decodedTimes := base64.StdEncoding.EncodeToString(marshalled)
	latCookie := http.Cookie{Name: "lat", Value: fmt.Sprintf("%f", loc.lat), Expires: expiration, Domain: "sunrisewhen.com", Secure: true}
	lonCookie := http.Cookie{Name: "lon", Value: fmt.Sprintf("%f", loc.lon), Expires: expiration, Domain: "sunrisewhen.com", Secure: true}
	timesCookie := http.Cookie{Name: "times", Value: decodedTimes, Expires: expiration, Domain: "sunrisewhen.com", Secure: true}
	http.SetCookie(w, &latCookie)
	http.SetCookie(w, &lonCookie)
	http.SetCookie(w, &timesCookie)
	addHeaders(w)
	w.Write(marshalled)
}

func apiVisitor(w http.ResponseWriter, r *http.Request) {
	loc := userLocation(r.Header.Get("X-Forwarded-For"))
	data := preparedApidata(loc)
	marshalled, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(marshalled)
}
