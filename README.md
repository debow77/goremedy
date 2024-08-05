# goremedy

RemedyClient
===========

A Go client for interacting with the Remedy API through Rapid.

Installation
------------

To install the RemedyClient package, run the following command:

```bash
go get github.cerner.com/OHAIFedAutoSre/goremedy
```

Usage
-----
Authentication for Remedy endpoints in managed through the gorapid package. See [Creating a RapidClient instance](https://github.cerner.com/OHAIFedAutoSre/gorapid#creating-a-rapidclient-instance)

```bash
# Using the default loglevel
client, err := goremedy.NewRemedyClient()


# Setting loglevel to DEBUG
client, err := goremedy.NewRemedyClient(goremedy.RemedyClientConfig{LogLevel: "DEBUG",})
```

TODO - Complete usage


Making API requests
-----
TODO

Contributing
-----
Contributions are welcome! Please open a pull request or issue on GitHub to contribute to the RemedyClient package.

Acknowledgments
-----
The RemedyClient package was inspired by the https://github.cerner.com/CWxAutomation/php_remedy.

Running the examples
-----

Running a single example, you can see the example paths in examples/main.go

```bash
$ task run -- getcompany1

GetCompany - Single Company usage example:
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

Running all examples listed in examples/main.go

```bash
task run
```

