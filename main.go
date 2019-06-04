package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	stand "github.com/nats-io/nats-streaming-server/server"
)

func main() {
	stand.RunServer("test-cluster")

	nc, err := nats.Connect("demo.nats.io", nats.Name("NATS Client"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected!")
	nc.Subscribe("hi", func(m *nats.Msg) {
		log.Println("[Received] ", string(m.Data))
	})
	nc.Publish("hi", []byte("Hello NATS!"))

	sc, err := stan.Connect("test-cluster", "client-123", stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}
	sc.Publish("hi", []byte("Hello STAN!"))
	select {}
}
