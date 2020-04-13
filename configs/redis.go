package configs

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// ConnectRedis - redis connection
func ConnectRedis() redis.Conn {
	conn, err := redis.Dial("tcp", "crud-redis:6379")
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}
