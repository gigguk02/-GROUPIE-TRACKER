package main

import (
	"artists/internal/handler"
	"log"
)

func main() {
	if err := handler.Handler(); err != nil {
		log.Fatal(err)
	}
}
