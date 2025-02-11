package models

type ProjectsResponse struct {
	Value []Project `json:"value"`
	Count int       `json:"count"`
}
