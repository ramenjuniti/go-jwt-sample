package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var user *User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}
	if user.Email == "" {
		errRes(w, http.StatusBadRequest, fmt.Errorf("Email is null or empty"))
		return
	}
	if user.Password == "" {
		errRes(w, http.StatusBadRequest, fmt.Errorf("Password is null or empty"))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}

	user.Password = string(hash)

	trans, err := dbm.Begin()
	if err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}
	if err := trans.Insert(user); err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}
	if err := trans.Commit(); err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}

	user.Password = "[FILTERED]"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
