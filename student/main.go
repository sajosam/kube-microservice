package main

import (
	"fmt"
	"net/http"
)

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List of students")
}

func main() {
	http.HandleFunc("/list-users", listUsersHandler)
	http.ListenAndServe(":8080", nil)
}
