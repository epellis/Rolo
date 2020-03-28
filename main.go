package main

import (
	"fmt"

	"github.com/epellis/rolo/server"
)

func main() {
	s := server.Default()
	fmt.Println("Finished with message:", s.Run())
}
