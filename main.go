package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	nc, err := nats.Connect("127.0.0.1", nats.Name("NATS Client"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected!")
	nc.Subscribe("hi", func(m *nats.Msg) {
		log.Println("[Received]", string(m.Data))
	})
	nc.Publish("hi", []byte("Hello NATS!"))

	sc, err := stan.Connect("test-cluster", "client-123", stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}
	sc.Publish("hi", []byte("Hello STAN!"))
	sc.Subscribe("hi", func (m *stan.Msg){
		log.Println("[Received]", string(m.Data))
	}, stan.DeliverAllAvailable())
	select {}
}
