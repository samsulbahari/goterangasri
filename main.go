package main

import (
	"fmt"
	"goterangasri/config"
	"goterangasri/router"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	_, db_error := config.SetDatabase()
	if err != db_error {
		log.Fatal("db connection error")
	}
	fmt.Println("database connect")
	router.Setuprouter()

}
