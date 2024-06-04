package sub

import "C"
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"wb-L0/model"
	"wb-L0/pkg/app"
)

func SubcribeOrders(app app.App) {
	subject := os.Getenv("SUBJECT")
	_, err := app.Sc.Subscribe(subject, func(msg *stan.Msg) {
		var receivedMsg model.Order
		err := json.Unmarshal(msg.Data, &receivedMsg)
		if err != nil {
			log.Printf("Не удалось размаршаллировать сообщение: %v", err)
			return
		}
		//fmt.Printf("Получено сообщение: %+v\n", receivedMsg)
		app.C.CacheAdd(receivedMsg)
		err = app.Pool.InsertOrder(context.Background(), receivedMsg)
		if err != nil {
			fmt.Printf("error %s", err)
		}
		msg.Ack()
	}, stan.SetManualAckMode(), stan.DurableName(subject))

	if err != nil {
		log.Fatalf("Не удалось подписаться на сообщения: %v", err)
	}

}
