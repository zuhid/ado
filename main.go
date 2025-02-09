package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// get json data from Azure DevOps
	url := `https://dev.azure.com/tzather/_apis/projects`
	jsonStr := []byte(`{"key":"value"}`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer ")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	// fmt.Println("Response status:", string(body))

	// convert json data to struct
	type Project struct {
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		Url            string    `json:"url"`
		State          string    `json:"state"`
		Revision       int       `json:"revision"`
		Visibility     string    `json:"visibility"`
		LastUpdateTime time.Time `json:"lastUpdateTime"`
	}

	type ProjectsResponse struct {
		Value []Project `json:"value"`
		Count int       `json:"count"`
	}

	var projectsResponse ProjectsResponse
	err = json.Unmarshal(body, &projectsResponse)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	fmt.Println("Projects:", projectsResponse.Value)

	// save data to postgres database
	connStr := ``
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	for _, project := range projectsResponse.Value {
		_, err := db.Exec(
			`INSERT INTO projects (id, name, url, state, revision, visibility, lastupdatetime) 
			VALUES ($1, $2, $3, $4, $5, $6, $7) 
			ON CONFLICT (id) DO UPDATE 
			SET name = EXCLUDED.name, url = EXCLUDED.url, state = EXCLUDED.state, revision = EXCLUDED.revision, visibility = EXCLUDED.visibility, lastupdatetime = EXCLUDED.lastupdatetime`,
			project.ID, project.Name, project.Url, project.State, project.Revision, project.Visibility, project.LastUpdateTime,
		)
		if err != nil {
			fmt.Println("Error inserting project into database:", err)
			return
		}
	}

	fmt.Println("Data saved to database successfully")
}
