package main

import (
	"log"

	"github.com/IDarar/test-task/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
