package main

import (
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

	// Create a queue subscription on "updates" with queue name "workers"
	sub, err := nc.QueueSubscribe(messaging.IN_Q, messaging.GROUP, func(m *nats.Msg) {

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
