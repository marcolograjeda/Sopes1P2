package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	pb "./comunicacion"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func getmessaje() string {
	conn, err := grpc.Dial("35.192.143.50:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	c := pb.NewComunicandoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.Call(ctx, &pb.Llamada{Body: "Hola"})
	if err == nil {
		fmt.Println(r)
		log.Println(r.GetBody())
	} else {
		log.Println("Comunicacion incorrecta 1", err)
	}

	return `{ "Hola mundo": "Correcto"}`
}

func handlemensajes(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productsWithDiscountApplied := getmessaje()
	json.NewEncoder(w).Encode(productsWithDiscountApplied)
}

type Article struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          string `json:"age"`
	Infectedtype string `json:"infectedtype"`
	State        string `json:"state"`
}

func crearmensaje(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)

	productsWithDiscountApplied := sendmessajepython(string(reqBody))
	json.NewEncoder(w).Encode(productsWithDiscountApplied)
}

func sendmessajepython(as string) string {

	conn, err := grpc.Dial("35.192.143.50:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar : %v", err)
	}
	defer conn.Close()

	c := pb.NewComunicandoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := c.Call(ctx, &pb.Llamada{Body: as})
	if err == nil {
		fmt.Println(r)
		log.Println(r.GetBody())
		return r.GetBody()
	} else {
		log.Println("Comunicacion incorrecta 2", err)
	}

	return `{"Hola mundo": "Correcto"}`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	port := "11080"

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Esta trabajando")
	})

	myRouter.HandleFunc("/rutaManejador", handlemensajes)
	myRouter.HandleFunc("/crear", crearmensaje).Methods("POST")

	fmt.Println("El servidor esta en ejecucion", port)
	http.ListenAndServe(":"+port, myRouter)
}
