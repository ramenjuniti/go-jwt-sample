package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func signup(w http.ResponseWriter, r *http.Request) {
	u := User{}
	jwt := JWT{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}
	if u.Email == "" {
		errRes(w, http.StatusBadRequest, fmt.Errorf("Email is null or empty"))
		return
	}
	if u.Password == "" {
		errRes(w, http.StatusBadRequest, fmt.Errorf("Password is null or empty"))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}

	u.Password = string(hash)

	trans, err := dbm.Begin()
	if err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}
	if err := trans.Insert(&u); err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}
	if err := trans.Commit(); err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}

	token, err := createToken(u)
	if err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}

	jwt.Token = token

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jwt)
}
