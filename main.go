package atmsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllAccounts() {
	res, err := http.Get("http://localhost:8000/api/account/get/all")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	accountList := string(body)
	fmt.Println(accountList)
}

func GetAccountById(id int) {
	res, err := http.Get(fmt.Sprint("http://localhost:8000/api/account/get/", id))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	account := string(body)
	fmt.Println(account)
}

func Withdrawal(id int, value float64) {
	client := http.Client{}

	url := fmt.Sprint("http://localhost:8000/api/account/withdraw/", id)

	jsonData, err := json.Marshal(map[string]float64{"value": value})
	if err != nil {
		panic(err)
	}

	bodyRequest := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest(http.MethodPut, url, bodyRequest)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.StatusCode)
}

func Deposit(id int, value float64) {
	client := http.Client{}

	url := fmt.Sprint("http://localhost:8000/api/account/deposit/", id)

	jsonData, err := json.Marshal(map[string]float64{"value": value})
	if err != nil {
		panic(err)
	}

	bodyRequest := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest(http.MethodPut, url, bodyRequest)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.StatusCode)
}
