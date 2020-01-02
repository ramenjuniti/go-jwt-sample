package main

import (
	"encoding/json"
	"net/http"
)

func errRes(w http.ResponseWriter, status int, err error) {
	errMes := ErrMes{Message: err.Error()}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errMes)
}
