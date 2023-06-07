package functions

import (
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Producer(conn *kafka.Conn) {
	/* conn := kafka.WriterConfig{
		Brokers:   brokers,             //mesajları barındıran, işleyen ve alıp gönderen sunucular
		Topic:     "quickstart-events", //mesajların kategorize edilmesini sağlar, her konu birden fazla broker'da olabilir, topic içindeki bir veye birden fazla partition'a ayrılır.
		Balancer:  &kafka.Hash{},       //balancer, partition'ları brokerlar arasında yeniden dağıtarak veri yükünü dengeler
		BatchSize: 1,                   //her mesaj için ayrı bir istek gönder
		//BatchBytes:       1,                    //her mesaj için ayrı bir istek gönder
		CompressionCodec: kafka.Snappy.Codec(), //mesajların sıkıştırılması için kullanılan araç. Kafka sıkıştırma için seçenekler sunar (örn: Snappy)
	}

	//*Yazar oluştur
	w := kafka.NewWriter(conn)
	defer w.Close() */

	//*Mesaj yaz
	message := kafka.Message{
		Key:   []byte("my-key"),
		Value: []byte("hello kafkaaa"),
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(message)
	if err != nil {
		panic("could not write \n" + err.Error())
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
