package main

import (
	"encoding/json"
	"net/http"
)

func login(w http.ResponseWriter, req *http.Request) {
	type User struct {
		Username string
	}

	user := User{"Ned"}
	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/auth/login", login)
	http.Handle("/", http.FileServer(http.Dir("./client/public")))

	panic(http.ListenAndServe(":8080", nil))
}
