package ci

import (
	"encoding/json"
	"fmt"
	"goremedy"
	"log"
)

func GetComputerSystems() {
	company := "CERN_CWIM-CernerWorks Technology Improvement"

	fmt.Println("\nGetComputerSystems:", company)

	client, err := goremedy.NewRemedyClient(goremedy.RemedyClientConfig{})
	if err != nil {
		fmt.Printf("Error creating RemedyClient: %v\n", err)
		return
	}

	csGroup := client.GetCIClientGroup()

	filter := map[string]string{
		"fqdn": "CWXTIMDBUS01",
	}

	// Additional filters
	// filter := map[string]string{
	// 	"os": "Linux",
	// }
	data, err := csGroup.GetComputerSystems(company, filter)

	// Without filters
	// data, err := csGroup.GetComputerSystems("CERN_CWIM-CernerWorks Technology Improvement")
	if err != nil {
		log.Println(err)
		return
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(jsonBytes))
}

/* Example output of the above code with fqdn filter

[
  {
    "source": {
      "name": "prodl",
      "company": "CERN_CWIM-CernerWorks Technology Improvement",
      "region": "",
      "site": "",
      "description": "",
      "shortDescription": "",
      "item": "",
      "manufacturerName": "",
      "hostName": "",
      "model": "",
      "roomRack": "",
      "type": "",
      "monitored": "",
      "operatingSystem": "",
      "operatingSystemVersion": "",
      "primaryIp": "",
      "primaryUsage": "",
      "secondaryUsage": "",
      "domain": "",
      "assetLifeCycleStatus": {
        "value": "Deployed"
      },
      "markAsDeleted": "No",
      "instanceId": "IDGIW2WRIDVLKANTYRVFQXPVDKNPLF",
      "status": {
        "value": "Deployed"
      }
    },
    "destination": {
      "name": "CWXTIMDBUS01.CERNERASP.COM",
      "company": "TASP_MO-4584",
      "region": "US",
      "site": "CTC-LS-III",
      "description": "cwxtimdbus01: cwxtimdbus01.cernerasp.com",
      "shortDescription": "",
      "item": "Virtual",
      "manufacturerName": "VMware",
      "hostName": "CWXTIMDBUS01.CERNERASP.COM",
      "model": "VMware Virtual Platform",
      "roomRack": "12",
      "type": "Server",
      "monitored": "",
      "operatingSystem": "Oracle Linux",
      "operatingSystemVersion": "7.9",
      "primaryIp": "7.40.8.106",
      "primaryUsage": "CareAware iBus",
      "secondaryUsage": "",
      "domain": "",
      "assetLifeCycleStatus": {
        "value": "Deployed"
      },
      "markAsDeleted": "No",
      "instanceId": "OI-D431521832A511E8BDC5005056A21B2A",
      "status": {
        "value": "Deployed"
      }
    }
  }
]

*/
