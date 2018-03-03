package trackerapi

import (
	"fmt"
	"github.com/soypita/clirescue/user"
)

var (
	urlMe   string     = "https://www.pivotaltracker.com/services/v5/me"
	methodMe string = get
)

func Me() {
	var meResp = new(MeResponse)
	currentUser := autorization()
	parse(makeRequest(methodMe, urlMe), meResp)
	setUserData(meResp, currentUser)
	printUserData(currentUser)
}

func setUserData(response *MeResponse, currentUser *user.User) {
	currentUser.Username = response.Username
	currentUser.Name = response.Name
	currentUser.Email = response.Email
	currentUser.Initials = response.Initials
	currentUser.Timezone.Kind = response.Timezone.Kind
	currentUser.Timezone.Offset = response.Timezone.Offset
	currentUser.Timezone.OlsonName = response.Timezone.OlsonName
	currentUser.APIToken = response.APIToken
}

func printUserData(currentUser *user.User) {
	fmt.Println("Username: ", currentUser.Username)
	fmt.Println("Name: ", currentUser.Name)
	fmt.Println("Email: ", currentUser.Email)
	fmt.Println("Initials: ", currentUser.Initials)
	fmt.Println("Timezone: ", currentUser.Timezone.Kind, currentUser.Timezone.Offset, currentUser.Timezone.OlsonName)
}

type MeResponse struct {
	APIToken string `json:"api_token"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Initials string `json:"initials"`
	Timezone struct {
		Kind      string `json:"kind"`
		Offset    string `json:"offset"`
		OlsonName string `json:"olson_name"`
	} `json:"time_zone"`
}
