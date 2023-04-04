package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/sunangel-project/go-horizon-service/src/messaging"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	defer nc.Close()

	// Use a WaitGroup to wait for 10 messages to arrive
	wg := sync.WaitGroup{}
	wg.Add(10)

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	sub, err := ec.Subscribe(messaging.IN_Q, func(spot_msg *messaging.SpotMessage) {
		log.Printf("%d", spot_msg.Part.Id)
		wg.Done()
	})

	if err != nil {
		log.Fatal(err)
	}

	// Wait for messages to come in
	wg.Wait()

	sub.Unsubscribe()

	// Drain connection (Preferred for responders)
	nc.Drain()
}

func handle_message(message *nats.Msg) {
	var spot_message *messaging.SpotMessage

	println("Received message")

	payload := message.Data
	err := json.Unmarshal(payload, spot_message)
	if err != nil {
		println(err)
		panic(err)
	}

	println(spot_message)
}
