package testutils

import (
	"encoding/json"
	"goremedy/interfaces"
	"net/url"

	"github.cerner.com/OHAIFedAutoSre/gorapid"
)

// MockRapidClient implements the RapidClient interface for testing
type MockRapidClient struct {
	GetFunc func(urlPath string, params url.Values) (*gorapid.Response, error)
}

func (m *MockRapidClient) Get(urlPath string, params url.Values) (*gorapid.Response, error) {
	return m.GetFunc(urlPath, params)
}

func (m *MockRapidClient) Post(urlPath string, body gorapid.JSONBody) (*gorapid.Response, error) {
	return nil, nil
}

func (m *MockRapidClient) Put(urlPath string, body gorapid.JSONBody) (*gorapid.Response, error) {
	return nil, nil
}

func (m *MockRapidClient) Patch(urlPath string, body gorapid.JSONBody) (*gorapid.Response, error) {
	return nil, nil
}

func (m *MockRapidClient) Delete(urlPath string) (*gorapid.Response, error) {
	return nil, nil
}

func (m *MockRapidClient) BaseURL() string {
	return "http://rapid.com"
}

func (m *MockRapidClient) GetPaginated(basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error) {
	return nil, 0, nil
}

func (m *MockRapidClient) GetPage(basePath, urlPath string, params url.Values) (*interfaces.PageResponse, int, error) {
	return nil, 0, nil
}

// MockRapidClientInterface implements the RapidClientInterface for testing
type MockRapidClientInterface struct {
	MockClient *MockRapidClient
}

func (m *MockRapidClientInterface) GetRapidClient() interfaces.RapidClient {
	return m.MockClient
}

func (m *MockRapidClientInterface) BaseURL() string {
	return m.MockClient.BaseURL()
}

func (m *MockRapidClientInterface) GetPaginated(basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error) {
	return m.MockClient.GetPaginated(basePath, urlPath, params)
}
