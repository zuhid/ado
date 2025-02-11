package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"zuhid.com/ado/models"
)

func SaveProjects(config models.Config, projectsResponse models.ProjectsResponse) {
	db := getDb(config)

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
		}
	}
	db.Close()
	fmt.Println("Data saved to database successfully")
}

func getDb(config models.Config) *sql.DB {
	// save data to postgres database
	connStr := config.ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}
	return db
}
