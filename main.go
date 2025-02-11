package main

import (
	"zuhid.com/ado/clientApi"
	"zuhid.com/ado/db"
	"zuhid.com/ado/models"
)

func main() {
	config, _ := models.LoadConfig("config.json")
	model := clientApi.GetProjects(config)
	db.SaveProjects(config, model)
}
