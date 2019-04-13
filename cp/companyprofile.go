package cp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Company struct {
	Profile Profile `json:"CSCO"`
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

func (c *Company) getData(companyName string) error {
	url := fmt.Sprintf("https://financialmodelingprep.com/api/company/profile/%s", companyName)
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

	err = json.Unmarshal(body, &c)
	if err != nil {
		return errors.Wrap(err, "could not unmarshall json")
	}

	return nil
}

func (c Company) CompanyProfile(companyName string) *Company {
	c.getData(companyName)

	return &c
}

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello world")
// 	})

// 	port := "8080"

// 	fmt.Println(http.ListenAndServe(":"+port, nil))
