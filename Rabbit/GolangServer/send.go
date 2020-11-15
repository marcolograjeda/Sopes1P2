package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func crearmensaje(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)

	productsWithDiscountApplied := sendmessajepython(string(reqBody))
	json.NewEncoder(w).Encode(productsWithDiscountApplied)
}

func sendmessajepython(as string) string {

	conn, err := amqp.Dial("amqp://guest:guest@104.154.53.252:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"lab-so1", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(as),
		})
	failOnError(err, "Failed to publish a message")

	return `{ "status": "done"}`
}

func enableCors(w *http.ResponseWriter) {
	// (w).Header().Set("Access-Control-Allow-Origin", "")
}

func main() {
	port := "11090"

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "It is working.")
	})

	myRouter.HandleFunc("/insert", crearmensaje).Methods("POST")

	fmt.Println("Server running on", port)
	http.ListenAndServe(":"+port, myRouter)
}
