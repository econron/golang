package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

type Message struct {
	ServerID int64 `json:"id"`
	Order int64 `json:"order"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		var m Message
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("request body #%v", m)
		defer r.Body.Close()
	})
	log.Fatal(http.ListenAndServe(":8001", nil))
}