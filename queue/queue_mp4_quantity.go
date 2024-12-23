package queue

import (
	"app/cmd/core-service/service"
	"app/internal/connection"
	constant "app/internal/constants"
	queuepayload "app/internal/dto/queue_payload"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type queueUrlQuantity struct {
	connRabbitmq *amqp091.Connection
	service      service.Service
}
type QueueUrlQuantity interface {
	Worker()
}

func (q *queueUrlQuantity) Worker() {
	queueName := constant.QUEUE_URL_QUANTITY
	ch, err := q.connRabbitmq.Channel()

	if err != nil {
		log.Println("error chanel: ", err)
		return
	}

	qe, err := ch.QueueDeclare(
		string(queueName),
		true,
		false,
		false,
		false,
		amqp091.Table{},
	)
	if err != nil {
		log.Println("error queue declare: ", err)
		return
	}
	log.Printf("start %s", string(queueName))

	msgs, err := ch.Consume(
		qe.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("error consumer: ", err)
		return
	}

	for d := range msgs {
		go func(mess amqp091.Delivery) {
			var payload queuepayload.QueueUrlQuantityPayload
			err := json.Unmarshal(mess.Body, &payload)
			if err != nil {
				log.Println("error msg: ", err)
				mess.Reject(true)
				return
			}

			err = q.service.VideoService.UploadQuantityVideo(payload)
			if err != nil {
				log.Println("error upload url video: ", err)
				mess.Reject(true)
				return
			}

			mess.Ack(false)
		}(d)

	}
}

func NewQueueUrlQuantity() QueueUrlQuantity {
	return &queueUrlQuantity{
		connRabbitmq: connection.GetRabbitmq(),
		service:      service.Register(),
	}
}
