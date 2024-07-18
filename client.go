// package goremedy

// import (
// 	"goremedy/companies"

// 	"github.cerner.com/OHAIFedAutoSre/gorapid"
// )

// type CompanyClientGroup = companies.CompanyClientGroup

// // RemedyClientInterface defines the methods for RemedyClient
// type RemedyClientInterface interface {
// 	GetRapidClient() *gorapid.RapidClient
// 	GetCompanyClientGroup() *CompanyClientGroup
// }

// // RemedyClient represents the main client for all Remedy transactions
// type RemedyClient struct {
// 	Companies   *CompanyClientGroup
// 	rapidClient *gorapid.RapidClient
// }

// // NewRemedyClient creates a new RemedyClient instance
// func NewRemedyClient() (*RemedyClient, error) {
// 	rapidClient, err := gorapid.NewRapidClient()
// 	if err != nil {
// 		return nil, err
// 	}

// 	client := &RemedyClient{
// 		rapidClient: rapidClient,
// 	}

// 	client.Companies = NewCompanyClientGroup(client)

// 	return client, nil
// }

// // GetRapidClient returns the RapidClient
// func (rc *RemedyClient) GetRapidClient() *gorapid.RapidClient {
// 	return rc.rapidClient
// }

// // GetCompanyClientGroup returns the CompanyClientGroup
// func (rc *RemedyClient) GetCompanyClientGroup() *CompanyClientGroup {
// 	return rc.Companies
// }

package goremedy

import (
	"goremedy/companies"

	"github.cerner.com/OHAIFedAutoSre/gorapid"
)

type CompanyClientGroup interface {
	Get(mnemonics []string) ([]*companies.Company, error)
	GetCernerworks(mnemonics []string) ([]*companies.Company, error)
}

// RemedyClientInterface defines the methods for RemedyClient
type RemedyClientInterface interface {
	GetRapidClient() *gorapid.RapidClient
	GetCompanyClientGroup() CompanyClientGroup
}

// RemedyClient represents the main client for all Remedy transactions
type RemedyClient struct {
	Companies   CompanyClientGroup
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

	client.Companies = companies.NewCompanyClientGroup(client)

	return client, nil
}

// GetRapidClient returns the RapidClient
func (rc *RemedyClient) GetRapidClient() *gorapid.RapidClient {
	return rc.rapidClient
}

// GetCompanyClientGroup returns the CompanyClientGroup
func (rc *RemedyClient) GetCompanyClientGroup() CompanyClientGroup {
	return rc.Companies
}
