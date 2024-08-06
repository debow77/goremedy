package crq

import "goremedy/common"

// Common fields between CRQ and UTN
type CommonFields struct {
	ChangeID                       string              `json:"changeId"`
	UniversalTicketNumber          string              `json:"universalTicketNumber"`
	Summary                        string              `json:"summary"`
	RequestedStartDate             string              `json:"requestedStartDate"`
	RequestedEndDate               string              `json:"requestedEndDate"`
	ChangeManager                  string              `json:"changeManager"`
	Submitter                      string              `json:"submitter"`
	ActualEndDate                  string              `json:"actualEndDate"`
	ActualStartDate                string              `json:"actualStartDate"`
	ChangeType                     string              `json:"changeType"`
	ClientReferenceId              string              `json:"clientReferenceId"`
	ClientViewable                 string              `json:"clientViewable"`
	ClosedDate                     string              `json:"closedDate"`
	Company                        string              `json:"company"`
	CompletedDate                  string              `json:"completedDate"`
	Coordinator                    string              `json:"coordinator"`
	CorporateId                    string              `json:"corporateId"`
	IntegrationId                  string              `json:"integrationId"`
	LastModifiedby                 string              `json:"lastModifiedby"`
	LastModifiedDate               string              `json:"lastModifiedDate"`
	LeadTime                       string              `json:"leadTime"`
	LocationCompany                string              `json:"locationCompany"`
	LocationSite                   string              `json:"locationSite"`
	Manufacturer                   string              `json:"manufacturer"`
	ModelVersion                   string              `json:"modelVersion"`
	Notes                          string              `json:"notes"`
	OperationalCategorizationTier1 string              `json:"operationalCategorizationTier1"`
	OperationalCategorizationTier2 string              `json:"operationalCategorizationTier2"`
	OperationalCategorizationTier3 string              `json:"operationalCategorizationTier3"`
	OriginationDate                string              `json:"originationDate"`
	PerformanceRating              string              `json:"performanceRating"`
	PortalSolution                 string              `json:"portalSolution"`
	PortalSolutionFamily           string              `json:"portalSolutionFamily"`
	ProductCategorizationTier1     string              `json:"productCategorizationTier1"`
	ProductCategorizationTier2     string              `json:"productCategorizationTier2"`
	ProductCategorizationTier3     string              `json:"productCategorizationTier3"`
	ProductName                    string              `json:"productName"`
	RequestedBy                    string              `json:"requestedBy"`
	RequestId                      string              `json:"requestId"`
	Reviewer                       string              `json:"reviewer"`
	ReviewerLogin                  string              `json:"reviewerLogin"`
	SubmitDate                     string              `json:"submitDate"`
	TargetDate                     string              `json:"targetDate"`
	Template                       string              `json:"template"`
	VendorCompany                  string              `json:"vendorCompany"`
	VendorGroup                    string              `json:"vendorGroup"`
	VendorOrganization             string              `json:"vendorOrganization"`
	VendorTicketNumber             string              `json:"vendorTicketNumber"`
	ScheduledEndDate               string              `json:"scheduledEndDate"`
	ScheduledStartDate             string              `json:"scheduledStartDate"`
	Urgency                        common.Urgency      `json:"urgencyString"`
	Impact                         common.Impact       `json:"impactString"`
	ChangeTiming                   common.ChangeTiming `json:"changeTimingString"`
	Status                         common.Status       `json:"statusString"`
}

// CRQ represents a Change Request
type CRQResponse struct {
	CommonFields
	WorkLogs      []WorkLog      `json:"workLogs"`
	Relationships []Relationship `json:"relationships"`
	Auditlogs     []Auditlog     `json:"auditlogs"`
}

// UtnResponse represents the response when querying by UTN
type UtnResponse struct {
	Content []struct {
		CommonFields
	} `json:"content"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
}

// WorkLog represents a work log entry
type WorkLog struct {
	WorkLogID  string `json:"workLogId"`
	ChangeID   string `json:"changeId"`
	Submitter  string `json:"submitter"`
	SubmitDate string `json:"workLogSubmitDate"`
	Notes      string `json:"notes"`
}

// Relationship represents a relationship between changes
type Relationship struct {
	RelationshipID   string `json:"relationshipId"`
	ChangeID         string `json:"changeId"`
	RelatedTo        string `json:"relatedTo"`
	RelationshipType string `json:"relationshipType"`
}

// Auditlog represents an audit log entry
type Auditlog struct {
	AuditlogID string `json:"auditlogId"`
	Submitter  string `json:"submitter"`
	CreateDate string `json:"createDate"`
	Log        string `json:"log"`
}
