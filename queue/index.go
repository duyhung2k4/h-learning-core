package queue

func InitQueue() {
	queueUrlQuantity := NewQueueUrlQuantity()

	go func() {
		queueUrlQuantity.Worker()
	}()
}
