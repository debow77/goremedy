package common

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
