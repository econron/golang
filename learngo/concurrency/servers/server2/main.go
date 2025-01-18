package main

import (
	"net/http"
	"fmt"
	"os"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(os.Stdout, "%#v", r)
	})
	log.Fatal(http.ListenAndServe(":8002", nil))
}