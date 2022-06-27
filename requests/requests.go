package requests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	ghRepoOwner                  = os.Getenv("GITHUB_REPOSITORY_OWNER")
	ghRef                        = os.Getenv("GITHUB_REF")
	ghRefName                    = os.Getenv("GITHUB_REF_NAME")
	ghRepo                       = os.Getenv("GITHUB_REPOSITORY")
	ghToken                      = os.Getenv("GITHUB_TOKEN")
	ghEvent                      = os.Getenv("GITHUB_EVENT_NAME")
	ghActor                      = os.Getenv("GITHUB_ACTOR")
	ghWorkflow                   = os.Getenv("GITHUB_WORKFLOW")
	ghActionsIDTokenRequestURL   = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	ghActionsIDTokenRequestToken = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	ghAPIURL                     = os.Getenv("GITHUB_API_URL")
)

// Get URL of github repository
func GetRepoUrl() string {
	// https://api.github.com/repos/OWNER/REPO/labels
	return ghAPIURL + "/repos/" + ghRepo + "/labels"
}

// Get URL of current github pull request
func GetPRUrl() string {
	prNumber := strings.Split(ghRefName, "/")[0]
	return ghAPIURL + "/repos/" + ghRepo + "/issues/" + prNumber + "/labels"
}

// SendRequest function for sending requests to the github API
func SendRequest(requestMethod string, url string, requestBody []byte) (int, []byte) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	// Init request and add required headers
	request, err := http.NewRequest(requestMethod, url, bytes.NewBuffer(requestBody))
	request.Header.Add("Accept", "application/vnd.github.v3+json")
	request.Header.Add("Authorization", "token "+ghToken)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	// Send request
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
		
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	statusCode := resp.StatusCode
	return statusCode, body
}

func PrintResponse(status int, response []byte) {
	fmt.Println(">> ", status)
	fmt.Println("\t", string(response))
	fmt.Println()

}

func TestSendRequest(requestMethod string, url string, requestBody []byte) {
	request, err := http.NewRequest(requestMethod, url, bytes.NewBuffer(requestBody))
	request.Header.Add("Accept", "application/vnd.github.v3+json")
	request.Header.Add("Authorization", "token "+ghToken)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	for _, i := range request.Form {
		fmt.Println(i)

	}

}
