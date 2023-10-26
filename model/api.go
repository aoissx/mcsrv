package model

type VersionsApiResponse struct {
	ProjectID     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
	Error         string   `json:"error"`
}

type PaperBuildsAPIResponse struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Builds      []int  `json:"builds"`
	Error       string `json:"error"`
}
