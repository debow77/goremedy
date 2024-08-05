package interfaces

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.cerner.com/OHAIFedAutoSre/gorapid"
)

type RapidClientInterface interface {
	GetRapidClient() RapidClient
	BaseURL() string
	GetPaginated(basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error)
}

type RapidClient interface {
	Get(urlPath string, params url.Values) (*gorapid.Response, error)
	Post(urlPath string, body gorapid.JSONBody) (*gorapid.Response, error)
	Put(urlPath string, body gorapid.JSONBody) (*gorapid.Response, error)
	Patch(urlPath string, body gorapid.JSONBody) (*gorapid.Response, error)
	Delete(urlPath string) (*gorapid.Response, error)
	BaseURL() string
	GetPaginated(basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error)
	GetPage(basePath, urlPath string, params url.Values) (*PageResponse, int, error)
}

func NewRapidClient() (RapidClientInterface, error) {
	rapidClient, err := gorapid.NewRapidClient()
	if err != nil {
		return nil, err
	}
	return &rapidClientImpl{client: rapidClient}, nil
}

type rapidClientImpl struct {
	client *gorapid.RapidClient
}

func (r *rapidClientImpl) GetRapidClient() RapidClient {
	return &rapidClientWrapper{r.client}
}

func (r *rapidClientImpl) BaseURL() string {
	return r.client.BaseURL
}

func (r *rapidClientImpl) GetPaginated(basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error) {
	return r.GetRapidClient().GetPaginated(basePath, urlPath, params)
}

type rapidClientWrapper struct {
	*gorapid.RapidClient
}

func (r *rapidClientWrapper) Get(urlPath string, params url.Values) (*gorapid.Response, error) {
	return r.RapidClient.Get(urlPath, params)
}

func (r *rapidClientWrapper) Post(urlPath string, body gorapid.JSONBody) (*gorapid.Response, error) {
	return r.RapidClient.Post(urlPath, body)
}

func (r *rapidClientWrapper) Put(urlPath string, body gorapid.JSONBody) (*gorapid.Response, error) {
	return r.RapidClient.Put(urlPath, body)
}

func (r *rapidClientWrapper) Patch(urlPath string, body gorapid.JSONBody) (*gorapid.Response, error) {
	return r.RapidClient.Request("PATCH", urlPath, body, nil)
}

func (r *rapidClientWrapper) Delete(urlPath string) (*gorapid.Response, error) {
	return r.RapidClient.Delete(urlPath)
}

func (r *rapidClientWrapper) BaseURL() string {
	return r.RapidClient.BaseURL
}

// PageResponse represents a paginated API response
type PageResponse struct {
	Content    []json.RawMessage `json:"content"`
	TotalPages int               `json:"totalPages"`
}

func (r *rapidClientWrapper) GetPaginated(basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error) {
	const pageSize = 1000
	params.Set("size", fmt.Sprintf("%d", pageSize))
	params.Set("page", "-1")
	var allItems []json.RawMessage
	totalPages := 1
	var statusCode int
	for page := 0; page < totalPages; page++ {
		params.Set("page", fmt.Sprintf("%d", page))
		items, status, err := r.GetPage(basePath, urlPath, params)
		if err != nil {
			return nil, status, err
		}
		statusCode = status
		allItems = append(allItems, items.Content...)
		totalPages = items.TotalPages
	}
	return allItems, statusCode, nil
}

func (r *rapidClientWrapper) GetPage(basePath, urlPath string, params url.Values) (*PageResponse, int, error) {
	var retryCount int
	for {
		resp, err := r.Get(basePath+urlPath, params)
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
