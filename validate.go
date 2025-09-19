package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const maxChirpLength = 140

func handlerValidate(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type respVals struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondError(w, "Couldn't decode paramters", http.StatusInternalServerError, err)
		return
	}
	if len(params.Body) > maxChirpLength {
		respondError(w, "Chirp is too long", http.StatusBadRequest, nil)
		return
	}

	respondJSON(w, http.StatusOK, respVals{
		Valid: true,
	})
}

func respondError(w http.ResponseWriter, msg string, code int, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResp struct {
		Error string `json:"error"`
	}
	respondJSON(w, code, errorResp{
		Error: msg,
	})

}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}