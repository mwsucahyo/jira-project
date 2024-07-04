package api

type Config struct {
	JiraURL  string `json:"JIRA_URL"`
	APIToken string `json:"API_TOKEN"`
	Email    string `json:"EMAIL"`
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Sprint struct {
	ID            int    `json:"id"`
	Self          string `json:"self"`
	State         string `json:"state"`
	Name          string `json:"name"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	CompleteDate  string `json:"completeDate"`
	CreatedDate   string `json:"createdDate"`
	OriginBoardID int    `json:"originBoardId"`
	Goal          string `json:"goal"`
}

type Board struct {
	ID       int           `json:"id"`
	Self     string        `json:"self"`
	Name     string        `json:"name"`
	Type     string        `json:"type"`
	Location BoardLocation `json:"location"`
}

type BoardLocation struct {
	ProjectID    int    `json:"projectId"`
	DisplayName  string `json:"displayName"`
	ProjectName  string `json:"projectName"`
	ProjectKey   string `json:"projectKey"`
	ProjectType  string `json:"projectTypeKey"`
	AvatarURI    string `json:"avatarURI"`
	ProjectName2 string `json:"name"`
}

type JQLResponse struct {
	Expand     string  `json:"expand"`
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

type Issue struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Summary           string    `json:"summary"`
	Subtasks          []Subtask `json:"subtasks"`
	Customfield_10016 float32   `json:"customfield_10016"`
	Customfield_10024 float32   `json:"customfield_10024"`
}

type Subtask struct {
	ID     string        `json:"id"`
	Key    string        `json:"key"`
	Self   string        `json:"self"`
	Fields SubtaskFields `json:"fields"`
}

type SubtaskFields struct {
	Summary   string    `json:"summary"`
	Status    Status    `json:"status"`
	Priority  Priority  `json:"priority"`
	IssueType IssueType `json:"issuetype"`
}

type Status struct {
	Self           string         `json:"self"`
	Description    string         `json:"description"`
	IconURL        string         `json:"iconUrl"`
	Name           string         `json:"name"`
	ID             string         `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
}

type StatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}

type Priority struct {
	Self    string `json:"self"`
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

type IssueType struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	Name           string `json:"name"`
	Subtask        bool   `json:"subtask"`
	AvatarID       int    `json:"avatarId"`
	EntityID       string `json:"entityId"`
	HierarchyLevel int    `json:"hierarchyLevel"`
}

type IssueTask struct {
	Summary     string  `json:"summary"`
	JiraLink    string  `json:"jiraLink"`
	StoryPoints float32 `json:"storyPoints"`
}

type ResponseIssue struct {
	StoryPointsTotal float32
	Issues           []IssueTask `json:"issues"`
}
