package ci

import (
	"net/url"
	"strings"
	"unicode"
)

// GetComputerSystemByFqdn retrieves a computer system by its FQDN
func (cg *clientGroup) GetComputerSystemByFqdn(fqdn string) (*ConfigurationItem, error) {
	return cg.getConfigurationItem("computer-systems", url.Values{"hostName": {fqdn}})
}

// GetComputerSystemById retrieves a computer system by its ID
func (cg *clientGroup) GetComputerSystemById(id string) (*ConfigurationItem, error) {
	return cg.getConfigurationItem("computer-systems", url.Values{"instanceId": {id}})
}

// ComputerSystemIsDeployed checks if a computer system is deployed
func (cg *clientGroup) ComputerSystemIsDeployed(fqdn string) (bool, error) {
	ci, err := cg.GetComputerSystemByFqdn(fqdn)
	if err != nil {
		return false, err
	}
	if ci == nil || ci.Status == nil {
		return false, nil
	}
	return ci.Status.Value == "Deployed", nil
}

// GetComputerSystems retrieves computer systems based on company and filters
func (cg *clientGroup) GetComputerSystems(company string, queryFilters ...map[string]string) ([]*Relationship, error) {
	params := url.Values{
		"relationship.markAsDeleted": {"No"},
		"source.classId":             {"BMC.CORE:CERN_DOMAIN"},
		"source.company":             {company},
		"destination.hostNameExists": {"true"},
	}

	if len(queryFilters) > 0 {
		applyQueryFilters(params, queryFilters[0])
	}

	computerSystems, err := cg.getRelationships("assets/-/relationships", params)
	if err != nil {
		return nil, err
	}

	if len(computerSystems) == 0 {
		return cg.retryWithCaseVariations(company, queryFilters...)
	}

	return computerSystems, nil
}

func applyQueryFilters(params url.Values, filters map[string]string) {
	for key, value := range filters {
		switch key {
		case "domain":
			params.Set("source.name", value)
		case "os":
			params.Set("destination.operatingSystemLike", value)
		case "fqdn":
			params.Set("destination.hostNameLike", value)
		case "usage":
			params.Set("destination.primaryUsageLike", value)
		case "notUsage":
			params.Set("destination.primaryUsageNotIn", value)
		}
	}
}

func (cg *clientGroup) retryWithCaseVariations(company string, queryFilters ...map[string]string) ([]*Relationship, error) {
	retryFields := []string{"domain", "os", "fqdn", "usage", "notUsage"}

	for _, field := range retryFields {
		if value, ok := queryFilters[0][field]; ok {
			params := buildRetryParams(company, field, value)

			// Try with first character uppercase
			computerSystems, err := cg.getRelationships("assets/-/relationships", params)
			if err != nil {
				return nil, err
			}
			if len(computerSystems) > 0 {
				return computerSystems, nil
			}

			// Try with all characters uppercase
			params.Set(getParamKey(field), strings.ToUpper(value))
			computerSystems, err = cg.getRelationships("assets/-/relationships", params)
			if err != nil {
				return nil, err
			}
			if len(computerSystems) > 0 {
				return computerSystems, nil
			}
		}
	}

	return nil, nil
}

func buildRetryParams(company, field, value string) url.Values {
	params := url.Values{
		"relationship.markAsDeleted": {"No"},
		"source.classId":             {"BMC.CORE:CERN_DOMAIN"},
		"source.company":             {company},
		"destination.hostNameExists": {"true"},
	}
	params.Set(getParamKey(field), capitalizeFirstLetter(value))
	return params
}

func capitalizeFirstLetter(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	return string(unicode.ToUpper(r[0])) + string(r[1:])
}

func getParamKey(field string) string {
	switch field {
	case "domain":
		return "source.name"
	case "os":
		return "destination.operatingSystemLike"
	case "fqdn":
		return "destination.hostNameLike"
	case "usage":
		return "destination.primaryUsageLike"
	case "notUsage":
		return "destination.primaryUsageNotIn"
	default:
		return ""
	}
}
