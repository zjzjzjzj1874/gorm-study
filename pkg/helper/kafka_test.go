package helper

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
)

// more test case see: https://github.com/segmentio/kafka-go
func Test_Kafka(t *testing.T) {
	ctx := context.Background()

	t.Run("kafka writer producer", func(t *testing.T) {
		err := KafkaWriter.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte("Key-A"),
				Value: []byte("value-a"),
			},
			kafka.Message{
				Key:   []byte("Key-B"),
				Value: []byte("value-b"),
			},
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("kafka write success\n")
	})

	t.Run("kafka conn", func(t *testing.T) {
		err := KafkaWriter.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte("Key-A"),
				Value: []byte("value-a"),
			},
			kafka.Message{
				Key:   []byte("Key-B"),
				Value: []byte("value-b"),
			},
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("kafka write success\n")
	})

	t.Run("kafka reader consumer", func(t *testing.T) {
		msg, err := KafkaReader.ReadMessage(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("kafka consumer msg:%v success\n", msg)
	})
}
