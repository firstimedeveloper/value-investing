package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/firstimedeveloper/value-investing"
	"io/ioutil"
	"net/http"
)

type Company struct {
	Profile      Profile `json:"CSCO"`
	BalanceSheet balanceSheet.BalanceSheet
}
type Profile struct {
	Price       float64     `json:"Price"`
	Beta        string      `json:"Beta"`
	VolAvg      string      `json:"VolAvg"`
	MktCap      string      `json:"MktCap"`
	LastDiv     string      `json:"LastDiv"`
	Range       string      `json:"Range"`
	Changes     float64     `json:"Changes"`
	ChangesPerc string      `json:"ChangesPerc"`
	CompanyName string      `json:"companyName"`
	Exchange    string      `json:"exchange"`
	Industry    string      `json:"industry"`
	Website     string      `json:"website"`
	Description string      `json:"description"`
	CEO         string      `json:"CEO"`
	Sector      string      `json:"sector"`
	DateIsFiled interface{} `json:"date_is_filed"`
	DateBsFiled interface{} `json:"date_bs_filed"`
	DateCsFiled interface{} `json:"date_cs_filed"`
	Image       string      `json:"image"`
}

func main() {

	value := Company{}
	url := "https://financialmodelingprep.com/api/company/profile/CSCO"
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	//need this cause the json provided by the api is shitty
	body = bytes.TrimPrefix(body, []byte{60, 112, 114, 101, 62})
	body = bytes.TrimSuffix(body, []byte{60, 112, 114, 101, 62})

	err = json.Unmarshal(body, &value)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
}
