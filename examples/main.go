package main

import (
	"flag"
	"fmt"
	"goremedy/examples/ci"
	"goremedy/examples/company"
	"goremedy/examples/crq"
	"goremedy/examples/inc"
)

func main() {
	example := flag.String("example", "", "which example to run")
	flag.Parse()

	if *example != "" {
		switch *example {
		case "getcompany1":
			company.GetCompany()
		case "getcompany2":
			company.GetCompany2()
		case "getdomains":
			ci.GetDomains()
		case "getcomputersystembyfqdn":
			ci.GetComputerSystemByFqdn()
		case "getcomputersystembyid":
			ci.GetComputerSystemById()
		case "computersystemisdeployed":
			ci.ComputerSystemIsDeployed()
		case "getcomputersystems":
			ci.GetComputerSystems()
		case "getcrq":
			crq.GetCrq()
		case "getinc":
			inc.GetInc()
		default:
			fmt.Println("unknown example")
		}
	} else {
		fmt.Println("Running all examples:")
		company.GetCompany()
		company.GetCompany2()
		ci.GetDomains()
		ci.GetComputerSystemByFqdn()
		ci.GetComputerSystemById()
		ci.ComputerSystemIsDeployed()
		ci.GetComputerSystems()
		crq.GetCrq()
		inc.GetInc()
	}
}
