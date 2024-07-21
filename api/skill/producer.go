package skill

import (
	"encoding/json"
	"log"
	"os"

	_ "net/http/pprof"

	"github.com/rcrowley/go-metrics"

	"github.com/IBM/sarama"
)

// Sarama configuration options
var (
	broker              = os.Getenv("BROKER")
	version             = sarama.DefaultVersion.String()
	topic               = os.Getenv("TOPIC")
	producers           = 1
	verbose             = false
	recordsNumber int64 = 1
	recordsRate         = metrics.GetOrRegisterMeter("records.rate", nil)
)

type Producer struct {
	sarama sarama.SyncProducer
	topic  string
}

func CreateProducer(topic string) *Producer {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalln(err)
	}
	return &Producer{
		sarama: producer,
		topic:  topic,
	}
}

func (producer Producer) CloseProducer() {
	if err := producer.sarama.Close(); err != nil {
		log.Fatalln(err)
	}
}

func (producer Producer) ProduceMessage(key string, skill Skill, skillKey string) (Skill, error) {
	stringJson, err := json.Marshal(skill)
	if err != nil {
		return Skill{}, err
	}

	msg := &sarama.ProducerMessage{Topic: producer.topic, Value: sarama.StringEncoder(stringJson), Key: sarama.StringEncoder(key), Headers: []sarama.RecordHeader{
		sarama.RecordHeader{
			Key:   []byte("key"),
			Value: []byte(skillKey),
		},
	}}
	_, _, err = producer.sarama.SendMessage(msg)
	if err != nil {
		return Skill{}, err
	}
	return skill, nil
}
