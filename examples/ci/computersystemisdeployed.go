package ci

import (
	"fmt"
	"goremedy"
	"log"
)

func ComputerSystemIsDeployed() {
	node := "cwxtimdbus01.cernerasp.com"

	fmt.Println("\nChecking if node:", node, "is deployed")

	client, err := goremedy.NewRemedyClient(goremedy.RemedyClientConfig{})
	if err != nil {
		fmt.Printf("Error creating RemedyClient: %v\n", err)
		return
	}

	csGroup := client.GetCIClientGroup()

	status, err := csGroup.ComputerSystemIsDeployed(node)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(node, "Is Deployed:", status)
}

/* Example output of the above code

cwxtimdbus01.cernerasp.com Is Deployed: true

*/
