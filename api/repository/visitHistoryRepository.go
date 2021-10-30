package repository

import (
	"context"
	"time"
)

type VisitHistory struct {
	VisitedAt time.Time `firestore:"visitedAt"` // 構造体のフィールド名はアッパーキャメルで書かないと構造体に上手くマッピングしてくれないので注意
}

type VisitHistoryRepository interface {
	GetLatestOne(ctx context.Context) (*VisitHistory, error)
	Add(ctx context.Context, doc *VisitHistory) error
}
