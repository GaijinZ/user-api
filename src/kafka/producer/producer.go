package producer

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	KafkaCtx    context.Context
	KafkaWriter *kafka.Writer
)

const (
	topic         = "messages"
	brokerAddress = "kafka-1:9092"
)

func getKafkaWriter() *kafka.Writer {
	return KafkaWriter
}

func getKafkaCtx() context.Context {
	return KafkaCtx
}

func ProduceMessage(k, val string) error {
	w := getKafkaWriter()
	ctx := getKafkaCtx()

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(k),
		// create an arbitrary message payload for the value
		Value: []byte(val),
		Time:  time.Now(),
	})
	if err != nil {
		fmt.Println("could not write message " + err.Error())
		return err
	}
	return nil
}

func Setup() {
	KafkaCtx = context.Background()

	l := log.New(os.Stdout, "kafka writer: ", 0)

	KafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"brokerAddress"},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})
}
