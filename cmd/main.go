package main

import (
	"fmt"
	"sync"

	"github.com/csthompson/nats-scheduler/internal/config"
	nats "github.com/nats-io/nats.go"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	// Connect to a server
	nc, _ := nats.Connect(config.NatsUrl)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		wg.Done()
	})

	wg.Wait()
}
