package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"sync"

	"github.com/streadway/amqp"
)

func startQueueListener(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := amqp.Dial(settings.QueueConnectionString)
	if err != nil {
		log.Println(err)
	}
	receive(conn, settings.QueueChannel, wg)
}

func receive(conn *amqp.Connection, queueName string, wg *sync.WaitGroup) {
	// rabbitConnect()
	defer conn.Close()

	for {
		ch, err := conn.Channel()
		if err != nil {
			fmt.Println(err, "Failed to open a channel")
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			queueName, // name
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		autoAck := settings.AutoAck

		message, err := ch.Consume(q.Name, "", autoAck, false, false, false, nil)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(500 * time.Millisecond)

		for d := range message { // the d stands for Delivery
			fmt.Println(string(d.Body[:]))
			messageHandler(queueName, d.Body, wg)
		}
	}
}

func messageHandler(queueName string, message []byte, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	log.Println(queueName, message)
	m := pdfRequest{}
	json.Unmarshal(message, &m)
	go hookForAMQP(&m)
}
