package crq

import (
	"encoding/json"
	"goremedy/common"
	"testing"
)

func TestCRQResponseUnmarshal(t *testing.T) {
	jsonData := `{
		"changeId": "CRQ000123",
		"universalTicketNumber": "UTN123456",
		"summary": "Test CRQ",
		"urgencyString": "2-High",
		"impactString": "3-Moderate/Limited",
		"changeTimingString": "Normal",
		"statusString": "Pending",
		"workLogs": [
			{
				"workLogId": "WL001",
				"changeId": "CRQ000123",
				"submitter": "John Doe",
				"workLogSubmitDate": "2023-08-06T10:00:00Z",
				"notes": "Test work log"
			}
		],
		"relationships": [
			{
				"relationshipId": "REL001",
				"changeId": "CRQ000123",
				"relatedTo": "CRQ000124",
				"relationshipType": "Related"
			}
		],
		"auditlogs": [
			{
				"auditlogId": "AL001",
				"submitter": "Jane Doe",
				"createDate": "2023-08-06T11:00:00Z",
				"log": "Test audit log"
			}
		]
	}`

	var crq CRQResponse
	err := json.Unmarshal([]byte(jsonData), &crq)
	if err != nil {
		t.Fatalf("Failed to unmarshal CRQResponse: %v", err)
	}

	// Test common fields
	if crq.ChangeID != "CRQ000123" {
		t.Errorf("Expected ChangeID to be 'CRQ000123', got '%s'", crq.ChangeID)
	}
	if crq.UniversalTicketNumber != "UTN123456" {
		t.Errorf("Expected UniversalTicketNumber to be 'UTN123456', got '%s'", crq.UniversalTicketNumber)
	}
	if crq.Summary != "Test CRQ" {
		t.Errorf("Expected Summary to be 'Test CRQ', got '%s'", crq.Summary)
	}
	if crq.Urgency != common.UrgencyHigh {
		t.Errorf("Expected Urgency to be 'UrgencyHigh', got '%s'", crq.Urgency)
	}
	if crq.Impact != common.ImpactModerateLimited {
		t.Errorf("Expected Impact to be 'ImpactModerateLimited', got '%s'", crq.Impact)
	}
	if crq.ChangeTiming != common.ChangeTimingNormal {
		t.Errorf("Expected ChangeTiming to be 'ChangeTimingNormal', got '%s'", crq.ChangeTiming)
	}
	if crq.Status != common.StatusPending {
		t.Errorf("Expected Status to be 'StatusPending', got '%s'", crq.Status)
	}

	// Test WorkLogs
	if len(crq.WorkLogs) != 1 {
		t.Errorf("Expected 1 WorkLog, got %d", len(crq.WorkLogs))
	} else {
		wl := crq.WorkLogs[0]
		if wl.WorkLogID != "WL001" {
			t.Errorf("Expected WorkLogID to be 'WL001', got '%s'", wl.WorkLogID)
		}
	}

	// Test Relationships
	if len(crq.Relationships) != 1 {
		t.Errorf("Expected 1 Relationship, got %d", len(crq.Relationships))
	} else {
		rel := crq.Relationships[0]
		if rel.RelationshipID != "REL001" {
			t.Errorf("Expected RelationshipID to be 'REL001', got '%s'", rel.RelationshipID)
		}
	}

	// Test Auditlogs
	if len(crq.Auditlogs) != 1 {
		t.Errorf("Expected 1 Auditlog, got %d", len(crq.Auditlogs))
	} else {
		al := crq.Auditlogs[0]
		if al.AuditlogID != "AL001" {
			t.Errorf("Expected AuditlogID to be 'AL001', got '%s'", al.AuditlogID)
		}
	}
}

func TestUtnResponseUnmarshal(t *testing.T) {
	jsonData := `{
		"content": [
			{
				"changeId": "CRQ000123",
				"universalTicketNumber": "UTN123456",
				"summary": "Test CRQ",
				"urgencyString": "2-High",
				"impactString": "3-Moderate/Limited",
				"changeTimingString": "Normal",
				"statusString": "Pending"
			}
		],
		"totalElements": 1,
		"totalPages": 1
	}`

	var utnResp UtnResponse
	err := json.Unmarshal([]byte(jsonData), &utnResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal UtnResponse: %v", err)
	}

	if len(utnResp.Content) != 1 {
		t.Errorf("Expected 1 Content item, got %d", len(utnResp.Content))
	} else {
		crq := utnResp.Content[0]
		if crq.ChangeID != "CRQ000123" {
			t.Errorf("Expected ChangeID to be 'CRQ000123', got '%s'", crq.ChangeID)
		}
		if crq.UniversalTicketNumber != "UTN123456" {
			t.Errorf("Expected UniversalTicketNumber to be 'UTN123456', got '%s'", crq.UniversalTicketNumber)
		}
		if crq.Summary != "Test CRQ" {
			t.Errorf("Expected Summary to be 'Test CRQ', got '%s'", crq.Summary)
		}
		if crq.Urgency != common.UrgencyHigh {
			t.Errorf("Expected Urgency to be 'UrgencyHigh', got '%s'", crq.Urgency)
		}
		if crq.Impact != common.ImpactModerateLimited {
			t.Errorf("Expected Impact to be 'ImpactModerateLimited', got '%s'", crq.Impact)
		}
		if crq.ChangeTiming != common.ChangeTimingNormal {
			t.Errorf("Expected ChangeTiming to be 'ChangeTimingNormal', got '%s'", crq.ChangeTiming)
		}
		if crq.Status != common.StatusPending {
			t.Errorf("Expected Status to be 'StatusPending', got '%s'", crq.Status)
		}
	}

	if utnResp.TotalElements != 1 {
		t.Errorf("Expected TotalElements to be 1, got %d", utnResp.TotalElements)
	}
	if utnResp.TotalPages != 1 {
		t.Errorf("Expected TotalPages to be 1, got %d", utnResp.TotalPages)
	}
}
