package ci

import (
	"net/url"
)

// GetDomains retrieves domains for a given company
func (cg *clientGroup) GetDomains(company string) ([]*ConfigurationItem, error) {
	params := url.Values{"company": {company}}
	domains, err := cg.getConfigurationItems("domains", params)
	if err != nil {
		return nil, err
	}
	return filterDomains(domains), nil
}

// TODO: Implement and test these functions
// func (cg *clientGroup) DomainHasUsage(company, domain, usage string) (bool, error) {
//     params := url.Values{
//         "relationship.markAsDeleted":   {"No"},
//         "source.name":                  {domain},
//         "source.company":               {company},
//         "destination.primaryUsageLike": {usage},
//     }
//     relationships, err := cg.getRelationships("assets/-/relationships", params)
//     if err != nil {
//         return false, err
//     }
//     return len(relationships) > 0, nil
// }

// func (cg *clientGroup) GetDomainSite(company, domain string) (string, error) {
//     params := url.Values{
//         "relationship.markAsDeleted": {"No"},
//         "source.classId":             {"BMC.CORE:CERN_DOMAIN"},
//         "source.company":             {company},
//         "source.name":                {domain},
//         "destination.hostNameExists": {"true"},
//     }
//     relationships, err := cg.getRelationships("assets/-/relationships", params)
//     if err != nil {
//         return "", err
//     }
//     return getMostCommonSite(relationships), nil
// }
