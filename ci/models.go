package ci

// ConfigurationItem represents a configuration item in the system
type ConfigurationItem struct {
	Name                   string `json:"name"`
	Company                string `json:"company"`
	Region                 string `json:"region"`
	Site                   string `json:"site"`
	Description            string `json:"description"`
	ShortDescription       string `json:"shortDescription"`
	Item                   string `json:"item"`
	ManufacturerName       string `json:"manufacturerName"`
	HostName               string `json:"hostName"`
	Model                  string `json:"model"`
	RoomRack               string `json:"roomRack"`
	Type                   string `json:"type"`
	Monitored              string `json:"monitored"`
	OperatingSystem        string `json:"operatingSystem"`
	OperatingSystemVersion string `json:"operatingSystemVersion"`
	PrimaryIp              string `json:"primaryIp"`
	PrimaryUsage           string `json:"primaryUsage"`
	SecondaryUsage         string `json:"secondaryUsage"`
	Domain                 string `json:"domain"`
	AssetLifeCycleStatus   struct {
		Value string `json:"value"`
	} `json:"assetLifeCycleStatus"`
	MarkAsDeleted string `json:"markAsDeleted"`
	InstanceId    string `json:"instanceId"`
	Status        *struct {
		Value string `json:"value"`
	} `json:"status,omitempty"`
}

// Relationship represents a relationship between configuration items
type Relationship struct {
	Source      *ConfigurationItem `json:"source"`
	Destination *ConfigurationItem `json:"destination"`
}

// InvalidDomainNames is a list of domain names that should be filtered out
var InvalidDomainNames = []string{
	"PMO Reclaim",
	"PMO Reclaims",
	"Staging",
	"Unallocated",
}

// filterDomains removes invalid domains from the list
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

// contains checks if a string is present in a slice
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// getMostCommonCompany returns the most common company from relationships
func getMostCommonCompany(relationships []*Relationship) string {
	return getMaxKey(countCompanies(relationships))
}

// getMostCommonSite returns the most common site from relationships
func getMostCommonSite(relationships []*Relationship) string {
	return getMaxKey(countSites(relationships))
}

func countCompanies(relationships []*Relationship) map[string]int {
	companyCounts := make(map[string]int)
	for _, rel := range relationships {
		if rel.Source != nil && rel.Source.Company != "" {
			companyCounts[rel.Source.Company]++
		}
	}
	return companyCounts
}

func countSites(relationships []*Relationship) map[string]int {
	siteCounts := make(map[string]int)
	for _, rel := range relationships {
		if rel.Destination != nil && rel.Destination.Site != "" {
			siteCounts[rel.Destination.Site]++
		}
	}
	return siteCounts
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
