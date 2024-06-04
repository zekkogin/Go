package pub

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/nats-io/stan.go"
	"log"
	"wb-L0/model"
)

func PublishOrder(subject string, sc stan.Conn) {
	var fakeModel model.Order
	gofakeit.Struct(&fakeModel)
	msg, _ := json.Marshal(fakeModel)
	err := sc.Publish(subject, msg)
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	log.Printf("Published [%s] : '%s'\n", subject, msg)
}
