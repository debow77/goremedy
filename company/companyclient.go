package company

import (
	"encoding/json"
	"fmt"
	"goremedy/interfaces"
	"log/slog"
	"net/url"
	"strings"
	"time"
)

// Company represents a company entity
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

// ClientGroup is an interface for a company client group
type ClientGroup interface {
	GetCompany(mnemonics []string) ([]*Company, int, error)
	GetCernerworks(mnemonics []string) ([]*Company, int, error)
}

type clientGroup struct {
	client       interfaces.RapidClientInterface
	companyQuery *queryClient
}

// NewClientGroup creates a new company client group instance
func NewClientGroup(client interfaces.RapidClientInterface) (ClientGroup, error) {
	return &clientGroup{
		client:       client,
		companyQuery: newQueryClient(client),
	}, nil
}

type queryClient struct {
	client interfaces.RapidClientInterface
}

func newQueryClient(client interfaces.RapidClientInterface) *queryClient {
	return &queryClient{client: client}
}

// GetCompany returns a list of companies by mnemonic
func (cg *clientGroup) GetCompany(mnemonics []string) ([]*Company, int, error) {
	return cg.companyQuery.getClientCompanies(mnemonics, nil)
}

// GetCernerworks returns a list of Cernerworks companies by mnemonic
func (cg *clientGroup) GetCernerworks(mnemonics []string) ([]*Company, int, error) {
	filters := map[string]string{"mnemonic": "_"}
	return cg.companyQuery.getClientCompanies(mnemonics, filters)
}

// getClientCompanies returns a list of companies by mnemonic and filters
func (qc *queryClient) getClientCompanies(mnemonics []string, filters map[string]string) ([]*Company, int, error) {
	params := url.Values{
		"companyTypeIn": {"Customer"},
		"statusIn":      {"1"},
	}

	if len(mnemonics) > 0 {
		params.Set("mnemonicIn", strings.Join(mnemonics, "|"))
	}

	return qc.getPaginated("companies", params, filters)
}

// getPaginated returns a list of companies by URL path, params, and filters
func (qc *queryClient) getPaginated(urlPath string, params url.Values, filters map[string]string) ([]*Company, int, error) {
	slog.Debug("Getting companies", "params", params, "filters", filters)

	resp, statusCode, err := qc.client.GetPaginated(qc.getPath(), urlPath, params)
	if err != nil {
		return nil, statusCode, err
	}

	var companies []*Company
	for _, raw := range resp {
		var company Company
		if err := json.Unmarshal(raw, &company); err != nil {
			return nil, statusCode, fmt.Errorf("failed to unmarshal company: %w", err)
		}
		if qc.matchesFilters(&company, filters) {
			companies = append(companies, &company)
		}
	}

	return companies, statusCode, nil
}

// matchesFilters checks if a company matches the given filters
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
		}
	}
	return true
}

// getPath returns the base URL path for the query client
func (qc *queryClient) getPath() string {
	return "remedy-company-query-svc/v1/"
}
