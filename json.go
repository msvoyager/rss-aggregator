package main //cuz this should be in the main package

import (
	"encoding/json"
	"net/http"
	"log"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload) //convert into json string as bytes so we can write it in a binary format directly into http request
	if err != nil {
		log.Printf("\nfailed to marshal json response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
		 
	}
	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}
