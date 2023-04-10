package main

import (
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/sunangel-project/go-horizon-service/src/messaging"
	"github.com/sunangel-project/horizon"
	"github.com/sunangel-project/horizon/location"
)

func main() {
	nc := messaging.Connect()
	defer nc.Close()
	ec := messaging.EncodedConnection(nc)
	defer ec.Close()
	js := messaging.JetStream(nc)
	kv := messaging.KeyValueHorizon(js)

	// Use a WaitGroup to wait for 10 messages to arrive
	wg := sync.WaitGroup{}
	wg.Add(10)

	sub, err := ec.Subscribe(messaging.IN_Q, func(spot_msg *messaging.SpotMessage) {
		log.Printf("%d", spot_msg.Part.Id)
		handle_message(spot_msg, kv)
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

func handle_message(spot_msg *messaging.SpotMessage, kv nats.KeyValue) {

	loc := location.Location{
		Latitude:  spot_msg.Spot.Loc.Lat,
		Longitude: spot_msg.Spot.Loc.Lon,
	}
	radius := 500

	_ = horizon.NewHorizon(&loc, radius)

	log.Println("Calculated horizon")
	//log.Println(spot_horizon)
}
