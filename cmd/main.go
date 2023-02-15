package main

import (
	"log"
	specialtribble "special-tribble"
)

func main() {
	srv := new(specialtribble.Server)
	err := srv.Run("8000")
	if err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
