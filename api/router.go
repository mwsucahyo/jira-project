package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router, config Config) {
	// Initialize routes here
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Jira Project API"))
	})

	r.Get("/config", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("JIRA_URL: %s\nAPI_TOKEN: %s\nEMAIL: %s\n", config.JiraURL, config.APIToken, config.Email)))
	})

	r.Get("/get-project/{key}", func(w http.ResponseWriter, r *http.Request) {
		GetProjectByKey(w, r, config, chi.URLParam(r, "key"))
	})

	r.Get("/get-sprint/{sprintID}", func(w http.ResponseWriter, r *http.Request) {
		sprintID := chi.URLParam(r, "sprintID")
		sprintIDInt, err := strconv.Atoi(sprintID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to convert sprintID to int: %v", err), http.StatusInternalServerError)
			return
		}
		GetSprintByID(w, r, config, sprintIDInt)
	})

	r.Get("/get-board/{originBoardID}", func(w http.ResponseWriter, r *http.Request) {
		originBoardID := chi.URLParam(r, "originBoardID")
		originBoardIDInt, err := strconv.Atoi(originBoardID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to convert sprintID to int: %v", err), http.StatusInternalServerError)
			return
		}
		GetBoardByOriginBoardID(w, r, config, originBoardIDInt)
	})

	r.Post("/get-task-additional", func(w http.ResponseWriter, r *http.Request) {
		sprintID := r.FormValue("sprint_id")
		assignee := r.FormValue("assignee")

		sprintIDInt, err := strconv.Atoi(sprintID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to convert sprintID to int: %v", err), http.StatusInternalServerError)
			return
		}

		// Validate input
		if sprintID == "" {
			http.Error(w, "sprintID is required", http.StatusBadRequest)
			return
		}

		if assignee == "" {
			http.Error(w, "assignee is required", http.StatusBadRequest)
			return
		}

		// Handle the request
		GetTaskAdditional(w, r, config, sprintIDInt, assignee)
	})

}
