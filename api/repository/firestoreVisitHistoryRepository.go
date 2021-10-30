package repository

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

var collectionName = "visitHistory"

type FirestoreVisitHistoryRepository struct {
	client *firestore.Client
}

func NewFirestoreVisitHistoryRepository(client *firestore.Client) VisitHistoryRepository {
	return &FirestoreVisitHistoryRepository{client: client}
}

func (f FirestoreVisitHistoryRepository) GetLatestOne(ctx context.Context) (*VisitHistory, error) {
	iter := f.client.Collection(collectionName).OrderBy("visitedAt", firestore.Desc).Limit(1).Documents(ctx)
	latestDocSnapShot, err := iter.Next()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	latestVisitHistory := VisitHistory{}
	if err := latestDocSnapShot.DataTo(&latestVisitHistory); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &latestVisitHistory, nil
}

func (f FirestoreVisitHistoryRepository) Add(ctx context.Context, doc *VisitHistory) error {
	_, _, err := f.client.Collection(collectionName).Add(ctx, doc)

	if err != nil {
		log.Printf("an error has occurred: %s", err)
		return err
	}

	return nil
}
