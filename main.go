package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/firstimedeveloper/value-investing/bs"
	"github.com/firstimedeveloper/value-investing/cp"
	"io/ioutil"
	"net/http"
)

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
