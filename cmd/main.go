package main

import (
	"net/http" 
	"github.com/go-chi/chi"
	"github.com/avukadin/goapi/pkg/mongocon"
)


func getArticle(w http.ResponseWriter, r *http.Request){
	var out = []byte("Hello from getArticle");
	w.Write(out);
}

func main(){
	mongocon.getMongoConnection()
	r := chi.NewRouter();
	r.Get("/getArticle", getArticle)

	// http.ListenAndServe(":8080", r)
}
