package inc

import (
	"goremedy/testutils"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.cerner.com/OHAIFedAutoSre/gorapid"
)

func TestGet(t *testing.T) {
	mockResponse := `{"incidentId": "12345", "statusString": "Draft"}`
	mockClient := &testutils.MockRapidClient{
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

	mockInterface := &testutils.MockRapidClientInterface{MockClient: mockClient}

	cg, err := NewClientGroup(mockInterface)
	if err != nil {
		t.Fatalf("Failed to create ClientGroup: %v", err)
	}

	inc, err := cg.Get("12345")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if inc.IncidentId != "12345" {
		t.Errorf("Expected incidentId 12345, got %s", inc.IncidentId)
	}

	if inc.Status != "Draft" {
		t.Errorf("Expected Status Draft, got %s", inc.Status)
	}
}

func TestGetByUtn(t *testing.T) {
	mockResponse := `{"content": [{"incidentId": "12345", "universalTicketNumber": "UTN12345"}]}`
	mockClient := &testutils.MockRapidClient{
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

	mockInterface := &testutils.MockRapidClientInterface{MockClient: mockClient}

	cg, err := NewClientGroup(mockInterface)
	if err != nil {
		t.Fatalf("Failed to create ClientGroup: %v", err)
	}

	utnResp, err := cg.GetByUtn("UTN12345")
	if err != nil {
		t.Fatalf("GetByUtn failed: %v", err)
	}

	if len(utnResp.Content) != 1 {
		t.Fatalf("Expected 1 INC, got %d", len(utnResp.Content))
	}

	if utnResp.Content[0].IncidentId != "12345" {
		t.Errorf("Expected incidentId 12345, got %s", utnResp.Content[0].IncidentId)
	}

	if utnResp.Content[0].UniversalTicketNumber != "UTN12345" {
		t.Errorf("Expected UTN UTN12345, got %s", utnResp.Content[0].UniversalTicketNumber)
	}
}

func TestNewClientGroup(t *testing.T) {
	mockInterface := &testutils.MockRapidClientInterface{MockClient: &testutils.MockRapidClient{}}

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
