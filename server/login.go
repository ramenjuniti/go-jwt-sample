package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
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

	password := u.Password

	row := User{}
	if err := dbm.SelectOne(&row, `SELECT id, email, password FROM user WHERE email = ?`, u.Email); err != nil {
		errRes(w, http.StatusBadRequest, err)
		return
	}

	hashedPassword := row.Password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		errRes(w, http.StatusUnauthorized, err)
		return
	}

	token, err := createToken(u)
	if err != nil {
		errRes(w, http.StatusInternalServerError, err)
		return
	}

	jwt.Token = token

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jwt)
}
