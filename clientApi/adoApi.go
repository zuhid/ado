package clientApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"zuhid.com/ado/models"
)

func GetProjects(config models.Config) models.ProjectsResponse {
	return get[models.ProjectsResponse](config.AdoToken, config.AdoApi+`projects`)
}

func get[T any](token string, url string) T {
	// create request
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error making request:", err)
	}

	// unmarshal response
	model := new(T)
	err = json.Unmarshal(body, &model)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
	}
	return *model
}
