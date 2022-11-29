package main

import (
	"demo1/api/v1/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../config/local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	r := router.SetupRouter()
	r.Run(":3005")
}
