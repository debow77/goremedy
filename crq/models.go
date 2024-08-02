package crq

import "encoding/json"

type Urgency string
type Impact string
type ChangeTiming string
type Status string

const (
	UrgencyUnknown  Urgency = "Unknown"
	UrgencyCritical Urgency = "1-Critical"
	UrgencyHigh     Urgency = "2-High"
	UrgencyMedium   Urgency = "3-Medium"
	UrgencyLow      Urgency = "4-Low"

	ImpactUnknown             Impact = "Unknown"
	ImpactExtensiveWidespread Impact = "1-Extensive/Widespread"
	ImpactSignificantLarge    Impact = "2-Significant/Large"
	ImpactModerateLimited     Impact = "3-Moderate/Limited"
	ImpactMinorLocalized      Impact = "4-Minor/Localized"

	ChangeTimingUnknown   ChangeTiming = "Unknown"
	ChangeTimingEmergency ChangeTiming = "Emergency"
	ChangeTimingNormal    ChangeTiming = "Normal"
	ChangeTimingStandard  ChangeTiming = "Standard"

	StatusDraft                    Status = "Draft"
	StatusCancelled                Status = "Cancelled"
	StatusClosed                   Status = "Closed"
	StatusCompleted                Status = "Completed"
	StatusImplementationInProgress Status = "Implementation In Progress"
	StatusPending                  Status = "Pending"
	StatusPlanningInProgress       Status = "Planning In Progress"
	StatusRejected                 Status = "Rejected"
	StatusRequestForAuthorization  Status = "Request For Authorization"
	StatusScheduled                Status = "Scheduled"
	StatusScheduledForApproval     Status = "Scheduled For Approval"
	StatusScheduledForReview       Status = "Scheduled For Review"
)

// Common fields between CRQ and UTN
type CommonFields struct {
	ChangeID                       string       `json:"changeId"`
	UniversalTicketNumber          string       `json:"universalTicketNumber"`
	Summary                        string       `json:"summary"`
	RequestedStartDate             string       `json:"requestedStartDate"`
	RequestedEndDate               string       `json:"requestedEndDate"`
	ChangeManager                  string       `json:"changeManager"`
	Submitter                      string       `json:"submitter"`
	ActualEndDate                  string       `json:"actualEndDate"`
	ActualStartDate                string       `json:"actualStartDate"`
	ChangeType                     string       `json:"changeType"`
	ClientReferenceId              string       `json:"clientReferenceId"`
	ClientViewable                 string       `json:"clientViewable"`
	ClosedDate                     string       `json:"closedDate"`
	Company                        string       `json:"company"`
	CompletedDate                  string       `json:"completedDate"`
	Coordinator                    string       `json:"coordinator"`
	CorporateId                    string       `json:"corporateId"`
	IntegrationId                  string       `json:"integrationId"`
	LastModifiedby                 string       `json:"lastModifiedby"`
	LastModifiedDate               string       `json:"lastModifiedDate"`
	LeadTime                       string       `json:"leadTime"`
	LocationCompany                string       `json:"locationCompany"`
	LocationSite                   string       `json:"locationSite"`
	Manufacturer                   string       `json:"manufacturer"`
	ModelVersion                   string       `json:"modelVersion"`
	Notes                          string       `json:"notes"`
	OperationalCategorizationTier1 string       `json:"operationalCategorizationTier1"`
	OperationalCategorizationTier2 string       `json:"operationalCategorizationTier2"`
	OperationalCategorizationTier3 string       `json:"operationalCategorizationTier3"`
	OriginationDate                string       `json:"originationDate"`
	PerformanceRating              string       `json:"performanceRating"`
	PortalSolution                 string       `json:"portalSolution"`
	PortalSolutionFamily           string       `json:"portalSolutionFamily"`
	ProductCategorizationTier1     string       `json:"productCategorizationTier1"`
	ProductCategorizationTier2     string       `json:"productCategorizationTier2"`
	ProductCategorizationTier3     string       `json:"productCategorizationTier3"`
	ProductName                    string       `json:"productName"`
	RequestedBy                    string       `json:"requestedBy"`
	RequestId                      string       `json:"requestId"`
	Reviewer                       string       `json:"reviewer"`
	ReviewerLogin                  string       `json:"reviewerLogin"`
	SubmitDate                     string       `json:"submitDate"`
	TargetDate                     string       `json:"targetDate"`
	Template                       string       `json:"template"`
	VendorCompany                  string       `json:"vendorCompany"`
	VendorGroup                    string       `json:"vendorGroup"`
	VendorOrganization             string       `json:"vendorOrganization"`
	VendorTicketNumber             string       `json:"vendorTicketNumber"`
	ScheduledEndDate               string       `json:"scheduledEndDate"`
	ScheduledStartDate             string       `json:"scheduledStartDate"`
	Urgency                        Urgency      `json:"urgencyString"`
	Impact                         Impact       `json:"impactString"`
	ChangeTiming                   ChangeTiming `json:"changeTimingString"`
	Status                         Status       `json:"statusString"`
}

func (u *Urgency) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*u = Urgency(s)
	return nil
}

func (i *Impact) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*i = Impact(s)
	return nil
}

func (ct *ChangeTiming) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*ct = ChangeTiming(s)
	return nil
}

func (ct *Status) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*ct = Status(s)
	return nil
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
