package nosql

import (
	"context"
	"github.com/rmukubvu/pumpdata/bus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type LogStore struct {
	client *mongo.Client
	db     *mongo.Database
}

type MongoObject struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Data        string             `json:"data" bson:"data,omitempty"`
	CreatedDate string             `json:"created_date" bson:"created_date,omitempty"`
}

const (
	schemaName = "amakosi_logs"
)

func NewConnection(uri string) *LogStore {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := getContext()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &LogStore{db: client.Database(schemaName),
		client: client}
}

func (l *LogStore) InsertRecord(document bus.DataEvent) error {
	collection := l.db.Collection(document.Collection)
	ctx, _ := getContext()
	//then insert
	_, err := collection.InsertOne(ctx, document)
	return err
}

func (l *LogStore) closeDb() {
	if l.client != nil {
		ctx, _ := getContext()
		l.client.Disconnect(ctx)
	}
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
