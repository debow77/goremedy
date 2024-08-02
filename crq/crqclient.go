package crq

import (
	"encoding/json"
	"fmt"
	"goremedy/internal/common"
	"io"
	"log/slog"
	"net/url"
)

const (
	changeQueryPath = "remedy-change-query-svc/v1"
)

type ClientGroup interface {
	Get(changeID string) (*CRQResponse, error)
	GetByUtn(changeUtn string) (*UtnResponse, error)
}

type clientGroup struct {
	client   common.RemedyClientInterface
	crQuery  *queryClient
	crModify *modifyClient
}

func NewClientGroup(client common.RemedyClientInterface) (ClientGroup, error) {
	return &clientGroup{
		client:   client,
		crQuery:  newQueryClient(client),
		crModify: newModifyClient(client),
	}, nil
}

type queryClient struct {
	client common.RemedyClientInterface
}

func newQueryClient(client common.RemedyClientInterface) *queryClient {
	return &queryClient{client: client}
}

func (qc *queryClient) get(changeID string) (*CRQResponse, error) {
	urlPath := fmt.Sprintf("%s/changes/%s/all", changeQueryPath, changeID)
	slog.Debug("Getting CRQ", "urlPath", urlPath)

	rapidClient := qc.client.GetRapidClient()
	resp, err := rapidClient.Get(urlPath, url.Values{})
	if err != nil {
		return nil, fmt.Errorf("failed to get CRQ: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var crq CRQResponse
	if err := json.Unmarshal(body, &crq); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CRQ: %w", err)
	}
	return &crq, nil
}

func (qc *queryClient) getByUtn(changeUtn string) (*UtnResponse, error) {
	urlPath := fmt.Sprintf("%s/changes", changeQueryPath)
	params := url.Values{
		"universalTicketNumber": {changeUtn},
	}

	slog.Debug("Getting CRQ", "urlPath", urlPath, "params", params)

	rapidClient := qc.client.GetRapidClient()
	resp, err := rapidClient.Get(urlPath, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get CRQ by UTN: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var response UtnResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &response, nil
}

func (cg *clientGroup) Get(changeID string) (*CRQResponse, error) {
	return cg.crQuery.get(changeID)
}

func (cg *clientGroup) GetByUtn(changeUtn string) (*UtnResponse, error) {
	return cg.crQuery.getByUtn(changeUtn)
}

type modifyClient struct {
	client common.RemedyClientInterface
}

func newModifyClient(client common.RemedyClientInterface) *modifyClient {
	return &modifyClient{client: client}
}
