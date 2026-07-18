package main

import (
	"log"

	"whywelive.me/learning-go/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
