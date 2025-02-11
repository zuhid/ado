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
	body := get(config.AdoToken, config.AdoApi+`projects`)
	model := models.ProjectsResponse{}
	err := json.Unmarshal(body, &model)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		// return nil
	}
	return model
}

func get(token string, url string) []byte {
	// create request
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}

	// return response body
	return body
}
