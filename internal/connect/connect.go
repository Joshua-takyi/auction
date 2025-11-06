package connect

import (
	"context"
	"fmt"
	"github.com/supabase-community/supabase-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

var (
	SupabaseClient *supabase.Client
	MongoDbClient  *mongo.Client
)

func MongoDbConnect(url, password string) (*mongo.Client, error) {
	fullUrl := strings.Replace(url, "<password>", password, 1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	options := options.Client().ApplyURI(fullUrl)

	var err error
	MongoDbClient, err = mongo.Connect(ctx, options)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo %v", err)
	}

	if err := MongoDbClient.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb %v", err)
	}

	return MongoDbClient, nil
}

func DisconnectMongodb() error {
	if MongoDbClient == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := MongoDbClient.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("discounnection failed %v", err)
	}
	MongoDbClient = nil
	return nil
}

func ConnectSupabase(url, key string) (*supabase.Client, error) {
	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to supabase server")
	}
	SupabaseClient = client
	return SupabaseClient, nil
}

func DisconnectSupabase() error {
	SupabaseClient = nil
	return nil
}
