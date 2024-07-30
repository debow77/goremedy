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

// GetPaginated returns a list of items from a paginated API response
// func GetPaginated(client RemedyClientInterface, basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error) {
// 	const pageSize = 1000
// 	params.Set("size", fmt.Sprintf("%d", pageSize))
// 	params.Set("page", "-1")

// 	var allItems []json.RawMessage
// 	totalPages := 1
// 	var statusCode int

// 	for page := 0; page < totalPages; page++ {
// 		params.Set("page", fmt.Sprintf("%d", page))

// 		resp, err := client.GetRapidClient().Get(basePath+urlPath, params)
// 		fmt.Printf("resp to unmarshal API response: %v", resp.Status)
// 		if err != nil {
// 			return nil, 0, fmt.Errorf("failed to make API request: %v", err)
// 		}

// 		defer resp.Body.Close() // close the response body!

// 		statusCode = resp.Status

// 		var responseBody []byte
// 		responseBody, err = io.ReadAll(resp.Body)
// 		if err != nil {
// 			return nil, 0, fmt.Errorf("failed to read response body: %v", err)
// 		}

// 		var pageResp struct {
// 			Content    []json.RawMessage `json:"content"`
// 			TotalPages int               `json:"totalPages"`
// 		}

// 		if err := json.Unmarshal(responseBody, &pageResp); err != nil {
// 			return nil, 0, fmt.Errorf("failed to unmarshal API response: %v", err)
// 		}

// 		allItems = append(allItems, pageResp.Content...)
// 		totalPages = pageResp.TotalPages
// 	}

// 	return allItems, statusCode, nil
// }

func GetPaginated(client RemedyClientInterface, basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error) {
	const pageSize = 1000
	params.Set("size", fmt.Sprintf("%d", pageSize))
	params.Set("page", "-1")

	var allItems []json.RawMessage
	totalPages := 1
	var statusCode int

	for page := 0; page < totalPages; page++ {
		params.Set("page", fmt.Sprintf("%d", page))

		var retryCount int
		for {
			// Retry logic to account for random 429 errors
			resp, err := client.GetRapidClient().Get(basePath+urlPath, params)
			if err != nil {
				if retryCount < 3 { // adjust the retry count as needed
					retryCount++
					time.Sleep(time.Duration(retryCount*500) * time.Millisecond) // exponential backoff
					continue
				}
				return nil, 0, fmt.Errorf("failed to make API request: %v", err)
			}

			defer resp.Body.Close() // close the response body!

			statusCode = resp.Status

			if resp.Status == 429 {
				retryCount++
				time.Sleep(time.Duration(retryCount*500) * time.Millisecond) // exponential backoff
				continue
			}

			var responseBody []byte
			responseBody, err = io.ReadAll(resp.Body)
			if err != nil {
				return nil, 0, fmt.Errorf("failed to read response body: %v", err)
			}

			var pageResp struct {
				Content    []json.RawMessage `json:"content"`
				TotalPages int               `json:"totalPages"`
			}

			if err := json.Unmarshal(responseBody, &pageResp); err != nil {
				return nil, 0, fmt.Errorf("failed to unmarshal API response: %v", err)
			}

			allItems = append(allItems, pageResp.Content...)
			totalPages = pageResp.TotalPages
			break
		}
	}

	return allItems, statusCode, nil
}
