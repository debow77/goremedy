package common

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.cerner.com/OHAIFedAutoSre/gorapid"
)

type RemedyClientInterface interface {
	GetRapidClient() *gorapid.RapidClient
}

func GetPaginated(client RemedyClientInterface, basePath, urlPath string, params url.Values) ([]json.RawMessage, error) {
	const pageSize = 1000
	params.Set("size", fmt.Sprintf("%d", pageSize))
	params.Set("page", "0")

	var allItems []json.RawMessage
	totalPages := 1

	for page := 0; page < totalPages; page++ {
		params.Set("page", fmt.Sprintf("%d", page))

		resp, err := client.GetRapidClient().Get(basePath+urlPath, params)
		if err != nil {
			return nil, fmt.Errorf("failed to make API request: %v", err)
		}

		var pageResp struct {
			Content    []json.RawMessage `json:"content"`
			TotalPages int               `json:"totalPages"`
		}

		// if err := json.Unmarshal(resp, &pageResp); err != nil {
		if err := json.Unmarshal(resp.Body, &pageResp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
		}

		allItems = append(allItems, pageResp.Content...)
		totalPages = pageResp.TotalPages
	}

	return allItems, nil
}
