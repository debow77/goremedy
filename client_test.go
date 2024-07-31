package goremedy

import (
	"testing"
)

func TestNewRemedyClient(t *testing.T) {
	tests := []struct {
		name        string
		config      RemedyClientConfig
		expectError bool
	}{
		{
			name:        "Valid config",
			config:      RemedyClientConfig{LogLevel: "INFO"},
			expectError: false,
		},
		{
			name:        "Invalid log level",
			config:      RemedyClientConfig{LogLevel: "INVALID"},
			expectError: false, // It should default to INFO and not error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewRemedyClient(tt.config)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
				if client != nil {
					t.Errorf("Expected client to be nil, but got %v", client)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if client == nil {
					t.Errorf("Expected client to be non-nil, but got nil")
				}
				if client.GetRapidClient() == nil {
					t.Errorf("Expected RapidClient to be non-nil")
				}
				if client.GetCompanyClientGroup() == nil {
					t.Errorf("Expected CompanyClientGroup to be non-nil")
				}
				if client.GetCIClientGroup() == nil {
					t.Errorf("Expected CIClientGroup to be non-nil")
				}
			}
		})
	}
}
