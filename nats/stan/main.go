package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	opts := []nats.Option{
		nats.ClientCert("cert.pem", "key.pem"),
		nats.MaxReconnects(10),
	}
	nc, err := nats.Connect("tls://localhost:4443", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := stan.Connect(
		"clusterID",
		"clientID",
		stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalln(err)
	}
	sub, err := conn.Subscribe("sub", func(m *stan.Msg) {
		log.Printf("received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatalln(err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	sub.Unsubscribe()
	conn.Close()
}
