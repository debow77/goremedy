package ci

import (
	"fmt"
	"goremedy"
)

func GetDomains() {
	company := "CERN_CWIM-CernerWorks Technology Improvement"

	fmt.Println("\nGet Company Domains example:", company)

	client, err := goremedy.NewRemedyClient(goremedy.RemedyClientConfig{})
	if err != nil {
		fmt.Printf("Error creating RemedyClient: %v\n", err)
		return
	}

	ciGroup := client.GetCIClientGroup()
	domains, err := ciGroup.GetDomains(company)
	if err != nil {
		fmt.Printf("Error getting domains: %v\n", err)
		return
	}
	fmt.Printf("Found %d domains\n", len(domains))

	// Print domain names
	for _, domain := range domains {
		fmt.Printf("Domain: %s\n", domain.Name)
	}

}

/* Example output of the above code

Get Company Domains example: CERN_CWIM-CernerWorks Technology Improvement
Found 36 domains
Domain: aprod
Domain: hprod
Domain: hawk
Domain: kart
Domain: kartp
Domain: prodl
Domain: Prod
Domain: DIDOMAIN_CWXTE
Domain: DIDOMAIN_IPDEV
Domain: DBSVCS_TRNSTN
Domain: BOE_LAB
Domain: BOE_LAB_DEV
Domain: goldweb
Domain: DBSERVE_SQLSERVER
Domain: irnhd
Domain: PRODL
Domain: APROD
Domain: Zabbix
Domain: testr
Domain: DBSTEST
Domain: frnzy
Domain: RMAN
Domain: webapps
Domain: LS_RCAT_NP
Domain: LS_RCAT_PROD
Domain: KC_RCAT_NP
Domain: KC_RCAT_PROD
Domain: prodl dr
Domain: TIAIX
Domain: TIRHEL
Domain: T_ora
Domain: fpp
Domain: TILLY
Domain: TVPROD
Domain: prodl
Domain: HINES

*/
