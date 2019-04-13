package bs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
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

func (b *BalanceSheet) getData(companyName string) error {
	url := fmt.Sprintf("https://financialmodelingprep.com/api/financials/balance-sheet-statement/%s?period=quarter", companyName)
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "could not get url")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "could not read response body")
	}

	//need this cause the json provided by the api is shitty
	body = bytes.TrimPrefix(body, []byte{60, 112, 114, 101, 62})
	body = bytes.TrimSuffix(body, []byte{60, 112, 114, 101, 62})

	err = json.Unmarshal(body, &b)
	if err != nil {
		return errors.Wrap(err, "could not unmarshall json")
	}

	return nil
}

func (b BalanceSheet) CompanyBS(companyName string) *BalanceSheet {
	b.getData(companyName)

	return &b
}
