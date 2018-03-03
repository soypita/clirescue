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

const (
  get string = "GET"
  post string = "POST"
  put string = "PUT"
)

var (
  currentUser  *user.User = user.New()
  tokenLocation string     = homeDir() + "/.tracker"
  Stdout       *os.File   = os.Stdout
)

func autorization() *user.User {
  dat, err := ioutil.ReadFile(tokenLocation)
	check(err)
	if dat != nil {
		currentUser.APIToken = string(dat)
	} else {
		setCredentials()
	}
  return currentUser
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

func makeRequest(method string, url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
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

func parse(body []byte, resp interface{}) {
  err := json.Unmarshal(body, &resp)
  if err != nil {
    fmt.Println("error:", err)
  }
  // ioutil.WriteFile(tokenLocation, []byte(meResp.APIToken), 0644)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func homeDir() string {
	usr, _ := u.Current()
	return usr.HomeDir
}
