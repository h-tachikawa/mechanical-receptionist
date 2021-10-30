package repository

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"

	"cloud.google.com/go/firestore"
)

var collectionName = "visitHistory"

type FirestoreVisitHistoryRepository struct {
	client *firestore.Client
}

func NewFirestoreVisitHistoryRepository(ctx context.Context) VisitHistoryRepository {
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

	return &FirestoreVisitHistoryRepository{client: client}
}

func (f FirestoreVisitHistoryRepository) GetLatestOne(ctx context.Context) (*VisitHistory, error) {
	iter := f.client.Collection(collectionName).OrderBy("visitedAt", firestore.Desc).Limit(1).Documents(ctx)
	latestDocSnapShot, err := iter.Next()

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	latestVisitHistory := VisitHistory{}
	if err := latestDocSnapShot.DataTo(&latestVisitHistory); err != nil {
		log.Fatalln(err)
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
