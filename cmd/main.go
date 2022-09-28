package main

import (
	"log"
	"fmt"
	"encoding/json"
	"net/http" 
	"github.com/go-chi/chi"
	"github.com/avukadin/goapi/pkg/mongocon"
)

type Client struct { Handle string } 
func getArticle(w http.ResponseWriter, r *http.Request){	

	// Parse client handle
	var c Client
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil{
		log.Fatal(err)
	}

	// Connect to mongo
	mongoClient:=mongocon.GetMongoClient()
	clientID := mongocon.GetClientID(mongoClient, c.Handle)

	var out = []byte(clientID);
	w.Write(out);
	// fmt.Println(string(json.Marshal(c)))
	// fmt.Println(c.Handle)
	fmt.Println("Done!")
}

func main(){
	r := chi.NewRouter();
	r.Get("/getArticle", getArticle)

	http.ListenAndServe(":8080", r)
}
