package configs

import (
	"log"

	"github.com/go-redis/redis"
)

// RedisClient is a global variable.
// Use this variable instead of repeating calling ConnectRedis
var RedisClient *redis.Client

// ConnectRedis called in main.go
func ConnectRedis() {
	redisOptions := &redis.Options{
		DB:       0,
		Addr:     "crud-redis:6379",
		Password: "",
	}

	log.Print("Connecting Redis ... ")
	client := redis.NewClient(redisOptions)

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Error")
		log.Fatal(err.Error())
	}

	log.Println("Success")
	RedisClient = client
}
