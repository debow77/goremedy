package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.cerner.com/OHAIFedAutoSre/gorapid"
)

// RemedyClientInterface is an interface for a Remedy client
type RemedyClientInterface interface {
	GetRapidClient() *gorapid.RapidClient
}

// PageResponse represents a paginated API response
type PageResponse struct {
	Content    []json.RawMessage `json:"content"`
	TotalPages int               `json:"totalPages"`
}

// GetPaginated returns a list of items from a paginated API response
func GetPaginated(client RemedyClientInterface, basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error) {
	const pageSize = 1000
	params.Set("size", fmt.Sprintf("%d", pageSize))
	params.Set("page", "-1")

	var allItems []json.RawMessage
	totalPages := 1
	var statusCode int

	for page := 0; page < totalPages; page++ {
		params.Set("page", fmt.Sprintf("%d", page))

		items, status, err := getPage(client, basePath, urlPath, params)
		if err != nil {
			return nil, status, err
		}

		statusCode = status
		allItems = append(allItems, items.Content...)
		totalPages = items.TotalPages
	}

	return allItems, statusCode, nil
}

// getPage retrieves a single page of results from the API
func getPage(client RemedyClientInterface, basePath, urlPath string, params url.Values) (*PageResponse, int, error) {
	var retryCount int
	for {
		resp, err := client.GetRapidClient().Get(basePath+urlPath, params)
		if err != nil {
			if retryCount < 3 {
				retryCount++
				time.Sleep(time.Duration(retryCount*500) * time.Millisecond)
				continue
			}
			return nil, 0, fmt.Errorf("failed to make API request: %v", err)
		}
		defer resp.Body.Close()

		if resp.Status == 429 {
			retryCount++
			time.Sleep(time.Duration(retryCount*500) * time.Millisecond)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, resp.Status, fmt.Errorf("failed to read response body: %v", err)
		}

		var pageResp PageResponse
		if err := json.Unmarshal(body, &pageResp); err != nil {
			return nil, resp.Status, fmt.Errorf("failed to unmarshal API response: %v", err)
		}

		return &pageResp, resp.Status, nil
	}
}
