package main

import (
	"os"

	"github.com/IBM/sarama"
	"github.com/thanakize/skill-api-kafka/consumer/database"
	"github.com/thanakize/skill-api-kafka/consumer/skill"
)

// Sarama configuration options
var (
	broker   = os.Getenv("BROKER")
	version  = sarama.DefaultVersion.String()
	group    = os.Getenv("GROUP")
	topics   = os.Getenv("TOPICS")
	assignor = os.Getenv("ASSIGNOR")
	verbose  = false
)

func init() {
	println(broker, topics)
	if len(broker) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

}

func main() {

	database := database.InitDatabase()
	defer database.CloseDatabase()

	handler := skill.InitHandler(database)
	consumer := skill.NewConsumer(handler, []string{broker}, topics)
	defer consumer.CloseConection()

	consumer.Listen()

}
