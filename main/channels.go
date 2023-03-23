package main

import (
	"fmt"
	"github.com/Skaifai/RedisProject/util"
	"github.com/redis/go-redis/v9"
	"log"
)

var pubsub *redis.PubSub

func SubscribeAndListen() {
	fmt.Print("Enter the channels: ")
	channels := util.ReadAndCleanString()
	pubsub = currentConnection.Subscribe(ctx, channels)
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(pubsub)

	result, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Should be *Subscription, but others are possible if other actions have been
	// taken on pubsub since it was created.
	switch result.(type) {
	case *redis.Subscription:
		// subscribe succeeded
		fmt.Println("Subscribed current connection to channels: " + pubsub.String())
	case *redis.Message:
		// Assert type
		m, ok := result.(redis.Message)
		if ok {
			// received first message
			fmt.Println("Received first message.")
			fmt.Println("Channel: " + m.Channel)
			fmt.Println("Message: " + m.Payload)
		}
	case *redis.Pong:
		// pong received
		fmt.Println("Pong received.")
	default:
		log.Fatal("Error in redis has occurred.")
	}

	// Make a channel that will track if the exit command was issued
	quit := make(chan string)
	go func() {
		for {
			input := util.ReadAndCleanString()
			if input == "exit" {
				quit <- input
				close(quit)
				return
			}
		}
	}()

	// Make a channel that will receive messages
	ch := pubsub.Channel()
	for {
		select {
		case message := <-ch:
			fmt.Println("Channel: " + message.Channel + " \nMessage: " + message.Payload)
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}

	//for msg := range ch {
	//	fmt.Println("Channel: " + msg.Channel + " \nMessage: " + msg.Payload)
	//}

}
