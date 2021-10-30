package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/h-tachikawa/mechanical-receptionist/api/domain"

	"github.com/h-tachikawa/mechanical-receptionist/api/repository"

	"github.com/h-tachikawa/mechanical-receptionist/api/adapter"
)

type VisitHistory struct {
	VisitedAt time.Time `firestore:"visitedAt"` // 構造体のフィールド名はアッパーキャメルで書かないと構造体に上手くマッピングしてくれないので注意
}

func NotifyUseCase() error {
	ctx := context.Background()
	repo := repository.NewFirestoreVisitHistoryRepository(ctx)
	latestVisitHistory, err := repo.GetLatestOne(ctx)
	latestVisitedTime := latestVisitHistory.VisitedAt

	if err != nil {
		return err
	}

	currentTime := time.Now()
	notificationTargetSpec := domain.NewNotificationTargetSpecification(latestVisitedTime, currentTime)

	if !notificationTargetSpec.IsSatisfied() {
		log.Println("前回の実行から1分以内なので何もしません")
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
	err = lineNotifier.Notify(fmt.Sprintf("来客です。対応してください。\n訪問時刻 => %s", time.Now().Format("2006-01-02 15:04:05")))

	if err != nil {
		return err
	}

	return nil
}
