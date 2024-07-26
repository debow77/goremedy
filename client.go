package goremedy

import (
	"log"
	"os"

	"goremedy/ci"
	"goremedy/company"

	"github.cerner.com/OHAIFedAutoSre/gorapid"
)

type RemedyClientInterface interface {
	GetRapidClient() *gorapid.RapidClient
	GetCompanyClientGroup() company.ClientGroup
	GetCIClientGroup() ci.ClientGroup
}

type RemedyClient struct {
	rapidClient        *gorapid.RapidClient
	companyClientGroup company.ClientGroup
	ciClientGroup      ci.ClientGroup
	logger             *log.Logger
}

type RemedyClientConfig struct {
	Logger *log.Logger
}

func NewRemedyClient(config RemedyClientConfig) (*RemedyClient, error) {
	rapidClient, err := gorapid.NewRapidClient()
	if err != nil {
		return nil, err
	}

	if config.Logger == nil {
		config.Logger = log.New(os.Stdout, "RemedyClient: ", log.LstdFlags)
	}

	client := &RemedyClient{
		rapidClient: rapidClient,
		logger:      config.Logger,
	}

	client.companyClientGroup = company.NewClientGroup(client)
	client.ciClientGroup = ci.NewClientGroup(client)

	return client, nil
}

func (rc *RemedyClient) GetRapidClient() *gorapid.RapidClient {
	return rc.rapidClient
}

func (rc *RemedyClient) GetCompanyClientGroup() company.ClientGroup {
	return rc.companyClientGroup
}

func (rc *RemedyClient) GetCIClientGroup() ci.ClientGroup {
	return rc.ciClientGroup
}
