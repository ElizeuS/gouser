package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type Person struct {
	Name    string `json:"Name"`
	Andress string `json:"Andress"`
}

//Função responsável por receber os dados da fila (Queue) RabbitMQ.
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"saveUsersData", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			save(string(d.Body))
		}

	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Método responsável por salvar os dados em arquivos .json
func save(data string) {

	//contém o valor do repositório onde devem serem salvos os arquivos
	str := os.Getenv("NEW_CLIENTS")

	file, _ := json.MarshalIndent(data, "", " ")

	//atualmente o arquivo á salvo com o valor da data atual
	dt := time.Now()
	dir := fmt.Sprintf("%s%s.json", str, dt.String())

	_ = ioutil.WriteFile(dir, file, 0644)
}
