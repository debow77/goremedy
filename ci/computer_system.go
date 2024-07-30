package ci

import (
	"net/url"
	"strings"
	"unicode"
)

func (cg *clientGroup) GetComputerSystemByFqdn(fqdn string) (*ConfigurationItem, error) {
	return cg.getConfigurationItem("computer-systems", url.Values{"hostName": {fqdn}})
}

func (cg *clientGroup) GetComputerSystemById(id string) (*ConfigurationItem, error) {
	return cg.getConfigurationItem("computer-systems", url.Values{"instanceId": {id}})
}

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

func (cg *clientGroup) GetComputerSystems(company string, queryFilters ...map[string]string) ([]*Relationship, error) {
	if len(queryFilters) == 0 {
		queryFilters = []map[string]string{{}} // default to an empty filter map
	}

	params := url.Values{
		"relationship.markAsDeleted": {"No"},
		"source.classId":             {"BMC.CORE:CERN_DOMAIN"},
		"source.company":             {company},
		"destination.hostNameExists": {"true"},
	}

	retryFields := []string{"domain", "os", "fqdn", "usage", "notUsage"}

	for key, value := range queryFilters[0] {
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

	computerSystems, err := cg.getRelationships("assets/-/relationships", params)
	if err != nil {
		return nil, err
	}

	if len(computerSystems) == 0 {
		// Retry logic for each field
		for _, field := range retryFields {
			if value, ok := queryFilters[0][field]; ok {
				// Try with first character uppercase
				params.Set(getParamKey(field), capitalizeFirstLetter(value))
				computerSystems, err = cg.getRelationships("assets/-/relationships", params)
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

				// Reset the parameter to original value
				params.Set(getParamKey(field), value)
			}
		}
	}

	return computerSystems, nil
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

// if len(relationships) > 0 {
// 	fmt.Printf("First relationship: %+v\n", relationships[0])
// }
// if len(relationships) > 0 {
// 	rel := relationships[0]
// 	log.Printf("First relationship: Source=%+v, Destination=%+v", *rel.Source, *rel.Destination)
// }
// panic("here")
// var computerSystems []*ConfigurationItem
// var computerSystems []*ConfigurationItem

// for _, rel := range relationships {
// 	if rel.Destination != nil {
// 		computerSystems = append(computerSystems, rel.Destination)
// 	}
// }
