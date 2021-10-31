package domain

import "time"

type VisitHistory struct {
	VisitedAt time.Time `firestore:"visitedAt"` // 構造体のフィールド名はアッパーキャメルで書かないと構造体に上手くマッピングしてくれないので注意
}

func NewVisitHistory(visitedAt time.Time) VisitHistory {
	return VisitHistory{VisitedAt: visitedAt}
}
