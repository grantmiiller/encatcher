package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func PayloadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	vars := mux.Vars(r)
	var payload string
	if httpPayload, ok := vars["payload"]; ok {
		payload = httpPayload
	} else {
		fmt.Println("Could not find payload")
		return
	}
	rawDecodedText, err := base64.StdEncoding.DecodeString(payload)

	fmt.Println("====================================")
	if err != nil {
		fmt.Printf("ERROR: could not decode payload of '%s': %v\n", payload, err)
		return
	}
	fmt.Printf("DECODED: %s\n", rawDecodedText)
}

func BlanketAccept(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{payload}", PayloadHandler)
	r.PathPrefix("/").HandlerFunc(BlanketAccept)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
