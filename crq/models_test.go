package crq

import (
	"encoding/json"
	"testing"
)

func TestUrgencyUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected Urgency
	}{
		{`"1-Critical"`, UrgencyCritical},
		{`"2-High"`, UrgencyHigh},
		{`"3-Medium"`, UrgencyMedium},
		{`"4-Low"`, UrgencyLow},
		{`"Unknown"`, UrgencyUnknown},
	}

	for _, test := range tests {
		var u Urgency
		err := json.Unmarshal([]byte(test.input), &u)
		if err != nil {
			t.Errorf("Failed to unmarshal %s: %v", test.input, err)
		}
		if u != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, u)
		}
	}
}

func TestImpactUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected Impact
	}{
		{`"1-Extensive/Widespread"`, ImpactExtensiveWidespread},
		{`"2-Significant/Large"`, ImpactSignificantLarge},
		{`"3-Moderate/Limited"`, ImpactModerateLimited},
		{`"4-Minor/Localized"`, ImpactMinorLocalized},
		{`"Unknown"`, ImpactUnknown},
	}

	for _, test := range tests {
		var i Impact
		err := json.Unmarshal([]byte(test.input), &i)
		if err != nil {
			t.Errorf("Failed to unmarshal %s: %v", test.input, err)
		}
		if i != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, i)
		}
	}
}

func TestChangeTimingUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected ChangeTiming
	}{
		{`"Emergency"`, ChangeTimingEmergency},
		{`"Normal"`, ChangeTimingNormal},
		{`"Standard"`, ChangeTimingStandard},
		{`"Unknown"`, ChangeTimingUnknown},
	}

	for _, test := range tests {
		var ct ChangeTiming
		err := json.Unmarshal([]byte(test.input), &ct)
		if err != nil {
			t.Errorf("Failed to unmarshal %s: %v", test.input, err)
		}
		if ct != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, ct)
		}
	}
}

func TestStatusUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected Status
	}{
		{`"Draft"`, StatusDraft},
		{`"Cancelled"`, StatusCancelled},
		{`"Closed"`, StatusClosed},
		{`"Completed"`, StatusCompleted},
		{`"Implementation In Progress"`, StatusImplementationInProgress},
		{`"Pending"`, StatusPending},
		{`"Planning In Progress"`, StatusPlanningInProgress},
		{`"Rejected"`, StatusRejected},
		{`"Request For Authorization"`, StatusRequestForAuthorization},
		{`"Scheduled"`, StatusScheduled},
		{`"Scheduled For Approval"`, StatusScheduledForApproval},
		{`"Scheduled For Review"`, StatusScheduledForReview},
	}

	for _, test := range tests {
		var s Status
		err := json.Unmarshal([]byte(test.input), &s)
		if err != nil {
			t.Errorf("Failed to unmarshal %s: %v", test.input, err)
		}
		if s != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, s)
		}
	}
}
