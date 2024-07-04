package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func getProjects(config Config) ([]Project, error) {
	url := fmt.Sprintf("%s/rest/api/3/project", config.JiraURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.Email, config.APIToken)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func getSprintByID(config Config, sprintID int) (*Sprint, error) {
	url := fmt.Sprintf("%s/rest/agile/1.0/sprint/%d", config.JiraURL, sprintID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.Email, config.APIToken)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sprint := &Sprint{}
	err = json.Unmarshal(body, &sprint)
	if err != nil {
		return nil, err
	}

	return sprint, nil
}
func getBoardByOriginBoardID(config Config, originBoardID int) (*Board, error) {
	url := fmt.Sprintf("%s/rest/agile/1.0/board/%d", config.JiraURL, originBoardID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.Email, config.APIToken)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	board := &Board{}
	err = json.Unmarshal(body, &board)
	if err != nil {
		return nil, err
	}

	return board, nil
}

func getTaskAdditional(config Config, sprintID int, assignee string) (*JQLResponse, error) {
	taskTypes := []string{"Task", "Sub-task", "Subtask"}
	sprint, err := getSprintByID(config, sprintID)
	if err != nil {
		return nil, err
	}

	board, err := getBoardByOriginBoardID(config, sprint.OriginBoardID)
	if err != nil {
		return nil, err
	}

	if sprint == nil {
		return nil, fmt.Errorf("sprint not found")
	}

	startDateParsing, err := time.Parse(time.RFC3339, sprint.StartDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing timestamp start date: %v", err)
	}

	var completedDateParsing time.Time
	if sprint.CompleteDate != "" {
		completedDateParsing, err = time.Parse(time.RFC3339, sprint.CompleteDate)
		if err != nil {
			return nil, fmt.Errorf("error parsing timestamp completed date: %v", err)
		}
	} else {
		completedDateParsing, err = time.Parse(time.RFC3339, sprint.StartDate)
		if err != nil {
			return nil, fmt.Errorf("error parsing timestamp completed date: %v", err)
		}
	}

	startDate := startDateParsing.Format(DateFormat)         // Format "2006-01-02" represents "YYYY-MM-DD"
	completedDate := completedDateParsing.Format(DateFormat) // Format "2006-01-02" represents "YYYY-MM-DD"

	jql := fmt.Sprintf(`project = "%s" AND created >= "%s" AND created <= "%s" AND sprint = %d AND assignee = '%s' AND type IN (%s)`,
		board.Location.ProjectName, startDate, completedDate, sprintID, assignee, strings.Join(taskTypes, ","))

	fmt.Println("jql", jql)

	encodedJQL := url.QueryEscape(jql)

	// Build URL using fmt.Sprintf
	url := fmt.Sprintf("%s/rest/api/3/search?jql=%s&fields=key,summary,customfield_10016,subtasks", config.JiraURL, encodedJQL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.Email, config.APIToken)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// print
	fmt.Println(string(body))

	jqlResponse := &JQLResponse{}
	err = json.Unmarshal(body, &jqlResponse)
	if err != nil {
		return nil, err
	}

	return jqlResponse, nil
}
