package skill

type Skill struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Tags        []string `json:"tags"`
}

type SkillResponse struct {
	Status string `json:"status"`
	Data   Skill  `json:"data"`
}
type SkillsResponse struct {
	Status string  `json:"status"`
	Data   []Skill `json:"data"`
}
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type DeleteResponse = ErrorResponse
