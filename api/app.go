package main

import (
	"log"
	"net/http"
	"os"

	"github.com/h-tachikawa/mechanical-receptionist/api/handler"

	"github.com/joho/godotenv"
)

func main() {
	log.Print("starting server...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("can not read .env file: %v", err)
	}

	if os.Getenv("LINE_NOTIFY_ENDPOINT") == "" || os.Getenv("LINE_NOTIFY_TOKEN") == "" {
		log.Fatalln("Please set all of required environment variables.")
		return
	}

	http.HandleFunc("/notify", handler.NotifyHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
