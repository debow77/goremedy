package crq_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"goremedy/crq"
	"io"
	"net/http"
	"net/url"
	"testing"
)

// MockRemedyClient is a mock implementation of the RemedyClientInterface
type MockRemedyClient struct {
	Response *http.Response
	Error    error
}
type RapidClientInterface interface {
	GenerateToken() (string, error)
}
type MockRapidClient struct{}

func (m *MockRapidClient) GenerateToken() (string, error) {
	// Return a mock token or an error
	return "mock-token", nil
}
func (m *MockRemedyClient) GetRapidClient() RapidClientInterface {
	return &MockRapidClient{}
}
func (m *MockRemedyClient) Get(basePath, urlPath string, params url.Values) (*http.Response, error) {
	return m.Response, m.Error
}
func TestClientGroup_Get(t *testing.T) {
	// Setup mock response data
	crqData := crq.CRQResponse{
		CommonFields: crq.CommonFields{
			ChangeID:              "CRQ123456",
			UniversalTicketNumber: "UTN123456",
			Summary:               "Test Change Request",
		},
	}
	body, _ := json.Marshal(crqData)
	mockResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(body)),
	}
	// Initialize the mock client and clientGroup
	mockClient := &MockRemedyClient{Response: mockResponse, Error: nil}
	clientGroup, err := crq.NewClientGroup(mockClient)
	if err != nil {
		t.Fatalf("expected no error creating client group, got %v", err)
	}
	// Call the Get method and check the results
	result, err := clientGroup.Get("CRQ123456")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.ChangeID != crqData.ChangeID {
		t.Errorf("expected ChangeID %s, got %s", crqData.ChangeID, result.ChangeID)
	}
	if result.UniversalTicketNumber != crqData.UniversalTicketNumber {
		t.Errorf("expected UniversalTicketNumber %s, got %s", crqData.UniversalTicketNumber, result.UniversalTicketNumber)
	}
}
func TestClientGroup_GetByUtn(t *testing.T) {
	// Setup mock response data
	utnData := crq.UtnResponse{
		Content: []struct {
			crq.CommonFields
		}{
			{
				CommonFields: crq.CommonFields{
					ChangeID:              "CRQ123456",
					UniversalTicketNumber: "UTN123456",
					Summary:               "Test Change Request",
				},
			},
		},
		TotalElements: 1,
		TotalPages:    1,
	}
	body, _ := json.Marshal(utnData)
	mockResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(body)),
	}
	// Initialize the mock client and clientGroup
	mockClient := &MockRemedyClient{Response: mockResponse, Error: nil}
	clientGroup, err := crq.NewClientGroup(mockClient)
	if err != nil {
		t.Fatalf("expected no error creating client group, got %v", err)
	}
	// Call the GetByUtn method and check the results
	result, err := clientGroup.GetByUtn("UTN123456")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(result.Content) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result.Content))
	}
	if result.Content[0].ChangeID != utnData.Content[0].ChangeID {
		t.Errorf("expected ChangeID %s, got %s", utnData.Content[0].ChangeID, result.Content[0].ChangeID)
	}
	if result.Content[0].UniversalTicketNumber != utnData.Content[0].UniversalTicketNumber {
		t.Errorf("expected UniversalTicketNumber %s, got %s", utnData.Content[0].UniversalTicketNumber, result.Content[0].UniversalTicketNumber)
	}
}
func TestClientGroup_Get_Error(t *testing.T) {
	// Setup mock error response
	mockClient := &MockRemedyClient{Response: nil, Error: errors.New("some error")}
	clientGroup, err := crq.NewClientGroup(mockClient)
	if err != nil {
		t.Fatalf("expected no error creating client group, got %v", err)
	}
	// Call the Get method and expect an error
	_, err = clientGroup.Get("CRQ123456")
	if err == nil {
		t.Fatal("expected error, got none")
	}
}
func TestClientGroup_GetByUtn_Error(t *testing.T) {
	// Setup mock error response
	mockClient := &MockRemedyClient{Response: nil, Error: errors.New("some error")}
	clientGroup, err := crq.NewClientGroup(mockClient)
	if err != nil {
		t.Fatalf("expected no error creating client group, got %v", err)
	}
	// Call the GetByUtn method and expect an error
	_, err = clientGroup.GetByUtn("UTN123456")
	if err == nil {
		t.Fatal("expected error, got none")
	}
}
