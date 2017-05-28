package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/streadway/amqp"
)

type pdfRequest struct {
	URL            string `json:"url"`
	Grayscale      bool   `json:"grayscale"`
	Landscape      bool   `json:"landscape"`
	TargetFileName string `json:"targetFileName"`
	TargetFileDest string `json:"targetFileDest"`
}

func startQueueListener(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := amqp.Dial(settings.QueueConnectionString)
	if err != nil {
		log.Println(err)
	}
	rabbitReceive(conn, settings.QueueChannel, wg)
}

// func rabbitConnect() *amqp.Connection {
// 	conn, err := amqp.Dial(settings.QueueConnectionString)
// 	if err != nil {
// 		for conn == nil {
// 			fmt.Println(err, "Waiting 15 seconds and attempting to connect to RabbitMQ again.")
// 			time.Sleep(time.Duration(15) * time.Second)
// 			conn, err = amqp.Dial(settings.QueueConnectionString)
// 		}
// 	}
// }

func rabbitSend(conn *amqp.Connection, queueName string, body string) {
	// rabbitConnect()
	defer conn.Close()

	fmt.Println("Sending", body, "to", queueName)
	ch, ERR := conn.Channel()
	if ERR != nil {
		fmt.Println(ERR, "Unable to open channel.")
		return
	}

	q, ERR := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if ERR != nil {
		fmt.Println(ERR, "Failed to declare a queue")
		return
	}

	ERR = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if ERR != nil {
		fmt.Println(ERR, "Failed to publish message")
	}

}

func rabbitReceive(conn *amqp.Connection, queueName string, wg *sync.WaitGroup) {
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
	spew.Dump(message)
	json.Unmarshal(message, &m)
	go hookForAMQP(&m)
}
