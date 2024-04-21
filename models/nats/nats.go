package nats

import (
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

type natsRequest struct {
	postId  uint
	message string
}

type natsResponse struct {
	Conn         *nats.Conn
	Subscription *nats.Subscription
}

type NatsClient struct {
	Conn          *nats.Conn
	Subs          map[string]*nats.Subscription
	UniqueReplyTo map[string]string
	ReqTimeout    time.Duration
}

func makeNatsClient() *NatsClient {
	natsName := os.Getenv("NATS_HOSTNAME")
	natsPort := os.Getenv("NATS_PORT")

	natsUrl := "nats://" + natsName + ":" + natsPort
	natsConn, err := nats.Connect(natsUrl)

	if err != nil {
		panic(err)
	}
	connect := &NatsClient{
		Conn:          natsConn,
		Subs:          make(map[string]*nats.Subscription),
		UniqueReplyTo: make(map[string]string),
		ReqTimeout:    time.Second * 5,
	}

	return connect
}
