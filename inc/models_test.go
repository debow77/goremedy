package inc

import (
	"encoding/json"
	"goremedy/common"
	"testing"
)

func TestINCResponseUnmarshal(t *testing.T) {
	jsonData := `{
		"IncidentId": "INC000123",
		"universalTicketNumber": "UTN123456",
		"summary": "Test Inc",
		"urgencyString": "2-High",
		"impactString": "3-Moderate/Limited",
		"changeTimingString": "Normal",
		"statusString": "Pending",
		"workLogs": [
			{
				"workLogId": "WL001",
				"IncidentId": "INC000123",
				"submitter": "Jon Snow",
				"workLogSubmitDate": "2023-08-06T10:00:00Z",
				"notes": "Test work log"
			}
		],
		"relationships": [
			{
				"relationshipId": "REL001",
				"IncidentId": "INC000123",
				"relatedTo": "INC000124",
				"relationshipType": "Related"
			}
		],
		"auditlogs": [
			{
				"auditlogId": "AL001",
				"submitter": "Jane Snow",
				"createDate": "2023-08-06T11:00:00Z",
				"log": "Test audit log"
			}
		]
	}`

	var inc IncResponse
	err := json.Unmarshal([]byte(jsonData), &inc)
	if err != nil {
		t.Fatalf("Failed to unmarshal IncResponse: %v", err)
	}

	// Test common fields
	if inc.IncidentId != "INC000123" {
		t.Errorf("Expected IncidentId to be 'INC000123', got '%s'", inc.IncidentId)
	}
	if inc.UniversalTicketNumber != "UTN123456" {
		t.Errorf("Expected UniversalTicketNumber to be 'UTN123456', got '%s'", inc.UniversalTicketNumber)
	}
	if inc.Summary != "Test Inc" {
		t.Errorf("Expected Summary to be 'Test Inc', got '%s'", inc.Summary)
	}
	if inc.Urgency != common.UrgencyHigh {
		t.Errorf("Expected Urgency to be 'UrgencyHigh', got '%s'", inc.Urgency)
	}
	if inc.Impact != common.ImpactModerateLimited {
		t.Errorf("Expected Impact to be 'ImpactModerateLimited', got '%s'", inc.Impact)
	}
	if inc.Status != common.StatusPending {
		t.Errorf("Expected Status to be 'StatusPending', got '%s'", inc.Status)
	}

	// Test WorkLogs
	if len(inc.WorkLogs) != 1 {
		t.Errorf("Expected 1 WorkLog, got %d", len(inc.WorkLogs))
	} else {
		wl := inc.WorkLogs[0]
		if wl.WorkLogID != "WL001" {
			t.Errorf("Expected WorkLogID to be 'WL001', got '%s'", wl.WorkLogID)
		}
	}

	// Test Relationships
	if len(inc.Relationships) != 1 {
		t.Errorf("Expected 1 Relationship, got %d", len(inc.Relationships))
	} else {
		rel := inc.Relationships[0]
		if rel.RelationshipID != "REL001" {
			t.Errorf("Expected RelationshipID to be 'REL001', got '%s'", rel.RelationshipID)
		}
	}

	// Test Auditlogs
	if len(inc.Auditlogs) != 1 {
		t.Errorf("Expected 1 Auditlog, got %d", len(inc.Auditlogs))
	} else {
		al := inc.Auditlogs[0]
		if al.AuditlogID != "AL001" {
			t.Errorf("Expected AuditlogID to be 'AL001', got '%s'", al.AuditlogID)
		}
	}
}

func TestINCUtnResponseUnmarshal(t *testing.T) {
	jsonData := `{
		"content": [
			{
				"IncidentId": "INC000123",
				"universalTicketNumber": "UTN123456",
				"summary": "Test INC",
				"urgencyString": "2-High",
				"impactString": "3-Moderate/Limited",
				"changeTimingString": "Normal",
				"statusString": "Pending"
			}
		],
		"totalElements": 1,
		"totalPages": 1
	}`

	var utnResp IncUtnResponse
	err := json.Unmarshal([]byte(jsonData), &utnResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal UtnResponse: %v", err)
	}

	if len(utnResp.Content) != 1 {
		t.Errorf("Expected 1 Content item, got %d", len(utnResp.Content))
	} else {
		inc := utnResp.Content[0]
		if inc.IncidentId != "INC000123" {
			t.Errorf("Expected IncidentId to be 'INC000123', got '%s'", inc.IncidentId)
		}
		if inc.UniversalTicketNumber != "UTN123456" {
			t.Errorf("Expected UniversalTicketNumber to be 'UTN123456', got '%s'", inc.UniversalTicketNumber)
		}
		if inc.Summary != "Test INC" {
			t.Errorf("Expected Summary to be 'Test INC', got '%s'", inc.Summary)
		}
		if inc.Urgency != common.UrgencyHigh {
			t.Errorf("Expected Urgency to be 'UrgencyHigh', got '%s'", inc.Urgency)
		}
		if inc.Impact != common.ImpactModerateLimited {
			t.Errorf("Expected Impact to be 'ImpactModerateLimited', got '%s'", inc.Impact)
		}
		if inc.Status != common.StatusPending {
			t.Errorf("Expected Status to be 'StatusPending', got '%s'", inc.Status)
		}
	}

	if utnResp.TotalElements != 1 {
		t.Errorf("Expected TotalElements to be 1, got %d", utnResp.TotalElements)
	}
	if utnResp.TotalPages != 1 {
		t.Errorf("Expected TotalPages to be 1, got %d", utnResp.TotalPages)
	}
}
