package ci

import (
	"encoding/json"
	"fmt"
	"goremedy/interfaces"
	"net/url"
	"strings"
)

// ClientGroup interface defines the methods for CI operations
type ClientGroup interface {
	GetDomains(company string) ([]*ConfigurationItem, error)
	GetComputerSystemByFqdn(fqdn string) (*ConfigurationItem, error)
	GetComputerSystemById(id string) (*ConfigurationItem, error)
	ComputerSystemIsDeployed(fqdn string) (bool, error)
	GetComputerSystems(company string, queryFilters ...map[string]string) ([]*Relationship, error)
}

type clientGroup struct {
	client interfaces.RapidClientInterface
}

// NewClientGroup creates a new CI client group
func NewClientGroup(client interfaces.RapidClientInterface) (ClientGroup, error) {
	return &clientGroup{client: client}, nil
}

func (cg *clientGroup) getPath() string {
	if strings.Contains(strings.ToLower(cg.client.BaseURL()), "staging") {
		return "remedy-asset-query-svc/v5/"
	}
	return "remedy-asset-query-svc/v1/"
}

func (cg *clientGroup) getConfigurationItems(urlPath string, params url.Values) ([]*ConfigurationItem, error) {
	responses, _, err := cg.client.GetPaginated(cg.getPath(), urlPath, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get configuration items: %w", err)
	}

	var configItems []*ConfigurationItem
	for _, resp := range responses {
		var ci ConfigurationItem
		if err := json.Unmarshal(resp, &ci); err != nil {
			return nil, fmt.Errorf("failed to unmarshal configuration item: %w", err)
		}
		configItems = append(configItems, &ci)
	}

	return configItems, nil
}

func (cg *clientGroup) getConfigurationItem(urlPath string, params url.Values) (*ConfigurationItem, error) {
	responses, _, err := cg.client.GetPaginated(cg.getPath(), urlPath, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get configuration item: %w", err)
	}

	if len(responses) == 0 {
		return nil, nil
	}

	var ci ConfigurationItem
	if err := json.Unmarshal(responses[0], &ci); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration item: %w", err)
	}

	return &ci, nil
}

func (cg *clientGroup) getRelationships(urlPath string, params url.Values) ([]*Relationship, error) {
	responses, _, err := cg.client.GetPaginated(cg.getPath(), urlPath, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get relationships: %w", err)
	}

	var relationships []*Relationship
	for _, resp := range responses {
		var rel Relationship
		if err := json.Unmarshal(resp, &rel); err != nil {
			return nil, fmt.Errorf("failed to unmarshal relationship: %w", err)
		}
		relationships = append(relationships, &rel)
	}

	return relationships, nil
}
