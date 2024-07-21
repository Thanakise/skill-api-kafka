package skill

import (
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

type Consumer struct {
	SaramaConsumer          sarama.Consumer
	SaramaPartitionConsumer sarama.PartitionConsumer
	Handler                 Handler
}

func NewConsumer(handler *Handler, brokers []string, topics string) *Consumer {
	consumer, err := sarama.NewConsumer(brokers, sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	partitionConsumer, err := consumer.ConsumePartition(topics, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	return &Consumer{
		SaramaConsumer:          consumer,
		SaramaPartitionConsumer: partitionConsumer,
		Handler:                 *handler,
	}
}

func (c Consumer) CloseConection() {
	if err := c.SaramaConsumer.Close(); err != nil {
		log.Fatalln(err)
	}
	if err := c.SaramaPartitionConsumer.Close(); err != nil {
		log.Fatalln(err)
	}
}

func (c Consumer) Listen() {

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
ConsumerLoop:
	for {
		select {
		case msg := <-c.SaramaPartitionConsumer.Messages():
			c.Handler.ActiveHandler(msg)
		case <-signals:
			break ConsumerLoop
		}
	}
}
