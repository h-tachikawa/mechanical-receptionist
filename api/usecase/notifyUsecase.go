package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/h-tachikawa/mechanical-receptionist/api/repository"

	"github.com/h-tachikawa/mechanical-receptionist/api/adapter"

	firebase "firebase.google.com/go"
)

type VisitHistory struct {
	VisitedAt time.Time `firestore:"visitedAt"` // 構造体のフィールド名はアッパーキャメルで書かないと構造体に上手くマッピングしてくれないので注意
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

	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln("an error occurred", err)
		}
	}(client)

	repo := repository.NewFirestoreVisitHistoryRepository(client)

	latestVisitHistory, err := repo.GetLatestOne(ctx)

	if err != nil {
		return err
	}

	current := time.Now()

	fmt.Println("current", current)
	fmt.Println("latest", latestVisitHistory.VisitedAt)

	durationAsSec := current.Sub(latestVisitHistory.VisitedAt).Seconds()

	fmt.Println(durationAsSec)
	if durationAsSec < 60 {
		fmt.Println("前回の実行から1分以内なので何もしません")
		return nil
	}

	err = repo.Add(ctx, latestVisitHistory)
	if err != nil {
		return err
	}

	connSettings := &adapter.ConnectionSettings{
		Endpoint: os.Getenv("LINE_NOTIFY_ENDPOINT"),
		Token:    os.Getenv("LINE_NOTIFY_TOKEN"),
	}

	lineNotifier := adapter.NewLineNotifier(connSettings)
	err = lineNotifier.Notify("来客です。対応してください。\n" +
		"訪問時刻 => " + time.Now().Format("2006-01-02 15:04:05"))

	if err != nil {
		return err
	}

	return nil
}
