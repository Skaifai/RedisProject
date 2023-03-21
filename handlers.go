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
	isConnectedToDB = true
	return client
}

func SetString() {
	fmt.Print("Enter the key: ")
	key := util.ReadAndCleanString()

	fmt.Print("Enter the value: ")
	val := util.ReadAndCleanString()

	fmt.Print("Enter the expiry duration: ")
	input := util.ReadAndCleanString()

	exp, err := time.ParseDuration(input)
	if err != nil {
		log.Fatal(err)
	}

	currentConnection.Set(ctx, key, val, exp)
}

func GetString() {
	fmt.Print("Enter the key: ")
	key := util.ReadAndCleanString()
	val, err := currentConnection.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Key: " + key)
	fmt.Println("Value: " + val)
}

func UpdateString() {
	fmt.Print("Enter the key: ")
	key := util.ReadAndCleanString()

	fmt.Print("Enter the value: ")
	val := util.ReadAndCleanString()

	oldValue, _ := currentConnection.Get(ctx, key).Result()

	success := currentConnection.SetXX(ctx, key, val, redis.KeepTTL).Val()
	if !success {
		fmt.Println("Key-value with the specified key does not exist.")
		return
	}
	fmt.Println("Key: " + key)
	fmt.Println("Value: " + oldValue + " --> " + val)
}

func DeleteString() {
	fmt.Print("Enter the key: ")
	key := util.ReadAndCleanString()

	val, err := currentConnection.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	currentConnection.Del(ctx, key)
	fmt.Println("Deleted key-value pair.")
	fmt.Println("Key: " + key)
	fmt.Println("Value: " + val)
}
