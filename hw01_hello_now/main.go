package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Could not get ntp time: %s", err)
	}

	fmt.Println("current time: " + time.Now().Round(time.Second).String())
	fmt.Println("exact time: " + ntpTime.Round(time.Second).String())
}
