package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/mwsucahyo/jira-project/api"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read environment variables
	appPort := os.Getenv("APP_PORT")
	jiraURL := os.Getenv("JIRA_URL")
	apiToken := os.Getenv("API_TOKEN")
	email := os.Getenv("EMAIL")

	config := api.Config{
		JiraURL:  jiraURL,
		APIToken: apiToken,
		Email:    email,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Mount API routes
	r.Route("/api/v1", func(r chi.Router) {
		api.Routes(r, config)
	})

	fmt.Println("Server is running on port " + appPort + "...")
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}
