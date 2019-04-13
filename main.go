package main

import (
	"fmt"
	"github.com/firstimedeveloper/value-investing/bs"
	"github.com/firstimedeveloper/value-investing/cp"
)

func main() {

	value1 := cp.Company{}
	value2 := bs.BalanceSheet{}
	companyName := "CSCO"

	fmt.Println(*value1.CompanyProfile(companyName))
	fmt.Println(*value2.CompanyBS(companyName))

}
