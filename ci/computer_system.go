package ci

import (
	"net/url"
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

func (cg *clientGroup) GetComputerSystems(company string, queryFilters map[string]string) ([]*ConfigurationItem, error) {
	params := url.Values{
		"relationship.markAsDeleted": {"No"},
		"source.classId":             {"BMC.CORE:CERN_DOMAIN"},
		"source.company":             {company},
		"destination.hostNameExists": {"true"},
	}

	for key, value := range queryFilters {
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

	// return cg.getRelationships("assets/-/relationships", params)
	relationships, err := cg.getRelationships("assets/-/relationships", params)
	if err != nil {
		return nil, err
	}

	var computerSystems []*ConfigurationItem
	for _, rel := range relationships {
		if rel.Destination != nil {
			computerSystems = append(computerSystems, rel.Destination)
		}
	}

	return computerSystems, nil
}
