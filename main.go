package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const port int = 8080

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

var users []User = []User{
	{ID: 1, Username: "jane"},
	{ID: 2, Username: "joe"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&newUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	users = append(users, newUser)
	usersMsg := "updated users:"
	for _, user := range users {
		usersMsg += fmt.Sprintf(" %s", user.Username)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(usersMsg))
}

func main() {
	http.HandleFunc("GET /users", getUsers)
	http.HandleFunc("POST /users", addUser)

	http.Handle("GET /metrics", promhttp.Handler())

	log.Println("Server staring on port:", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
