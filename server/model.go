package main

type User struct {
	ID       int    `json:"id" db:"id, primarykey, autoincrement"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type ErrMes struct {
	Message string `json:"message"`
}
