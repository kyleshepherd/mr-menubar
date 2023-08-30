package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Data struct {
		MRs MRs `json:"currentUser"`
	} `json:"data"`
}

type MRs struct {
	Username string `json:"username"`
	Assigned struct {
		Count int `json:"count"`
	} `json:"assignedMergeRequests"`
	Review struct {
		Count int `json:"count"`
	} `json:"reviewRequestedMergeRequests"`
}

func GetMRs(token string) (*MRs, error) {
	queryMap := map[string]string{
		"query": `
			{
				currentUser {
					username
					assignedMergeRequests(state: opened) {
							count
					}
					reviewRequestedMergeRequests(state:opened) {
							count
					}
				}
			}
		`,
	}

	jsonQuery, err := json.Marshal(queryMap)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON instance %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://gitlab.com/api/graphql", bytes.NewBuffer(jsonQuery))
	if err != nil {
		return nil, fmt.Errorf("error creating request %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: time.Second * 5}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request %v", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading data %v", err)
	}

	resStruct := Response{}

	err = json.Unmarshal(data, &resStruct)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON %v", err)
	}

	return &resStruct.Data.MRs, nil
}
