package crq

import (
	"encoding/json"
	"goremedy/interfaces"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

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

func TestGet(t *testing.T) {
	mockResponse := `{"changeId": "12345", "statusString": "Draft"}`
	mockClient := &MockRapidClient{
		GetFunc: func(urlPath string, params url.Values) (*gorapid.Response, error) {
			return &gorapid.Response{
				Body:         io.NopCloser(strings.NewReader(mockResponse)),
				Status:       http.StatusOK,
				Headers:      http.Header{"Content-Type": []string{"application/json"}},
				Error:        nil,
				ResponseTime: 100 * time.Millisecond,
				RequestURL:   urlPath,
			}, nil
		},
	}

	mockInterface := &MockRapidClientInterface{MockClient: mockClient}

	cg, err := NewClientGroup(mockInterface)
	if err != nil {
		t.Fatalf("Failed to create ClientGroup: %v", err)
	}

	crq, err := cg.Get("12345")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if crq.ChangeID != "12345" {
		t.Errorf("Expected ChangeID 12345, got %s", crq.ChangeID)
	}

	if crq.Status != "Draft" {
		t.Errorf("Expected Status Draft, got %s", crq.Status)
	}
}

func TestGetByUtn(t *testing.T) {
	mockResponse := `{"content": [{"changeId": "12345", "universalTicketNumber": "UTN12345"}]}`
	mockClient := &MockRapidClient{
		GetFunc: func(urlPath string, params url.Values) (*gorapid.Response, error) {
			return &gorapid.Response{
				Body:         io.NopCloser(strings.NewReader(mockResponse)),
				Status:       http.StatusOK,
				Headers:      http.Header{"Content-Type": []string{"application/json"}},
				Error:        nil,
				ResponseTime: 100 * time.Millisecond,
				RequestURL:   urlPath,
			}, nil
		},
	}

	mockInterface := &MockRapidClientInterface{MockClient: mockClient}

	cg, err := NewClientGroup(mockInterface)
	if err != nil {
		t.Fatalf("Failed to create ClientGroup: %v", err)
	}

	utnResp, err := cg.GetByUtn("UTN12345")
	if err != nil {
		t.Fatalf("GetByUtn failed: %v", err)
	}

	if len(utnResp.Content) != 1 {
		t.Fatalf("Expected 1 change, got %d", len(utnResp.Content))
	}

	if utnResp.Content[0].ChangeID != "12345" {
		t.Errorf("Expected ChangeID 12345, got %s", utnResp.Content[0].ChangeID)
	}

	if utnResp.Content[0].UniversalTicketNumber != "UTN12345" {
		t.Errorf("Expected UTN UTN12345, got %s", utnResp.Content[0].UniversalTicketNumber)
	}
}

func TestNewClientGroup(t *testing.T) {
	mockInterface := &MockRapidClientInterface{MockClient: &MockRapidClient{}}

	cg, err := NewClientGroup(mockInterface)
	if err != nil {
		t.Fatalf("NewClientGroup failed: %v", err)
	}

	if cg == nil {
		t.Fatal("NewClientGroup returned nil")
	}

	_, ok := cg.(*clientGroup)
	if !ok {
		t.Error("NewClientGroup did not return a *clientGroup")
	}
}
