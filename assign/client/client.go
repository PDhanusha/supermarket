package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type item struct {
	Item	string
	Value	float64
}

Var Items []item

func main() {
	client := &http.Client{}
	client.Timeout = time.Second * 15
	temp := Items{item{Item:Apple,Value:20.4 },
                    item{Item:Orange,Value:21.4}}
	j, err := json.Marshal(temp)
	req, err := http.NewRequest("GET", "http://localhost:9455/testget", bytes.NewBuffer(j))
	if err != nil {
		fmt.Println("1st err")
		panic(err)
	}
	//defer req.Body.Close()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("2nd err")
		panic(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("3rd err")
		panic(err)
	}
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Body : ", string(respBody))
}

