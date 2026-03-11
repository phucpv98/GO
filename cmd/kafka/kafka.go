package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	KafkaURL   = "localhost:19092"
	KafkaTopic = "user_topic_vip"
)

// for Consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,           // Danh sách các broker Kafka
		GroupID:        groupID,           // GroupID giúp quản lý offset và phân phối tải giữa các consumer trong cùng một nhóm
		Topic:          topic,             // Tên topic mà consumer sẽ đọc
		MinBytes:       10e3,              // 10KB
		MaxBytes:       10e6,              // 10MB
		CommitInterval: time.Second,       // Tự động commit offset sau mỗi giây
		StartOffset:    kafka.FirstOffset, // Bắt đầu đọc từ đầu topic nếu chưa có offset nào được commit
	})
}

// for Producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL), // Địa chỉ của Kafka broker
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // Cân bằng tải giữa các broker
	}
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua ban chung khoan
func newStock(msg, typ string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typ

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["info"] = s
	body["action"] = "action"

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	// Viết message bởi Producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("Failed to write message to Kafka: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Message sent successfully"})
}

// Consumer hóng mua ATC và bán ATC
func RegisterConsumerATC(id int) {
	kafkaGroup := "consumer-group-"
	reader := getKafkaReader(KafkaURL, KafkaTopic, kafkaGroup)
	defer reader.Close() // Đảm bảo đóng reader khi không còn sử dụng - giải phóng tài nguyên sau khi hoàn thành RegisterConsumerATC.

	// Bắt đầu vòng lặp để liên tục đọc tin nhắn từ Kafka - coi những ai hóng chuyện.
	fmt.Printf("Consumer(%d) Hong Phien ATC::\n", id)
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Failed to read message from Kafka: %v", err)
			continue
		}

		fmt.Printf(
			"Consumer(%d), hong topic:%v, partition:%v, offset:%v, time:%d %s = %s\n",
			id, msg.Topic, msg.Partition, msg.Offset, msg.Time.Unix(), string(msg.Key), string(msg.Value),
		)
	}
}

func main() {
	r := gin.Default() // Khởi tạo Kafka producer một lần duy nhất khi ứng dụng khởi động - tái sử dụng producer cho tất cả các yêu cầu.
	kafkaProducer = getKafkaWriter(KafkaURL, KafkaTopic)
	defer kafkaProducer.Close() // Đảm bảo đóng producer khi không còn sử dụng - giải phóng tài nguyên sau khi hoàn thành main.

	r.POST("action/stock", actionStock)

	// đăng ký 2 user để mua Stock trong ATC (id=1) (id=2)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run(":8999")
}

// ex: curl -X POST "http://localhost:8999/action/stock?msg=HPG&type=MUA"

// CDC: Change Data Capture - là một kỹ thuật trong quản lý dữ liệu, cho phép theo dõi và ghi lại các thay đổi xảy ra trong cơ sở dữ liệu. Khi có sự thay đổi (chèn, cập nhật, xóa) trong cơ sở dữ liệu, CDC sẽ ghi lại thông tin về thay đổi đó, bao gồm loại thay đổi, dữ liệu trước và sau khi thay đổi, thời gian thay đổi, v.v. Điều này giúp các hệ thống khác có thể phản ứng kịp thời với những thay đổi này, ví dụ như cập nhật dữ liệu trong hệ thống phân tích hoặc đồng bộ hóa dữ liệu giữa các hệ thống khác nhau.
// => Debezium : là một nền tảng mã nguồn mở được sử dụng để thực hiện CDC. Nó hỗ trợ nhiều loại cơ sở dữ liệu khác nhau như MySQL, PostgreSQL, MongoDB, v.v. Debezium có thể ghi lại các thay đổi trong cơ sở dữ liệu và gửi chúng đến các hệ thống khác thông qua Kafka hoặc các hệ thống xử lý dữ liệu khác. Điều này giúp các nhà phát triển dễ dàng xây dựng các ứng dụng phản ứng với sự thay đổi dữ liệu trong thời gian thực.
