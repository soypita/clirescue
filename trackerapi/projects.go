package trackerapi

import (
	"fmt"
)

var (
	urlProjects   string     = "https://www.pivotaltracker.com/services/v5/projects"
	methodProjects string = get
)


func Projects() {
  var projecteResponses = make([]ProjectsResponse, 0)
	autorization()
  parse(makeRequest(methodProjects, urlProjects), &projecteResponses)
	for _, project := range projecteResponses {
    printProjectsData(project)
  }
}

func printProjectsData(proj ProjectsResponse) {
  fmt.Println("Id: ", proj.Id)
  fmt.Println("Name: ", proj.Name)
  fmt.Println("Version: ", proj.Version)
  fmt.Println("Type: ", proj.Type)
  fmt.Println("StartAt: ", proj.StartAt)
  fmt.Println("CreatedAt: ", proj.CreatedAt)
  fmt.Println("UpdatedAt: ", proj.UpdatedAt)

}

type ProjectsResponse struct {
	Id int `json:"id"`
	Name     string `json:"name"`
	Version    int `json:"version"`
	Type string `json:"project_type"`
	StartAt string `json:"start_time"`
  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
}
