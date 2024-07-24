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
	Handler                 HandlerInterface
}

func NewConsumer(handler HandlerInterface, brokers []string, topics string) *Consumer {
	consumer, err := sarama.NewConsumer(brokers, sarama.NewConfig())

	if err != nil {
		panic("error try to new consumer" + err.Error())
	}

	partitionConsumer, err := consumer.ConsumePartition(topics, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err.Error())
	}
	return &Consumer{
		SaramaConsumer:          consumer,
		SaramaPartitionConsumer: partitionConsumer,
		Handler:                 handler,
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
