package common

import (
	"encoding/json"
	"testing"
)

func TestUrgencyUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Urgency
	}{
		{"Critical", `"1-Critical"`, UrgencyCritical},
		{"High", `"2-High"`, UrgencyHigh},
		{"Medium", `"3-Medium"`, UrgencyMedium},
		{"Low", `"4-Low"`, UrgencyLow},
		{"Unknown", `"Unknown"`, UrgencyUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u Urgency
			err := json.Unmarshal([]byte(tt.input), &u)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if u != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, u)
			}
		})
	}
}

func TestImpactUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Impact
	}{
		{"Extensive/Widespread", `"1-Extensive/Widespread"`, ImpactExtensiveWidespread},
		{"Significant/Large", `"2-Significant/Large"`, ImpactSignificantLarge},
		{"Moderate/Limited", `"3-Moderate/Limited"`, ImpactModerateLimited},
		{"Minor/Localized", `"4-Minor/Localized"`, ImpactMinorLocalized},
		{"Unknown", `"Unknown"`, ImpactUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i Impact
			err := json.Unmarshal([]byte(tt.input), &i)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if i != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, i)
			}
		})
	}
}

func TestChangeTimingUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ChangeTiming
	}{
		{"Emergency", `"Emergency"`, ChangeTimingEmergency},
		{"Normal", `"Normal"`, ChangeTimingNormal},
		{"Standard", `"Standard"`, ChangeTimingStandard},
		{"Unknown", `"Unknown"`, ChangeTimingUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ct ChangeTiming
			err := json.Unmarshal([]byte(tt.input), &ct)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if ct != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, ct)
			}
		})
	}
}

func TestStatusUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Status
	}{
		{"Draft", `"Draft"`, StatusDraft},
		{"Cancelled", `"Cancelled"`, StatusCancelled},
		{"Closed", `"Closed"`, StatusClosed},
		{"Completed", `"Completed"`, StatusCompleted},
		{"Implementation In Progress", `"Implementation In Progress"`, StatusImplementationInProgress},
		{"Pending", `"Pending"`, StatusPending},
		{"Planning In Progress", `"Planning In Progress"`, StatusPlanningInProgress},
		{"Rejected", `"Rejected"`, StatusRejected},
		{"Request For Authorization", `"Request For Authorization"`, StatusRequestForAuthorization},
		{"Scheduled", `"Scheduled"`, StatusScheduled},
		{"Scheduled For Approval", `"Scheduled For Approval"`, StatusScheduledForApproval},
		{"Scheduled For Review", `"Scheduled For Review"`, StatusScheduledForReview},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Status
			err := json.Unmarshal([]byte(tt.input), &s)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if s != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, s)
			}
		})
	}
}
