package helper

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

var KafkaClientConn *kafka.Conn
var KafkaWriter *kafka.Writer
var KafkaReader *kafka.Reader

func init() {
	// TODO put these configs into yaml
	topic := "my-topic"
	partition := 0

	var err error
	KafkaClientConn, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		panic(err)
	}

	_ = KafkaClientConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_ = KafkaClientConn.SetReadDeadline(time.Now().Add(10 * time.Second))
}

// 可以启动多个kafka的writer和reader.By the way:如果是以reader或者writer启动的,失败后请记得调用Close的方法
func init() {
	KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(""),
		Topic:    "",
		Balancer: &kafka.LeastBytes{},
	}

	rc := kafka.ReaderConfig{
		Brokers:        []string{"brokers"},
		GroupID:        "GroupID",
		Topic:          "Topic",
		MinBytes:       10e3, //10KB
		MaxBytes:       10e6, //10MB
		CommitInterval: time.Second,
	}

	KafkaReader = kafka.NewReader(rc)
}
