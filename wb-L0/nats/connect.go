package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"time"
)

func Connect(ctx context.Context, clientID string) stan.Conn {

	// Connect to NATS
	opts := []nats.Option{nats.Name("NATS Streaming")}
	nc, err := nats.Connect(os.Getenv("NATS_URL"), opts...)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to NATS-Streaming
	deadline, _ := ctx.Deadline()
	sc, err := stan.Connect(os.Getenv("CLUSTER_ID"), clientID, stan.NatsURL(nc.ConnectedUrl()), stan.ConnectWait(time.Until(deadline)))

	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.NatsURL(nc.ConnectedUrl()))
	}

	select {
	case <-ctx.Done():
		nc.Close()
		log.Fatalf("Context canceled before successful connection: %v", ctx.Err())
	default:
	}

	return sc
}
