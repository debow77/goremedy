package companies

import (
    "encoding/json"
    "fmt"
    "net/url"
    "strings"
    "time"
    "goremedy"
)

// Company represents a Remedy Company
type Company struct {
    ID           string    `json:"id,omitempty"`
    Name         string    `json:"name,omitempty"`
    Mnemonic     string    `json:"mnemonic,omitempty"`
    Type         string    `json:"type,omitempty"`
    Status       int       `json:"status,omitempty"`
    CreatedDate  time.Time `json:"createdDate,omitempty"`
    ModifiedDate time.Time `json:"modifiedDate,omitempty"`
}

// CompanyClientGroup contains the Remedy company clients
type CompanyClientGroup struct {
    client       goremedy.RemedyClientInterface
    companyQuery *CompanyQueryClient
}

// NewCompanyClientGroup creates a new CompanyClientGroup
func NewCompanyClientGroup(client goremedy.RemedyClientInterface) *CompanyClientGroup {
    return &CompanyClientGroup{
        client:       client,
        companyQuery: NewCompanyQueryClient(client),
    }
}

// Get retrieves companies
func (cg *CompanyClientGroup) Get(mnemonics []string) ([]*Company, error) {
    return cg.companyQuery.GetClientCompanies(mnemonics, nil)
}

// GetCernerworks retrieves CernerWorks companies
func (cg *CompanyClientGroup) GetCernerworks(mnemonics []string) ([]*Company, error) {
    filters := map[string]string{"mnemonic": "_"}
    return cg.companyQuery.GetClientCompanies(mnemonics, filters)
}

// CompanyQueryClient interacts with the Remedy Company API
type CompanyQueryClient struct {
    client goremedy.RemedyClientInterface
}

// NewCompanyQueryClient creates a new CompanyQueryClient
func NewCompanyQueryClient(client goremedy.RemedyClientInterface) *CompanyQueryClient {
    return &CompanyQueryClient{
        client: client,
    }
}

// GetClientCompanies gathers client companies
func (cq *CompanyQueryClient) GetClientCompanies(mnemonics []string, filters map[string]string) ([]*Company, error) {
    params := url.Values{
        "companyTypeIn": {"Customer"},
        "statusIn":      {"1"},
    }

    if len(mnemonics) > 0 {
        params.Set("mnemonicIn", strings.Join(mnemonics, "|"))
    }

    companies, err := cq.getPaginated("companies", params, filters)
    if err != nil {
        return nil, err
    }

    return companies, nil
}

// getPaginated performs paginated HTTP GET requests against the base API URL for the company client
func (cq *CompanyQueryClient) getPaginated(urlPath string, params url.Values, filters map[string]string) ([]*Company, error) {
    params.Set("size", "1000")
    params.Set("page", "0")

    var allCompanies []*Company
    totalPages := 1

    for page := 0; page < totalPages; page++ {
        params.Set("page", fmt.Sprintf("%d", page))

        resp, err := cq.client.GetRapidClient().Get(cq.getPath()+urlPath, params)
        if err != nil {
            return nil, err
        }

        var pageResp struct {
            Content    []*Company `json:"content"`
            TotalPages int        `json:"totalPages"`
        }

        if err := json.Unmarshal(resp, &pageResp); err != nil {
            return nil, err
        }

        filteredCompanies := cq.filterCompanies(pageResp.Content, filters)
        allCompanies = append(allCompanies, filteredCompanies...)

        totalPages = pageResp.TotalPages
    }

    return allCompanies, nil
}

// filterCompanies filters the companies based on the provided filters
func (cq *CompanyQueryClient) filterCompanies(companies []*Company, filters map[string]string) []*Company {
    if len(filters) == 0 {
        return companies
    }

    var filteredCompanies []*Company
    for _, company := range companies {
        include := true
        for key, value := range filters {
            switch key {
            case "mnemonic":
                if !strings.Contains(company.Mnemonic, value) {
                    include = false
                }
            }
        }
        if include {
            filteredCompanies = append(filteredCompanies, company)
        }
    }

    return filteredCompanies
}

// getPath returns the base path for company-related API endpoints
func (cq *CompanyQueryClient) getPath() string {
    return "remedy-company-query-svc/v1/"
}
