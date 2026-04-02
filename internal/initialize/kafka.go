package initialize

import (
	"go-ecommerce/global"
	"log"

	"github.com/segmentio/kafka-go"
)

// Khởi tạo Kafka
var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:19092"),
		Topic:    "otp-auth-topic",    // Topic
		Balancer: &kafka.LeastBytes{}, // Cân bằng tải
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Failed to close kafka producer: %v", err)
	}
}
