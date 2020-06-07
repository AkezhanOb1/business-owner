package mq


import (
	"encoding/json"
	pb "github.com/AkezhanOb1/business-owner/api/mq"
	"github.com/AkezhanOb1/business-owner/config"
	"github.com/streadway/amqp"
)

//BusinessRegistrationEmail is a func
func BusinessRegistrationEmail(email string) error{
	conn, err := amqp.Dial(config.RabbitConnection)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		config.WelcomeEmilQueue, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	req := pb.WelcomeEmailRequest{
		Email:	email,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		return err
	}

	return nil
}
