package inc

import "goremedy/common"

// Common fields between CRQ and UTN
type CommonFields struct {
	AlternateContactCompany                  string         `json:"alternateContactCompany"`
	AlternateContactDepartment               string         `json:"alternateContactDepartment"`
	AlternateContactEmail                    string         `json:"alternateContactEmail"`
	AlternateContactFirstName                string         `json:"alternateContactFirstName"`
	AlternateContactId                       string         `json:"alternateContactId"`
	AlternateContactLastName                 string         `json:"alternateContactLastName"`
	AlternateContactOrganization             string         `json:"alternateContactOrganization"`
	AlternateContactPeopleId                 string         `json:"alternateContactPeopleId"`
	AlternateContactPhoneNumber              string         `json:"alternateContactPhoneNumber"`
	AlternateContactSite                     string         `json:"alternateContactSite"`
	AlternateContactSiteGroup                string         `json:"alternateContactSiteGroup"`
	AlternateContactSiteId                   string         `json:"alternateContactSiteId"`
	AssignedGroup                            string         `json:"assignedGroup"`
	AssignedGroupId                          string         `json:"assignedGroupId"`
	AssignedSupportCompany                   string         `json:"assignedSupportCompany"`
	AssignedSupportOrganization              string         `json:"assignedSupportOrganization"`
	Assignee                                 string         `json:"assignee"`
	AssigneeId                               string         `json:"assigneeId"`
	AssigneeLoginId                          string         `json:"assigneeLoginId"`
	Broadcasted                              string         `json:"broadcasted"`
	CareImpactReview                         string         `json:"careImpactReview"`
	CareImpactReviewString                   string         `json:"careImpactReviewString"`
	CernerProjectOrg                         string         `json:"cernerProjectOrg"`
	Ci                                       string         `json:"ci"`
	CiTagNumber                              string         `json:"ciTagNumber"`
	Classification                           string         `json:"classification"`
	ClientReferenceId                        string         `json:"clientReferenceId"`
	ClientViewableStatus                     string         `json:"clientViewableStatus"`
	ClosedDate                               string         `json:"closedDate"`
	Company                                  string         `json:"company"`
	ContactEmail                             string         `json:"contactEmail"`
	ContactFirstName                         string         `json:"contactFirstName"`
	ContactId                                string         `json:"contactId"`
	ContactLastName                          string         `json:"contactLastName"`
	ContactOrganization                      string         `json:"contactOrganization"`
	ContactPeopleId                          string         `json:"contactPeopleId"`
	ContactPhoneNumber                       string         `json:"contactPhoneNumber"`
	Department                               string         `json:"department"`
	FullName                                 string         `json:"fullName"`
	IncidentId                               string         `json:"incidentId"`
	IncidentType                             string         `json:"incidentType"`
	IncidentTypeString                       string         `json:"incidentTypeString"`
	IntegrationId                            string         `json:"integrationId"`
	IsFed                                    string         `json:"isFed"`
	ItilClassification                       string         `json:"itilClassification"`
	LastModifiedby                           string         `json:"lastModifiedby"`
	LastModifiedDate                         string         `json:"lastModifiedDate"`
	LastResolvedDate                         string         `json:"lastResolvedDate"`
	LocationCompany                          string         `json:"locationCompany"`
	Notes                                    string         `json:"notes"`
	OperationalCategorizationTier1           string         `json:"operationalCategorizationTier1"`
	OperationalCategorizationTier2           string         `json:"operationalCategorizationTier2"`
	OperationalCategorizationTier3           string         `json:"operationalCategorizationTier3"`
	OriginationDate                          string         `json:"originationDate"`
	Owner                                    string         `json:"owner"`
	OwnerGroup                               string         `json:"ownerGroup"`
	OwnerGroupId                             string         `json:"ownerGroupId"`
	OwnerLoginId                             string         `json:"ownerLoginId"`
	OwnerSupportCompany                      string         `json:"ownerSupportCompany"`
	OwnerSupportOrganization                 string         `json:"ownerSupportOrganization"`
	PortalSolution                           string         `json:"portalSolution"`
	PortalSolutionFamily                     string         `json:"portalSolutionFamily"`
	PreviousStatus                           string         `json:"previousStatus"`
	Priority                                 string         `json:"priority"`
	PriorityString                           string         `json:"priorityString"`
	ProbableCause                            string         `json:"probableCause"`
	ProductCategorizationTier1               string         `json:"productCategorizationTier1"`
	ProductCategorizationTier2               string         `json:"productCategorizationTier2"`
	ProductCategorizationTier3               string         `json:"productCategorizationTier3"`
	ProductManufacturer                      string         `json:"productManufacturer"`
	ProductModelVersion                      string         `json:"productModelVersion"`
	ProductName                              string         `json:"productName"`
	Region                                   string         `json:"region"`
	ReportedDate                             string         `json:"reportedDate"`
	ReportedSource                           string         `json:"reportedSource"`
	ReportedSourceString                     string         `json:"reportedSourceString"`
	RequestId                                string         `json:"requestId"`
	RequiredResolutionDate                   string         `json:"requiredResolutionDate"`
	Resolution                               string         `json:"resolution"`
	ResolutionProductCategorizationTier1     string         `json:"resolutionProductCategorizationTier1"`
	ResolutionProductCategorizationTier2     string         `json:"resolutionProductCategorizationTier2"`
	ResolutionProductCategorizationTier3     string         `json:"resolutionProductCategorizationTier3"`
	ResolutionOperationalCategorizationTier1 string         `json:"resolutionOperationalCategorizationTier1"`
	ResolutionOperationalCategorizationTier2 string         `json:"resolutionOperationalCategorizationTier2"`
	ResolutionOperationalCategorizationTier3 string         `json:"resolutionOperationalCategorizationTier3"`
	ResolutionProductManufacturer            string         `json:"resolutionProductManufacturer"`
	ResolutionProductModelVersion            string         `json:"resolutionProductModelVersion"`
	ResolutionProductName                    string         `json:"resolutionProductName"`
	RespondedDate                            string         `json:"respondedDate"`
	ReviewForPHI                             string         `json:"reviewForPHI"`
	ReviewForPHIString                       string         `json:"reviewForPHIString"`
	Service                                  string         `json:"service"`
	ServiceClass                             string         `json:"serviceClass"`
	ServiceReconId                           string         `json:"serviceReconId"`
	ServiceRequestId                         string         `json:"serviceRequestId"`
	Site                                     string         `json:"site"`
	SiteGroup                                string         `json:"siteGroup"`
	StageNumber                              string         `json:"stageNumber"`
	StatusReasonString                       string         `json:"statusReasonString"`
	StatusReason                             string         `json:"statusReason"`
	SubmitDate                               string         `json:"submitDate"`
	Submitter                                string         `json:"submitter"`
	Summary                                  string         `json:"summary"`
	SupplierCompany                          string         `json:"supplierCompany"`
	SupplierGroup                            string         `json:"supplierGroup"`
	SupplierOrganization                     string         `json:"supplierOrganization"`
	SupplierTicketNumber                     string         `json:"supplierTicketNumber"`
	TargetDate                               string         `json:"targetDate"`
	Template                                 string         `json:"template"`
	TemplateId                               string         `json:"templateId"`
	UniversalTicketNumber                    string         `json:"universalTicketNumber"`
	Vip                                      string         `json:"vip"`
	VipString                                string         `json:"vipString"`
	Urgency                                  common.Urgency `json:"urgencyString"`
	Impact                                   common.Impact  `json:"impactString"`
	Status                                   common.Status  `json:"statusString"`
}

// INC represents a ...
type IncResponse struct {
	CommonFields
	WorkLogs      []WorkLog      `json:"workLogs"`
	Relationships []Relationship `json:"relationships"`
	Auditlogs     []Auditlog     `json:"auditlogs"`
}

type IncUtnResponse struct {
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
