package redis

import (
	"fmt"

	config "github.com/511636330/go-conf"
	"github.com/go-redis/redis/v8"
)

var clients = make(map[string]*redis.Client)

func GetClient(conn string) *redis.Client {
	if client, ok := clients[conn]; ok && client != nil {
		return client
	}

	return connect(conn)

}

func connect(conn string) *redis.Client {
	host := config.GetString(fmt.Sprintf("database.redis.%s.host", conn))
	if host == "" {
		conn = "default"
	}
	host = config.GetString(fmt.Sprintf("database.redis.%s.host", conn), "127.0.0.1")
	port := config.GetString(fmt.Sprintf("database.redis.%s.port", conn), "6379")
	password := config.GetString(fmt.Sprintf("database.redis.%s.password", conn), "")
	db := config.GetInt(fmt.Sprintf("database.redis.%s.db", conn), "0")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	clients[conn] = client

	return client
}
