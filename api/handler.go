package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetProjectByKey(w http.ResponseWriter, r *http.Request, config Config, key string) {
	projects, err := getProjects(config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get projects: %v", err), http.StatusInternalServerError)
		return
	}

	// Apply filter if a key is provided
	if key != "" {
		filteredProjects := []Project{}
		for _, project := range projects {
			if project.Key == key {
				filteredProjects = append(filteredProjects, project)
			}
		}
		projects = filteredProjects
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

func GetSprintByID(w http.ResponseWriter, r *http.Request, config Config, sprintID int) {
	sprint, err := getSprintByID(config, sprintID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get projects: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(sprint)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

func GetBoardByOriginBoardID(w http.ResponseWriter, r *http.Request, config Config, originBoardID int) {
	board, err := getBoardByOriginBoardID(config, originBoardID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get projects: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(board)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

func GetTaskAdditional(w http.ResponseWriter, r *http.Request, config Config, sprintID int, assignee string) {
	tasks, err := getTaskAdditional(config, sprintID, assignee)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get projects: %v", err), http.StatusInternalServerError)
		return
	}

	if tasks == nil {
		http.Error(w, "No tasks found", http.StatusNotFound)
		return
	}

	responseIssues := ResponseIssue{}
	issues := []IssueTask{}

	for _, issue := range tasks.Issues {
		jiraLink := fmt.Sprintf("%s/browse/%s", config.JiraURL, issue.Key)
		storyPoints := issue.Fields.Customfield_10016 + issue.Fields.Customfield_10024
		issues = append(issues, IssueTask{
			Summary:     issue.Fields.Summary,
			JiraLink:    jiraLink,
			StoryPoints: storyPoints,
		})

		responseIssues.StoryPointsTotal += storyPoints
	}

	responseIssues.Issues = issues

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseIssues)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
