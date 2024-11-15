package queue

import (
	"app/config"
	"app/constant"
	queuepayload "app/dto/queue_payload"
	"app/service"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type queueUrlQuantity struct {
	connRabbitmq *amqp091.Connection
	videoService service.VideoService
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

			err = q.videoService.UploadQuantityVideo(payload)
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
		connRabbitmq: config.GetRabbitmq(),
		videoService: service.NewVideoService(),
	}
}
