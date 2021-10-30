package adapter

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ConnectionSettings struct {
	Endpoint string
	Token    string
}

type LineNotifier struct {
	connSettings *ConnectionSettings
}

func NewLineNotifier(connSettings *ConnectionSettings) NotifierAdapter {
	return &LineNotifier{connSettings}
}

func (receiver *LineNotifier) Notify(message string) error {
	body := strings.NewReader("message=" + message)
	req, _ := http.NewRequest(http.MethodPost, receiver.connSettings.Endpoint, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", receiver.connSettings.Token))

	client := new(http.Client)
	resp, err := client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(respBody))

	if err != nil {
		fmt.Println("error", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		err := errors.New("http request failed")
		return err
	}

	return nil
}
