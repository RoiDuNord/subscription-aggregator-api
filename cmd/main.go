package main

import (
	"log"
	"subscription-aggregator-api/app"
)

func main() {
	if err := app.MustStart(); err != nil {
		log.Fatalf("App not started: %v", err)
		return
	}
}
