package crq

import (
	"fmt"
	"goremedy"
)

func GetCrq() {
	fmt.Println("\nGet CRQ usage example:")

	changeID := "CRQ000005174722"
	utn := "452335879"

	remedyClient, err := goremedy.NewRemedyClient()
	if err != nil {
		fmt.Printf("Error creating RemedyClient: %v\n", err)
		return
	}

	crqGroup := remedyClient.GetCRQClientGroup()
	crq, err := crqGroup.Get(changeID)

	if err != nil {
		fmt.Printf("Error getting crq: %v\n", err)
		return
	}

	if crq != nil {
		fmt.Printf("CRQ found:\n")
		fmt.Printf("  ChangeID: %s\n", crq.ChangeID)
		fmt.Printf("  Urgency: %s\n", crq.Urgency)
		fmt.Printf("  Impact: %s\n", crq.Impact)
		fmt.Printf("  UniversalTicketNumber: %s\n", crq.UniversalTicketNumber)
		fmt.Printf("  Status: %s\n", crq.Status)
		fmt.Printf("  Summary: %s\n", crq.Summary)
		fmt.Printf("  RequestedStartDate: %s\n", crq.RequestedStartDate)
		fmt.Printf("  RequestedEndDate: %s\n", crq.RequestedEndDate)
		fmt.Printf("  ChangeManager: %s\n", crq.ChangeManager)
		fmt.Printf("  Submitter: %s\n", crq.Submitter)

		// Just different printing methods I was playing with....
		// fmt.Printf("CRQ found: %+v\n", crq)
		// b, _ := json.MarshalIndent(crq, "", "  ")
		// fmt.Println(string(b))
	} else {
		fmt.Println("No CRQ found for the given ID.")
	}

	utn_crq, err := crqGroup.GetByUtn(utn)
	if err != nil {
		fmt.Printf("\nError getting crq by utn: %v\n", err)
		return
	}

	if utn_crq != nil {
		fmt.Printf("\nCRQ by Utn found:\n")
		fmt.Printf("  ChangeID: %s\n", utn_crq.Content[0].ChangeID)
		fmt.Printf("  Urgency: %s\n", utn_crq.Content[0].Urgency)
		fmt.Printf("  Impact: %s\n", utn_crq.Content[0].Impact)
		fmt.Printf("  UniversalTicketNumber: %s\n", utn_crq.Content[0].UniversalTicketNumber)
		fmt.Printf("  Status: %s\n", utn_crq.Content[0].Status)
		fmt.Printf("  Summary: %s\n", utn_crq.Content[0].Summary)
		fmt.Printf("  RequestedStartDate: %s\n", utn_crq.Content[0].RequestedStartDate)
		fmt.Printf("  RequestedEndDate: %s\n", utn_crq.Content[0].RequestedEndDate)
		fmt.Printf("  ChangeManager: %s\n", utn_crq.Content[0].ChangeManager)
		fmt.Printf("  Submitter: %s\n", utn_crq.Content[0].Submitter)
	} else {
		fmt.Println("No CRQ found for the given ID.")
	}
}

/* Example output of the above code

Get CRQ usage example:
CRQ found:
  ChangeID: CRQ000005174722
  Urgency: 4-Low
  Impact: 4-Minor/Localized
  UniversalTicketNumber: 452335879
  Status: 0
  Summary: CERN_DARS - WebSphere ODR cert renewal testing.
  RequestedStartDate:
  RequestedEndDate:
  ChangeManager: Jon Boldt
  Submitter: dp7168

CRQ by Utn found:
  ChangeID: CRQ000005174722
  Urgency: 4-Low
  Impact: 4-Minor/Localized
  UniversalTicketNumber: 452335879
  Status: 0
  Summary: CERN_DARS - WebSphere ODR cert renewal testing.
  RequestedStartDate:
  RequestedEndDate:
  ChangeManager: Jon Boldt
  Submitter: dp7168

*/
