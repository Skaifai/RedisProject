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

// Connect function establishes a connection to a Redis server using the provided host address, database number,
// and expiry duration. It uses the cache package to create a new Redis client object and then calls the Ping method of
// that client to test the connection. If the connection is successful, the function sets a global variable
// isConnectedToDB to true and returns the client object.
// The isConnectedToDB variable is used to check if a connection to the Redis server has already been established
// before executing other functions that require the Redis client object.
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

	// Run a timer with the connection's duration. When the connection expires, the program will warn the user.
	connectionTimer = time.NewTimer(exp)
	go func() {
		<-connectionTimer.C
		fmt.Println("\nConnection expired. Please, connect again to resume operations.")
		fmt.Print("Enter command: ")
	}()
	return client
}
