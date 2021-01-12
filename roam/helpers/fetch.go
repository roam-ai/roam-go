package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	apiBaseURL = "https://sdk.geospark.co"

	// INVALIDAPIKEY is error reponse for invalid api key
	invalidapikey = "Entered API Key is Invalid"
)

// APIResponse is struct to unmarshal response json from api
type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"msg"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

// ProjectData is struct to unmarshal details endpoint data
type ProjectData struct {
	AccountID string `json:"account_id"`
	ProjectID string `json:"project_id"`
}

// GroupData is struct to unmarshal group data
type GroupData struct {
	UserIDs []string `json:"user_ids"`
}

// fetch makes get request to api, decodes body and returns data
func fetch(apikey, url string) (data interface{}, err error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	request.Header.Add("api-key", apikey)
	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusUnauthorized {
		return nil, errors.New(invalidapikey)
	}
	if response.StatusCode == http.StatusOK {
		var apiResponse APIResponse
		err := json.NewDecoder(response.Body).Decode(&apiResponse)
		if err != nil {
			return nil, err
		}
		return apiResponse.Data, nil
	}
	return nil, errors.New("Unknown Error")
}

// GetProjectDetails gets the account_id and project_id of given apikey
func GetProjectDetails(apikey string) (accountID, projectID string, err error) {
	getDetailsURL := apiBaseURL + "/api/details"
	data, err := fetch(apikey, getDetailsURL)
	if err != nil {
		return "", "", err
	}
	var projectData ProjectData
	marshal, _ := json.Marshal(data)
	err = json.Unmarshal(marshal, &projectData)
	if err != nil {
		return "", "", err
	}
	return projectData.AccountID, projectData.ProjectID, nil
}

// GetGroupData gets all users from given group_id
func GetGroupData(apikey, groupID string) (userIDs []string, err error) {
	getGroupURL := apiBaseURL + "/api/group/" + groupID
	data, err := fetch(apikey, getGroupURL)
	if err != nil {
		return userIDs, err
	}
	var groupData GroupData
	marshal, _ := json.Marshal(data)
	err = json.Unmarshal(marshal, &groupData)
	if err != nil {
		return userIDs, err
	}
	return groupData.UserIDs, nil
}
