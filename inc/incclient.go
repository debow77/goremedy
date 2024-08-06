package inc

import (
	"encoding/json"
	"fmt"
	"goremedy/interfaces"
	"io"
	"log/slog"
	"net/url"
)

const (
	incQueryPath = "/remedy-incident-query-svc/v1"
)

type ClientGroup interface {
	Get(changeID string) (*IncResponse, error)
	GetByUtn(incUtn string) (*IncUtnResponse, error)
}

type clientGroup struct {
	client   interfaces.RapidClientInterface
	incQuery *queryClient
}

func NewClientGroup(client interfaces.RapidClientInterface) (ClientGroup, error) {
	return &clientGroup{
		client:   client,
		incQuery: newQueryClient(client),
	}, nil
}

type queryClient struct {
	client interfaces.RapidClientInterface
}

func newQueryClient(client interfaces.RapidClientInterface) *queryClient {
	return &queryClient{client: client}
}

func (qc *queryClient) get(incNumber string) (*IncResponse, error) {
	urlPath := fmt.Sprintf("%s/incidents/%s/all", incQueryPath, incNumber)
	slog.Debug("Getting INC", "urlPath", urlPath)

	resp, err := qc.client.GetRapidClient().Get(urlPath, url.Values{})
	if err != nil {
		return nil, fmt.Errorf("failed to get INC: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read get INC response body: %w", err)
	}
	// fmt.Println(string(body))
	var inc IncResponse
	if err := json.Unmarshal(body, &inc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal INC: %w", err)
	}

	return &inc, nil
}

func (qc *queryClient) getByUtn(incUtn string) (*IncUtnResponse, error) {
	urlPath := fmt.Sprintf("%s/incidents", incQueryPath)
	params := url.Values{
		"universalTicketNumber": {incUtn},
	}

	slog.Debug("Getting INC", "urlPath", urlPath, "params", params)

	resp, err := qc.client.GetRapidClient().Get(urlPath, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get INC by UTN: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read INC UTN response body: %w", err)
	}
	// fmt.Println(string(body))
	var response IncUtnResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &response, nil
}

func (inc *clientGroup) Get(incNumber string) (*IncResponse, error) {
	return inc.incQuery.get(incNumber)
}

func (inc *clientGroup) GetByUtn(incUtn string) (*IncUtnResponse, error) {
	return inc.incQuery.getByUtn(incUtn)
}
