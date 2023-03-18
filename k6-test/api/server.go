package main

import (
	"log"
	"net/http"

	"github.com/carlosdamazio/k6-test/internal/db"
	"github.com/carlosdamazio/k6-test/internal/rest/handles"
	"github.com/carlosdamazio/k6-test/internal/services/message"
)

func main() {
	http.HandleFunc("/messages", handles.NewMessageHandle(message.New(db.Messages)).MessageHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
