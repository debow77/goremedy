package goremedy

import (
	"encoding/json"
	"goremedy/ci"
	"goremedy/company"
	"goremedy/crq"
	"goremedy/interfaces"
	"log/slog"
	"net/url"
	"os"
)

// RemedyClient represents a Remedy client
type RemedyClient struct {
	rapidClient        interfaces.RapidClientInterface
	companyClientGroup company.ClientGroup
	ciClientGroup      ci.ClientGroup
	crqClientGroup     crq.ClientGroup
	config             *RemedyClientConfig
}

// RemedyClientConfig defines the configuration for a Remedy client
type RemedyClientConfig struct {
	LogLevel string
}

// logLevels maps string log levels to slog.Level values
var logLevels = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

// NewRemedyClient creates a new Remedy client instance
func NewRemedyClient(config ...RemedyClientConfig) (*RemedyClient, error) {
	var cfg RemedyClientConfig
	if len(config) > 0 {
		cfg = config[0]
	} else {
		cfg = RemedyClientConfig{
			LogLevel: "INFO", // default log level
		}
	}

	setLogLevel(cfg)

	rapidClientInterface, err := interfaces.NewRapidClient()
	if err != nil {
		return nil, err
	}

	client := &RemedyClient{
		rapidClient: rapidClientInterface,
		config:      &cfg,
	}

	client.companyClientGroup, err = company.NewClientGroup(client)
	if err != nil {
		return nil, err
	}

	client.ciClientGroup, err = ci.NewClientGroup(client)
	if err != nil {
		return nil, err
	}

	client.crqClientGroup, err = crq.NewClientGroup(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// setLogLevel sets the log level based on the config
func setLogLevel(config RemedyClientConfig) {
	level, ok := logLevels[config.LogLevel]
	if !ok {
		level = slog.LevelInfo // Default to INFO if invalid level provided
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

	if config.LogLevel != "INFO" {
		if config.LogLevel == "WARN" {
			slog.Warn("Log level set", "level", config.LogLevel)
		} else if config.LogLevel == "ERROR" {
			slog.Error("Log level set", "level", config.LogLevel)
		} else if config.LogLevel == "DEBUG" {
			slog.Debug("Log level set", "level", config.LogLevel)
		}
	}
}

// GetRapidClient returns the Rapid client instance
func (rc *RemedyClient) GetRapidClient() interfaces.RapidClient {
	return rc.rapidClient.GetRapidClient()
}

// BaseURL returns the interfaces base url instances for Rapid
func (rc *RemedyClient) BaseURL() string {
	return rc.rapidClient.BaseURL()
}

// GetPaginated returns the interfaces GetPaginated instances
func (rc *RemedyClient) GetPaginated(basePath, urlPath string, params url.Values) ([]json.RawMessage, int, error) {
	return rc.rapidClient.GetPaginated(basePath, urlPath, params)
}

// GetCompanyClientGroup returns the company client group instance
func (rc *RemedyClient) GetCompanyClientGroup() company.ClientGroup {
	return rc.companyClientGroup
}

// GetCIClientGroup returns the CI client group instance
func (rc *RemedyClient) GetCIClientGroup() ci.ClientGroup {
	return rc.ciClientGroup
}

// GetCRQClientGroup returns the CRQ client group instance
func (rc *RemedyClient) GetCRQClientGroup() crq.ClientGroup {
	return rc.crqClientGroup
}
