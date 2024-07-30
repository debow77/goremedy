package ci

import (
	"encoding/json"
	"fmt"
	"goremedy"
	"log"
)

func GetComputerSystemByFqdn() {
	node := "cwxtimdbus01.cernerasp.com"

	fmt.Println("\nGetComputerSystemByFqdn:", node)

	client, err := goremedy.NewRemedyClient(goremedy.RemedyClientConfig{})
	if err != nil {
		fmt.Printf("Error creating RemedyClient: %v\n", err)
		return
	}

	csGroup := client.GetCIClientGroup()

	server_info, err := csGroup.GetComputerSystemByFqdn(node)
	if err != nil {
		log.Println(err)
		return
	}

	jsonBytes, err := json.MarshalIndent(server_info, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(jsonBytes))
}

/* Example output of the above code

Name: CWXTIMDBUS01.CERNERASP.COM
{
  "name": "CWXTIMDBUS01.CERNERASP.COM",
  "company": "TASP_MO-4584",
  "region": "US",
  "site": "CTC-LS-III",
  "description": "cwxtimdbus01: cwxtimdbus01.cernerasp.com",
  "shortDescription": "cwxtimdbus01",
  "item": "Virtual",
  "manufacturerName": "VMware",
  "model": "VMware Virtual Platform",
  "roomRack": "12",
  "type": "Server",
  "monitored": "",
  "operatingSystem": "Oracle Linux",
  "operatingSystemVersion": "7.9",
  "primaryIp": "7.40.8.106",
  "primaryUsage": "CareAware iBus",
  "secondaryUsage": "",
  "domain": "cernerasp.com",
  "assetLifeCycleStatus": {
    "value": "Deployed"
  },
  "markAsDeleted": "No",
  "instanceId": "OI-D431521832A511E8BDC5005056A21B2A",
  "status": {
    "value": "Deployed"
  }
}

*/
