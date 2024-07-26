package ci

type ConfigurationItem struct {
	Name                 string `json:"name"`
	Company              string `json:"company"`
	Site                 string `json:"site"`
	AssetLifeCycleStatus struct {
		Value string `json:"value"`
	} `json:"assetLifeCycleStatus"`
	MarkAsDeleted string `json:"markAsDeleted"`
	Status        *struct {
		Value string `json:"value"`
	} `json:"status,omitempty"`
}

type Relationship struct {
	Source      *ConfigurationItem `json:"source"`
	Destination *ConfigurationItem `json:"destination"`
}

var InvalidDomainNames = []string{
	"PMO Reclaim",
	"PMO Reclaims",
	"Staging",
	"Unallocated",
}

func filterDomains(domains []*ConfigurationItem) []*ConfigurationItem {
	var filteredDomains []*ConfigurationItem
	for _, domain := range domains {
		if !contains(InvalidDomainNames, domain.Name) &&
			(domain.AssetLifeCycleStatus.Value != "Disposed" || domain.AssetLifeCycleStatus.Value == "") &&
			(domain.MarkAsDeleted != "Yes" || domain.MarkAsDeleted == "") {
			filteredDomains = append(filteredDomains, domain)
		}
	}
	return filteredDomains
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func getMostCommonCompany(relationships []*Relationship) string {
	companyCounts := make(map[string]int)
	for _, rel := range relationships {
		if rel.Source != nil && rel.Source.Company != "" {
			companyCounts[rel.Source.Company]++
		}
	}
	return getMaxKey(companyCounts)
}

func getMostCommonSite(relationships []*Relationship) string {
	siteCounts := make(map[string]int)
	for _, rel := range relationships {
		if rel.Destination != nil && rel.Destination.Site != "" {
			siteCounts[rel.Destination.Site]++
		}
	}
	return getMaxKey(siteCounts)
}

func getMaxKey(m map[string]int) string {
	var maxKey string
	var maxValue int
	for k, v := range m {
		if v > maxValue {
			maxKey = k
			maxValue = v
		}
	}
	return maxKey
}
