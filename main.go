package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Skaifai/RedisProject/cache"
	"github.com/Skaifai/RedisProject/util"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
	"time"
)

const version = "1.0"

var ctx = context.Background()

// Declaring the connection variable which will store the current connection to the database
var currentConnection redis.Client

func main() {
	// Declaring an instance of a reader class, which will handle the input from the console
	reader := bufio.NewReader(os.Stdin)

	// Declaring an input and err variable in advance, because we use input value as the stop condition for the loop
	var input string
	var err error

	for input != "exit" {
		fmt.Print("Enter command: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			break
		}
		input = util.CleanString(input)
		switch input {
		case "connect":
			fmt.Println("Connecting...")
			currentConnection = Connect()
		case "disconnect":
			fmt.Println("Disconnecting...")
			currentConnection.Close()
		case "exit":
			break
		default:
			fmt.Println("Incorrect input.")
		}

	}
}

func Connect() redis.Client {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the host address: ")
	host, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	host = util.CleanString(host)

	fmt.Print("Enter the db: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = util.CleanString(input)

	db, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter the expiry duration: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = util.CleanString(input)

	exp, err := time.ParseDuration(input)
	if err != nil {
		log.Fatal(err)
	}

	client := cache.NewClient(host, db, exp)
	return client
}
