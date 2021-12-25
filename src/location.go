package main

import (
	"log"
	"net"
)

func userLocation(remoteAddress string) Location {
	ip := net.ParseIP(remoteAddress)
	record, err := DB.City(ip)
	if err != nil {
		log.Println(err)
	}
	Location := Location{
		lat: float32(record.Location.Latitude),
		lon: float32(record.Location.Longitude),
	}
	return Location
}
