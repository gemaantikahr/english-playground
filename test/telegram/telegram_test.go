package telegram

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSendMessage(t *testing.T) {

	token := "5907329537:AAFLJB6iuL3o9z79i3IBYtO5IjmF_0zdYEU"
	chatID := "658555117"
	message := "Hello, this is a message sent from my Go bot."

	url := "https://api.telegram.org/bot" + token + "/sendMessage"

	var jsonStr = []byte(`{"chat_id": "` + chatID + `", "text": "` + message + `"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println("error", err.Error())
		}
	}(resp.Body)

	result, _ := io.ReadAll(resp.Body)
	println("the string result ", string(result))

}

func TestGetUpdate(t *testing.T) {

	token := "5907329537:AAFLJB6iuL3o9z79i3IBYtO5IjmF_0zdYEU"
	url := "https://api.telegram.org/bot" + token + "/getUpdates"
	resp, _ := http.Get(url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println("error ", err.Error())
		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
