package repository

import (
	"context"
	"github.com/aydinnyunus/wallet-tracker/domain/repository"
	"github.com/go-redis/redis/v8"
)


// global constants for file
const ()

// global variables (not cool) for this file
var ()

// ConnectToRedis establishes a live connection to your agents database.
// Make sure to have permissions and network configurations so that use can connect to database. Usually,
// database ports and hosts are not public in enterprise networks. So, that part is on you to check.
func ConnectToRedis(dbConfig repository.Database) (*redis.Client, context.Context, error) {

	var ctx = context.Background()
	var rdb = redis.NewClient(&redis.Options{
		Addr:     dbConfig.DBAddr + dbConfig.DBPort,
		Password: dbConfig.DBPass, // no password set
		DB:       0,               // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, nil, err
	}

	// handle possible errors.

	return rdb, ctx, err
}

func AddRedis(rdb *redis.Client, ctx context.Context, key, value string) error {

	err := rdb.RPush(ctx, key, value).Err()
	if err != nil {
		panic(err)
	}

	return err
}

func ReadRedis(rdb *redis.Client, ctx context.Context, key string, limit int) []string {

	val, err := rdb.LRange(ctx, key, 0, int64(limit-1)).Result()
	if err != nil {
		panic(err)
	}
	return val

}
