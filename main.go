package main

import (
	"log"
	"backend/db"
	"backend/db/redis"
	"backend/rest"
)

func main() {
	redis.ConnectRedis("127.0.0.1:6379", "")
	db.ConnectDB("127.0.0.1:3306", "?parseTime=true");
	log.Fatal(rest.RunAPI("0.0.0.0:8000"))
}