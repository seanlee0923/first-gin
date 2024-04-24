package nats

import (
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

type Post struct {
	Id      uint
	Title   string `json:"title"`
	Content string `json:"content"`
	Regdate time.Time
}

func NatsConnection() (*nats.Conn, *nats.EncodedConn) {
	natsName := os.Getenv("NATS_HOSTNAME")
	natsPort := os.Getenv("NATS_PORT")

	natsUrl := "nats://" + natsName + ":" + natsPort
	natsConn, err := nats.Connect(natsUrl)
	if err != nil {
		panic(err)
	}
	encodedConn, _ := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)
	return natsConn, encodedConn
}

func publishPost(topic string, post Post) {
	_, encodedConn := NatsConnection()

	encodedConn.Publish(topic, "")
}
