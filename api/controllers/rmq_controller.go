package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type RMQ struct {
	ch *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (channel *RMQ) GetChannel() (ch *amqp.Channel) {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Printf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	conn, err := amqp.Dial(os.Getenv("AQMP_URL"))

	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

	return ch

}

func (channel *RMQ) PublishData(data string) {

	ch := channel.GetChannel()
	q, err := ch.QueueDeclare(
		"process_req", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := data
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
