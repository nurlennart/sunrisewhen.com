package main

import (
	"log"
	"net/http"

	"github.com/oschwald/geoip2-golang"
)

type pagedata struct {
	Sunrise            string
	Sunset             string
	UTC_Sunrise        string
	UTC_Sunset         string
	CivilTwilightBegin string
	CivilTwilightEnd   string
	Lat                float32
	Lon                float32
}

type apidata struct {
	Sunrise string
	Sunset  string
}

type Location struct {
	lat float32
	lon float32
}

var DB, err = geoip2.Open("GeoLite2-City.mmdb")

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/", apiVisitor)
	log.Println("up and running. - 127.0.0.1:49151")
	http.ListenAndServe("127.0.0.1:49151", nil)
	defer DB.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookieExistent(w, r, "times") {
			cachedTimeVisitor(w, r)
		} else {
			if cookieExistent(w, r, "lat") && cookieExistent(w, r, "lon") {
				cachedLocationVisitor(w, r)
			} else {
				emptyVisitor(w, r)
			}
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func preparedPagedata(loc Location) pagedata {
	suntimes := Sunapi(loc.lat, loc.lon)
	timezone := timezone(loc.lat, loc.lon)
	data := pagedata{
		Sunrise:            Local(suntimes.Results.Sunrise, timezone.ZoneName),
		Sunset:             Local(suntimes.Results.Sunset, timezone.ZoneName),
		UTC_Sunrise:        UTC(suntimes.Results.Sunrise),
		UTC_Sunset:         UTC(suntimes.Results.Sunset),
		CivilTwilightBegin: Local(suntimes.Results.CivilTwilightBegin, timezone.ZoneName),
		CivilTwilightEnd:   Local(suntimes.Results.CivilTwilightEnd, timezone.ZoneName),
		Lat:                loc.lat,
		Lon:                loc.lon,
	}
	return data
}

func preparedApidata(loc Location) apidata {
	suntimes := Sunapi(loc.lat, loc.lon)
	timezone := timezone(loc.lat, loc.lon)
	data := apidata{
		Sunrise: Api(suntimes.Results.Sunrise, timezone.ZoneName),
		Sunset:  Api(suntimes.Results.Sunset, timezone.ZoneName),
	}
	return data
}
