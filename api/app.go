package main

import (
	"log"
	"net/http"
	"os"

	"github.com/h-tachikawa/mechanical-receptionist/api/handler"
	"github.com/joho/godotenv"
)

func main() {
	requiredEnvVars := []string{
		"LINE_NOTIFY_ENDPOINT",
		"LINE_NOTIFY_TOKEN",
		"GCP_PROJECT_ID",
		"FIRESTORE_EMULATOR_HOST",
	}

	log.Print("starting server...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("can not read .env file: %v", err)
	}

	for _, envVarName := range requiredEnvVars {
		if os.Getenv(envVarName) == "" {
			log.Fatalln("Please set all of required environment variables.")
			return
		}
	}

	http.HandleFunc("/notify", handler.NewNotificationHandler().Handle)

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
