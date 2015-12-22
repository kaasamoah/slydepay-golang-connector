// mobile
package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MobileClient struct {
}

func CreateOrder() (token string) {
	var url = "http://stage.airty.me/api/getItemList"
	var dataBytes = []byte(`{"username": "user@slydepay.com.gh","password":"1234567890"}`)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dataBytes))

	if err != nil {
		result := "Sorry, an error occurred"
		return result
	}

	if resp.StatusCode != 200 {
		result := fmt.Sprintln("Error communicating with server: ", resp.StatusCode)
		return result
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result := string(body)

	log.Println("Server returned response: ", result)
	return result
}
