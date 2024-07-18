package main

import (
	"fmt"
	"log"
)

func Main() {
	// Create a new RemedyClient
	client, err := remedy.NewRemedyClient()
	if err != nil {
		log.Fatalf("Error creating Remedy client: %v", err)
	}

	// Call Get method with no mnemonics (to get all companies)
	companies, err := client.Companies.Get(nil)
	if err != nil {
		log.Fatalf("Error getting companies: %v", err)
	}

	// Print the results
	fmt.Printf("Found %d companies\n", len(companies))
	for _, company := range companies {
		fmt.Printf("Company: %s (Mnemonic: %s)\n", company.Name, company.Mnemonic)
	}

	// Test with specific mnemonics
	mnemonics := []string{"CERN_CWIM"}
	specificCompanies, err := client.Companies.Get(mnemonics)
	if err != nil {
		log.Fatalf("Error getting specific companies: %v", err)
	}

	fmt.Printf("\nFound %d specific companies\n", len(specificCompanies))
	for _, company := range specificCompanies {
		fmt.Printf("Company: %s (Mnemonic: %s)\n", company.Name, company.Mnemonic)
	}
}
