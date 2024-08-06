package inc

import (
	"fmt"
	"goremedy"
)

func GetInc() {
	fmt.Println("\nGet INC usage example:")

	incNumber := "INC000029331406"
	utn := "432062289"

	remedyClient, err := goremedy.NewRemedyClient()
	if err != nil {
		fmt.Printf("Error creating RemedyClient: %v\n", err)
		return
	}

	incGroup := remedyClient.GetINCClientGroup()

	inc, err := incGroup.Get(incNumber)

	if err != nil {
		fmt.Printf("Error getting inc: %v\n", err)
		return
	}

	if inc != nil {
		fmt.Printf("INC found:\n")
		fmt.Printf("  IncidentId: %s\n", inc.IncidentId)
		fmt.Printf("  Company: %s\n", inc.Company)
		fmt.Printf("  ContactFirstName: %s\n", inc.ContactFirstName)
		fmt.Printf("  ContactId: %s\n", inc.ContactId)
		fmt.Printf("  UniversalTicketNumber: %s\n", inc.UniversalTicketNumber)
		fmt.Printf("  IsFed: %s\n", inc.IsFed)
		fmt.Printf("  Template: %s\n", inc.Template)
		fmt.Printf("  Region: %s\n", inc.Region)
		fmt.Printf("  Urgency: %s\n", inc.Urgency)
		fmt.Printf("  Impact: %s\n", inc.Impact)
	} else {
		fmt.Println("No INC found for the given ID.")
	}

	utn_inc, err := incGroup.GetByUtn(utn)
	if err != nil {
		fmt.Printf("\nError getting crq by utn: %v\n", err)
		return
	}

	if utn_inc != nil {
		fmt.Printf("\nINC by Utn found:\n")
		fmt.Printf("  Status: %s\n", utn_inc.Content[0].Status)
		fmt.Printf("  Company: %s\n", utn_inc.Content[0].Company)
		fmt.Printf("  ContactFirstName: %s\n", utn_inc.Content[0].ContactFirstName)
		fmt.Printf("  ContactId: %s\n", utn_inc.Content[0].ContactId)
		fmt.Printf("  UniversalTicketNumber: %s\n", utn_inc.Content[0].UniversalTicketNumber)
		fmt.Printf("  IsFed: %s\n", utn_inc.Content[0].IsFed)
		fmt.Printf("  Template: %s\n", utn_inc.Content[0].Template)
		fmt.Printf("  Region: %s\n", utn_inc.Content[0].Region)
		fmt.Printf("  Urgency: %s\n", utn_inc.Content[0].Urgency)
		fmt.Printf("  Impact: %s\n", utn_inc.Content[0].Impact)

	} else {
		fmt.Println("No CRQ found for the given ID.")
	}
}

/* Example output of the above code

Get INC usage example:
INC found:
  IncidentId: INC000029331406
  Company: CERN_OLYM-Cerner Olympus Systems
  ContactFirstName: Cerner
  ContactId: ASSOCIATE
  UniversalTicketNumber: 432062289
  IsFed: False
  Template:
  Region:
  Urgency: 1-Critical
  Impact: 1-Extensive/Widespread

INC by Utn found:
  Status: Cancelled
  Company: CERN_OLYM-Cerner Olympus Systems
  ContactFirstName: Cerner
  ContactId: ASSOCIATE
  UniversalTicketNumber: 432062289
  IsFed: False
  Template:
  Region:
  Urgency: 1-Critical
  Impact: 1-Extensive/Widespread

*/
