package crawler

import (
	"context"
	"errors"

	"github.com/loggerhead/doger/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	Database string
	URI      string
	client   *mongo.Client
	db       *mongo.Database
}

var cli client

func InitDB(dbname, url string) (err error) {
	cli.Database = dbname
	cli.URI = url

	if cli.client, err = mongo.NewClient(options.Client().ApplyURI(url)); err != nil {
		return
	}
	if err = cli.client.Connect(context.Background()); err != nil {
		return
	}

	cli.db = cli.client.Database(cli.Database)
	return
}

func Save(raw interface{}) (err error) {
	if raw == nil {
		err = errors.New("nil value")
		return
	}

	switch v := raw.(type) {
	case *api.UserInfo:
		table := cli.db.Collection("user_info")
		_, err = table.InsertOne(context.Background(), v)
	case *api.DynamicDetail:
		table := cli.db.Collection("dynamic_detail")
		_, err = table.InsertOne(context.Background(), v)
	case *api.DynamicComment:
		table := cli.db.Collection("dynamic_comment")
		_, err = table.InsertOne(context.Background(), v)
	default:
		err = errors.New("unkown type")
	}

	if IsDup(err) {
		err = nil
	}
	return
}

func IsDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}
