package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const GETBLOCKCOUNT  = "getblockcount"
const GETBLOCKNEWADDRESS  = "getnewaddress"
const GETDIFFICULTY  = "getdifficulty"
func main() {
	s := GETBLOCKCOUNT
	body := bitcoincode(s)
	fmt.Println(body)
}


//通过变量与bitcoind进行交互
func bitcoincode(s string) string {
	str := strings.Fields(s)
	url := "http://user:pwd@127.0.0.1:8332"
	curl1 := `{"jsonrpc":"2.0","id":"curltest","method":"`
	curl2 := `","params":[`
	curl3 := `]}`

	var quest string
	switch len(str) {
	case 1:
		quest = fmt.Sprintln(curl1+str[0]+curl2+curl3)
	case 2:
		quest = fmt.Sprintln(curl1+str[0]+curl2+"\""+str[1]+"\""+curl3)
	}

	fmt.Println(quest)
	var jsonStr = []byte(quest)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Encoding","UTF-8")
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Basic "+"user:pwd")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body)
}
