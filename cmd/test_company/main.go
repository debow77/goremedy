package main

import (
	"fmt"
	"log"

	"goremedy"
)

func main() {
	// Create a new RemedyClient
	remedyClient, err := goremedy.NewRemedyClient()
	if err != nil {
		log.Fatalf("Failed to create RemedyClient: %v", err)
	}

	// Call the Get method with a sample mnemonic
	companies, err := remedyClient.Companies.Get([]string{"cern_cwim"})
	if err != nil {
		log.Fatalf("Error fetching companies: %v", err)
	}

	// Print the retrieved companies
	for _, company := range companies {
		fmt.Printf("ID: %s, Name: %s, Mnemonic: %s\n", company.ID, company.Name, company.Mnemonic)
	}
}
