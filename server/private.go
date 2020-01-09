package main

import "net/http"

func private(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))
}
