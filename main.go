package main

import (
	"log"

	"github.com/epellis/rolo/server"
)

func main() {
	s, err := server.Default()
	if err != nil {
		log.Fatalln("Error setting up server:", err)
	}
	log.Fatalln("Finished with message:", s.Run())
}
