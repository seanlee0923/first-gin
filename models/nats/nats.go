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

func makeNatsClient() *nats.Conn {
	natsName := os.Getenv("NATS_HOSTNAME")
	natsPort := os.Getenv("NATS_PORT")

	natsUrl := "nats://" + natsName + ":" + natsPort
	natsConn, err := nats.Connect(natsUrl)

	if err != nil {
		panic(err)
	}

	return natsConn
}
