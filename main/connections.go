package main

import (
	"fmt"
	"github.com/Skaifai/RedisProject/cache"
	"github.com/Skaifai/RedisProject/util"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

func Connect() redis.Client {
	fmt.Print("Enter the host address: ")
	host := util.ReadAndCleanString()

	fmt.Print("Enter the db: ")
	input := util.ReadAndCleanString()

	db, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter the expiry duration: ")
	input = util.ReadAndCleanString()

	exp, err := time.ParseDuration(input)
	if err != nil {
		log.Fatal(err)
	}

	client := cache.NewClient(host, db, exp)
	output, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(output)
		fmt.Println("Connection successful.")
	}
	isConnectedToDB = true
	return client
}
