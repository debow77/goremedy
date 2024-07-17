package remedy

import (
	"goremedy"
)

// RemedyClient represents the main client for all Remedy transactions
type RemedyClient struct {
	// ChangeRequests *CRClientGroup
	Companies   *CompanyClientGroup
	rapidClient *gorapid.RapidClient
}

// NewRemedyClient creates a new RemedyClient instance
func NewRemedyClient() (*RemedyClient, error) {
	rapidClient, err := gorapid.NewRapidClient()
	if err != nil {
		return nil, err
	}

	client := &RemedyClient{
		rapidClient: rapidClient,
	}

	// client.ChangeRequests = NewCRClientGroup(client)
	client.Companies = NewCompanyClientGroup(client)

	return client, nil
}
