package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/h-tachikawa/mechanical-receptionist/adapter"
	"github.com/h-tachikawa/mechanical-receptionist/domain"
	"github.com/h-tachikawa/mechanical-receptionist/repository"
)

type NotificationUseCase struct{}

func NewNotificationUseCase() NotificationUseCase {
	return NotificationUseCase{}
}

func (n NotificationUseCase) Execute() error {
	ctx := context.Background()
	visitHistoryRepo := repository.NewFirestoreVisitHistoryRepository(ctx)
	latestVisitHistory, err := visitHistoryRepo.GetLatest(ctx)
	latestVisitedTime := latestVisitHistory.VisitedAt

	if err != nil {
		return err
	}

	currentTime := time.Now()
	notificationTargetSpecification := domain.NewNotificationTargetSpecification(latestVisitedTime, currentTime)

	if !notificationTargetSpecification.IsSatisfied() {
		log.Println("前回の実行から1分以内なので何もしません")
		return nil
	}

	current := domain.NewVisitHistory(currentTime)

	err = visitHistoryRepo.Add(ctx, &current)
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
