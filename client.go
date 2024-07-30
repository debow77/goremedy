package goremedy

import (
	"log/slog"

	"goremedy/ci"
	"goremedy/company"

	"github.cerner.com/OHAIFedAutoSre/gorapid"
)

// RemedyClientInterface defines the interface for a Remedy client
type RemedyClientInterface interface {
	// GetRapidClient returns the Rapid client instance
	GetRapidClient() *gorapid.RapidClient
	// GetCompanyClientGroup returns the company client group instance
	GetCompanyClientGroup() company.ClientGroup
	// GetCIClientGroup returns the CI client group instance
	GetCIClientGroup() ci.ClientGroup
}

// RemedyClient represents a Remedy client
type RemedyClient struct {
	// rapidClient is the Rapid client instance
	rapidClient *gorapid.RapidClient
	// companyClientGroup is the company client group instance
	companyClientGroup company.ClientGroup
	// ciClientGroup is the CI client group instance
	ciClientGroup ci.ClientGroup
}

// RemedyClientConfig defines the configuration for a Remedy client
type RemedyClientConfig struct {
	// LogLevel specifies the log level for the client
	LogLevel string
}

func init() {
	// Initialize the log level to INFO
	config := RemedyClientConfig{
		LogLevel: "INFO",
	}
	setLogLevel(config)
}

// setLogLevel sets the log level based on the config, can override from calling modules/packages
func setLogLevel(config RemedyClientConfig) {
	switch config.LogLevel {
	case "DEBUG":
		// slog.SetDefault(slog.LevelDebug)
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Log level debug")
	case "INFO":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "WARN":
		slog.SetLogLoggerLevel(slog.LevelWarn)
		slog.Warn("Log level warn")
	case "ERROR":
		slog.SetLogLoggerLevel(slog.LevelError)
		slog.Error("Log level error")
	default:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}
}

// NewRemedyClient creates a new Remedy client instance
func NewRemedyClient(config RemedyClientConfig) (*RemedyClient, error) {
	// Set the log level based on the config
	setLogLevel(config)
	// Create a new Rapid client instance
	rapidClient, err := gorapid.NewRapidClient()
	if err != nil {
		return nil, err
	}

	// Create a new Remedy client instance
	client := &RemedyClient{
		rapidClient: rapidClient,
	}

	// Create a new company client group instance
	client.companyClientGroup, err = company.NewClientGroup(client)
	if err != nil {
		return nil, err
	}

	// Create a new CI client group instance
	client.ciClientGroup, err = ci.NewClientGroup(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// GetRapidClient returns the Rapid client instance
func (rc *RemedyClient) GetRapidClient() *gorapid.RapidClient {
	return rc.rapidClient
}

// GetCompanyClientGroup returns the company client group instance
func (rc *RemedyClient) GetCompanyClientGroup() company.ClientGroup {
	return rc.companyClientGroup
}

// GetCIClientGroup returns the CI client group instance
func (rc *RemedyClient) GetCIClientGroup() ci.ClientGroup {
	return rc.ciClientGroup
}
