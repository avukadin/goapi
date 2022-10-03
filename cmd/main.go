package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http" 

	"github.com/go-chi/chi"
	"github.com/avukadin/goapi/pkg/mongocon"
)

type Client struct { Handle string } 

func getClientID(w http.ResponseWriter, r *http.Request){	

	// Parse client handle
	var c Client
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil{
		log.Fatal(err)
	}

	// Connect to mongo
	mongoClient := mongocon.GetMongoClient()
	clientID := mongocon.GetClientID(mongoClient, c.Handle)

	var out = []byte(clientID);
	w.Write(out);
}

func main(){
	r := chi.NewRouter();
	r.Get("/getClientID", getClientID)

	fmt.Println("API is running.")
	http.ListenAndServe(":8080", r)
}
