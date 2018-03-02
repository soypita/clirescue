package trackerapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	u "os/user"

	"github.com/soypita/clirescue/cmdutil"
	"github.com/soypita/clirescue/user"
)

var (
	URL          string     = "https://www.pivotaltracker.com/services/v5/me"
	FileLocation string     = homeDir() + "/.tracker"
	currentUser  *user.User = user.New()
	Stdout       *os.File   = os.Stdout
)

func Me() {
	dat, err := ioutil.ReadFile(FileLocation)
	check(err)
	if dat != nil {
		currentUser.APIToken = string(dat)
	} else {
		setCredentials()
	}
	parse(makeRequest())
	printUserData()
	ioutil.WriteFile(FileLocation, []byte(currentUser.APIToken), 0644)
}

func makeRequest() []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if currentUser.APIToken == "" {
		req.SetBasicAuth(currentUser.Username, currentUser.Password)
	} else {
		req.Header.Add("X-TrackerToken", currentUser.APIToken)
	}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}
	return body
}

func parse(body []byte) {
	var meResp = new(MeResponse)
	err := json.Unmarshal(body, &meResp)
	if err != nil {
		fmt.Println("error:", err)
	}

	setUserData(meResp)
	currentUser.APIToken = meResp.APIToken
}

func setUserData(response *MeResponse) {
	currentUser.Username = response.Username
	currentUser.Name = response.Name
	currentUser.Email = response.Email
	currentUser.Initials = response.Initials
	currentUser.Timezone.Kind = response.Timezone.Kind
	currentUser.Timezone.Offset = response.Timezone.Offset
	currentUser.Timezone.OlsonName = response.Timezone.OlsonName
	currentUser.APIToken = response.APIToken

}
func printUserData() {
	fmt.Println("Username: ", currentUser.Username)
	fmt.Println("Name: ", currentUser.Name)
	fmt.Println("Email: ", currentUser.Email)
	fmt.Println("Initials: ", currentUser.Initials)
	fmt.Println("Timezone: ", currentUser.Timezone.Kind, currentUser.Timezone.Offset, currentUser.Timezone.OlsonName)

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func setCredentials() {
	fmt.Fprint(Stdout, "Username: ")
	var username = cmdutil.ReadLine()
	cmdutil.Silence()
	fmt.Fprint(Stdout, "Password: ")

	var password = cmdutil.ReadLine()
	currentUser.Login(username, password)
	cmdutil.Unsilence()
}

func homeDir() string {
	usr, _ := u.Current()
	return usr.HomeDir
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
