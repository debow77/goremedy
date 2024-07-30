`	// GetComputerSystemById(id string) (*ConfigurationItem, error)
	// ComputerSystemIsDeployed(fqdn string) (bool, error)
	// GetBusinessServices(company, name string) ([]*ConfigurationItem, error)
	// GetBusinessServiceByName(company, name string) (*ConfigurationItem, error)
	// DomainHasUsage(company, domain, usage string) (bool, error)
	// GetComputerSystemCompany(fqdn string) (string, error)
	// GetComputerSystemMnemonic(fqdn string) (string, error)
	// GetComputerSystemDomains(fqdn string) ([]*Relationship, error)
	// GetComputerSystemGroups(fqdn string) ([]*Relationship, error)
	// GetDomainSite(company, domain string) (string, error)
	// GetComputerSystems(company string, queryFilters map[string]string) ([]*ConfigurationItem, error)
	// RelateToCr(changeId, instanceId string) error


	## Running the examples

	```bash
		# From the root of the goremedy folder
		$ go run examples/main.go -example=getcompany1
		Single Company usage example:
		2024/07/29 09:21:56 WARN Log level warn
		Status code: 200
		Found 1 companies
		Company found:
			CompanyId: CPY000000139896
			RemedyCompanyId: 0510
			Name: CernerWorks Technology Improvement
			Region: Internal
			Company: CERN_CWIM-CernerWorks Technology Improvement
			Mnemonic: CERN_CWIM
			FocusClient: 
			CompanyType: Customer
			ProdDataCenter: CTC-LS-III
			DrDataCenter: CTC-KC-I
			Status: 1
			Created Date: 0001-01-01 00:00:00 +0000 UTC
			Modified Date: 0001-01-01 00:00:00 +0000 UTC
	```