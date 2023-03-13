package main

import (
	f "kafka-go/functions"
)

func main() {

	brokers := []string{"3.69.23.63:9092"}

	/* 	conn, err := kafka.DialLeader(context.Background(), "tcp", "3.69.23.63:9092", "my-topic", 0)
	   	if err != nil {
	   		log.Fatal("failed to dial leader:", err)
	   	} */

	//todo: PRODUCER
	//f.Producer(conn)

	//todo: CONSUME
	f.Consumer(brokers)

	//todo: CREATE TOPICS

}
