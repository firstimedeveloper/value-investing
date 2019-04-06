package balanceSheet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BalanceSheet struct {
	Symbol     string `json:"symbol"`
	Financials []struct {
		Date                        string `json:"date"`
		CashAndCashEquivalents      string `json:"Cash and cash equivalents"`
		ShortTermInvestments        string `json:"Short-term investments"`
		CashAndShortTermInvestments string `json:"Cash and short-term investments"`
		Receivables                 string `json:"Receivables"`
		Inventories                 string `json:"Inventories"`
		TotalCurrentAssets          string `json:"Total current assets"`
		PropertyPlantEquipmentNet   string `json:"Property Plant & Equipment Net"`
		GoodwillAndIntangibleAssets string `json:"Goodwill and Intangible Assets"`
		LongTermInvestments         string `json:"Long-term investments"`
		TaxAssets                   string `json:"Tax assets"`
		TotalNonCurrentAssets       string `json:"Total non-current assets"`
		TotalAssets                 string `json:"Total assets"`
		Payables                    string `json:"Payables"`
		ShortTermDebt               string `json:"Short-term debt"`
		TotalCurrentLiabilities     string `json:"Total current liabilities"`
		LongTermDebt                string `json:"Long-term debt"`
		TotalDebt                   string `json:"Total debt"`
		DeferredRevenue             string `json:"Deferred revenue"`
		TaxLiabilities              string `json:"Tax Liabilities"`
		DepositLiabilities          string `json:"Deposit Liabilities"`
		TotalNonCurrentLiabilities  string `json:"Total non-current liabilities"`
		TotalLiabilities            string `json:"Total liabilities"`
		OtherComprehensiveIncome    string `json:"Other comprehensive income"`
		RetainedEarningsDeficit     string `json:"Retained earnings (deficit)"`
		ShareholdersEquity          string `json:"Shareholders Equity"`
		Investments                 string `json:"Investments"`
		NetDebt                     string `json:"Net Debt"`
	} `json:"financials"`
}

func (b BalanceSheet) balanceSheet() {

	value := BalanceSheet{}
	url := "https://financialmodelingprep.com/api/financials/balance-sheet-statement/AAPL?period=quarter"
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

	//fmt.Println(value)

	for _, i := range value.Financials {
		fmt.Printf("%s %s\n", i.Date, i.TotalAssets)
	}
}
