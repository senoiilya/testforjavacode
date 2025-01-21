package main

import (
	"log"

	"github.com/senoiilya/testforjavacode/internal/app"
)

func main() {
	srv := new(app.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}

}
