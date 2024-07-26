package company

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"goremedy/internal/common"
)

type Company struct {
	CompanyId       string    `json:"companyId,omitempty"`
	RemedyCompanyId string    `json:"remedyCompanyId,omitempty"`
	Name            string    `json:"name,omitempty"`
	Region          string    `json:"region,omitempty"`
	Company         string    `json:"company,omitempty"`
	Mnemonic        string    `json:"mnemonic,omitempty"`
	FocusClient     string    `json:"focusClient,omitempty"`
	CompanyType     string    `json:"companyType,omitempty"`
	Status          string    `json:"status,omitempty"`
	ProdDataCenter  string    `json:"prodDataCenter,omitempty"`
	DrDataCenter    string    `json:"drDataCenter,omitempty"`
	CreatedDate     time.Time `json:"createdDate,omitempty"`
	ModifiedDate    time.Time `json:"modifiedDate,omitempty"`
}

type ClientGroup interface {
	GetCompany(mnemonics []string) ([]*Company, error)
	GetCernerworks(mnemonics []string) ([]*Company, error)
}

type clientGroup struct {
	client       common.RemedyClientInterface
	companyQuery *queryClient
}

func NewClientGroup(client common.RemedyClientInterface) ClientGroup {
	return &clientGroup{
		client:       client,
		companyQuery: newQueryClient(client),
	}
}

type queryClient struct {
	client common.RemedyClientInterface
}

func newQueryClient(client common.RemedyClientInterface) *queryClient {
	return &queryClient{client: client}
}

func (cg *clientGroup) GetCompany(mnemonics []string) ([]*Company, error) {
	return cg.companyQuery.getClientCompanies(mnemonics, nil)
}

func (cg *clientGroup) GetCernerworks(mnemonics []string) ([]*Company, error) {
	filters := map[string]string{"mnemonic": "_"}
	return cg.companyQuery.getClientCompanies(mnemonics, filters)
}

func (qc *queryClient) getClientCompanies(mnemonics []string, filters map[string]string) ([]*Company, error) {
	params := url.Values{
		"companyTypeIn": {"Customer"},
		"statusIn":      {"1"},
	}

	if len(mnemonics) > 0 {
		params.Set("mnemonicIn", strings.Join(mnemonics, "|"))
	}

	return qc.getPaginated("companies", params, filters)
}

func (qc *queryClient) getPaginated(urlPath string, params url.Values, filters map[string]string) ([]*Company, error) {
	rawItems, err := common.GetPaginated(qc.client, qc.getPath(), urlPath, params)
	if err != nil {
		return nil, err
	}

	var companies []*Company
	for _, raw := range rawItems {
		var company Company
		if err := json.Unmarshal(raw, &company); err != nil {
			return nil, fmt.Errorf("failed to unmarshal company: %v", err)
		}
		if qc.matchesFilters(&company, filters) {
			companies = append(companies, &company)
		}
	}

	return companies, nil
}

func (qc *queryClient) matchesFilters(company *Company, filters map[string]string) bool {
	for key, value := range filters {
		switch key {
		case "mnemonic":
			if !strings.Contains(company.Mnemonic, value) {
				return false
			}
		case "name":
			if !strings.Contains(company.Name, value) {
				return false
			}
			// Add more cases as needed for other filter types
		}
	}
	return true
}

func (qc *queryClient) getPath() string {
	return "remedy-company-query-svc/v1/"
}
