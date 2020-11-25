package main

import (
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Could not get ntp time: %s", err)
	}

	println("current time: " + time.Now().Round(time.Second).String())
	println("exact time: " + ntpTime.Round(time.Second).String())
}
