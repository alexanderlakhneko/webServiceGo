package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/hello/", helloHandler)

	s := http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

type Resp struct {
	Message string
	Error   string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := strings.Replace(r.URL.Path, "/hello/", "", 1)

	resp := Resp{
		Message: fmt.Sprintf("hello, %s. Glad to see you again", name),
	}

	respJson, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)

	w.Write(respJson)
}
