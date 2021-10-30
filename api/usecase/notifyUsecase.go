package usecase

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type VisitHistory struct {
	VisitedAt time.Time `firestore:"visitedAt"` // 構造体のフィールド名はアッパーキャメルで書かないと構造体に上手くマッピングしてくれないので注意
}

type ConnectionSettings struct {
	endpoint string
	token    string
}

type LineNotifier struct {
	connSettings *ConnectionSettings
}

func NewLineNotifier(connSettings *ConnectionSettings) *LineNotifier {
	return &LineNotifier{connSettings}
}

func (receiver *LineNotifier) notify(message string) error {
	body := strings.NewReader("message=" + message)
	req, _ := http.NewRequest(http.MethodPost, receiver.connSettings.endpoint, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", receiver.connSettings.token))

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

func NotifyUseCase() error {
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: os.Getenv("GCP_PROJECT_ID"),
	}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	collectionName := "visitHistory"

	iter := client.Collection(collectionName).OrderBy("visitedAt", firestore.Desc).Limit(1).Documents(ctx)
	latestDocSnapShot, err := iter.Next()

	if err != nil {
		fmt.Println(err)
		return err
	}

	latestVisitHistory := VisitHistory{}
	if err := latestDocSnapShot.DataTo(&latestVisitHistory); err != nil {
		fmt.Println(err)
		return err
	}

	latestVisitedTime := latestVisitHistory.VisitedAt
	current := time.Now()

	fmt.Println("current", current)
	fmt.Println("latest", latestVisitedTime)

	durationAsSec := current.Sub(latestVisitedTime).Seconds()

	fmt.Println(durationAsSec)
	if durationAsSec < 60 {
		fmt.Println("前回の実行から1分以内なので何もしません")
		return nil
	}

	doc := make(map[string]interface{})
	doc["visitedAt"] = firestore.ServerTimestamp

	_, _, err = client.Collection(collectionName).Add(ctx, doc)

	if err != nil {
		log.Printf("an error has occurred: %s", err)
	}

	connSettings := &ConnectionSettings{
		endpoint: os.Getenv("LINE_NOTIFY_ENDPOINT"),
		token:    os.Getenv("LINE_NOTIFY_TOKEN"),
	}

	lineNotifier := NewLineNotifier(connSettings)
	err = lineNotifier.notify("来客です。対応してください。")

	if err != nil {
		return err
	}

	return nil
}
