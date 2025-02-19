package company

import (
	"fmt"
	"goremedy"
)

func GetCompany2() {
	fmt.Println("\nGetCompany - Multi Company usage example:")

	// This would be used overriding the LogLevel set in client.go
	// config := goremedy.RemedyClientConfig{
	// 	LogLevel: "DEBUG", // Set log level to DEBUG
	// }
	// remedyClient, err := goremedy.NewRemedyClient(config)

	// This would be used if not overriding the LogLevel set in client.go
	remedyClient, err := goremedy.NewRemedyClient(goremedy.RemedyClientConfig{})

	if err != nil {
		fmt.Printf("Error creating RemedyClient: %v\n", err)
		return
	}

	companyGroup := remedyClient.GetCompanyClientGroup()

	// Get companies by mnemonics
	companies, status, err := companyGroup.GetCompany([]string{"CERN_CWIM", "CERN_KUBE"})
	fmt.Println("Status code:", status)
	if err != nil {
		fmt.Printf("Error getting companies: %v\n", err)
		return
	}
	fmt.Printf("Found %d companies\n", len(companies))

	// Print the results
	if len(companies) > 0 {
		for _, company := range companies {
			fmt.Printf("Company found:\n")
			fmt.Printf("  CompanyId: %s\n", company.CompanyId)
			fmt.Printf("  RemedyCompanyId: %s\n", company.RemedyCompanyId)
			fmt.Printf("  Name: %s\n", company.Name)
			fmt.Printf("  Region: %s\n", company.Region)
			fmt.Printf("  Company: %s\n", company.Company)
			fmt.Printf("  Mnemonic: %s\n", company.Mnemonic)
			fmt.Printf("  FocusClient: %s\n", company.FocusClient)
			fmt.Printf("  CompanyType: %s\n", company.CompanyType)
			fmt.Printf("  ProdDataCenter: %s\n", company.ProdDataCenter)
			fmt.Printf("  DrDataCenter: %s\n", company.DrDataCenter)
			fmt.Printf("  Status: %s\n", company.Status)
			fmt.Printf("  Created Date: %s\n", company.CreatedDate)
			fmt.Printf("  Modified Date: %s\n", company.ModifiedDate)
		}
	} else {
		fmt.Println("No companies found for the given mnemonics.")
	}

}

/* Example output of the above code


Multi Company usage example:
2024/07/29 08:10:50 WARN Log level warn
Status code: 200
Found 2 companies
Company found:
  CompanyId: CPY000000139896
  RemedyCompanyId: 0510
  Name: CernerWorks Technology Improvement
  Region: Internal
  Company: CERN_CWIM-CernerWorks Technology Improvement
  Mnemonic: CERN_CWIM
  FocusClient:
  CompanyType: Customer
  ProdDataCenter: CTC-LS-III
  DrDataCenter: CTC-KC-I
  Status: 1
  Created Date: 0001-01-01 00:00:00 +0000 UTC
  Modified Date: 0001-01-01 00:00:00 +0000 UTC
Company found:
  CompanyId: CPY000000192673
  RemedyCompanyId: 247803
  Name: CWx Engineering
  Region: Internal
  Company: CERN_KUBE-CWx Engineering
  Mnemonic: CERN_KUBE
  FocusClient:
  CompanyType: Customer
  ProdDataCenter: CTC-LS-V
  DrDataCenter:
  Status: 1
  Created Date: 0001-01-01 00:00:00 +0000 UTC
  Modified Date: 0001-01-01 00:00:00 +0000 UTC


*/
