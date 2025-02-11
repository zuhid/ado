package main

import (
	"testing"

	"github.com/h2non/gock"
)

func TestGetFixedValue(t *testing.T) {
	defer gock.Off()

	gock.New("https://dev.azure.com/tzather/_apis/").Get("/projects").
		Reply(200).
		JSON(`{
    "count": 1,
    "value": [
        {
            "id": "bbebc0d4-2a8d-4509-89e8-f8e89fd499df",
            "name": "project01",
            "url": "https://dev.azure.com/organization/_apis/projects/bbebc0d4-2a8d-4509-89e8-f8e89fd499df",
            "state": "wellFormed",
            "revision": 101,
            "visibility": "private",
            "lastUpdateTime": "2024-01-11T19:19:19.19Z"
        }
    ]
}`)
	main()
}
