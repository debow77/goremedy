package main

import (
	"fmt"
	"goremedy"
)

func main() {

	fmt.Println("Company usage example:")
	ExampleUsage()
}

func ExampleUsage() {
	client, err := goremedy.NewRemedyClient(goremedy.RemedyClientConfig{})
	if err != nil {
		fmt.Printf("Error creating RemedyClient: %v\n", err)
		return
	}
	ciGroup := client.GetCIClientGroup()

	domains, err := ciGroup.GetDomains("CERN_CWIM-CernerWorks Technology Improvement")
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

// Example output
// Company usage example:
// Found 36 domains
// Domain: aprod
// Domain: hprod
// Domain: hawk
// Domain: kart
// Domain: kartp
// Domain: prodl
// Domain: Prod
// Domain: DIDOMAIN_CWXTE
// Domain: DIDOMAIN_IPDEV
// Domain: DBSVCS_TRNSTN
// Domain: BOE_LAB
// Domain: BOE_LAB_DEV
// Domain: goldweb
// Domain: DBSERVE_SQLSERVER
// Domain: irnhd
// Domain: PRODL
// Domain: APROD
// Domain: Zabbix
// Domain: testr
// Domain: DBSTEST
// Domain: frnzy
// Domain: RMAN
// Domain: webapps
// Domain: LS_RCAT_NP
// Domain: LS_RCAT_PROD
// Domain: KC_RCAT_NP
// Domain: KC_RCAT_PROD
// Domain: prodl dr
// Domain: TIAIX
// Domain: TIRHEL
// Domain: T_ora
// Domain: fpp
// Domain: TILLY
// Domain: TVPROD
// Domain: prodl
// Domain: HINES
