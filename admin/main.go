package main

import (
	"fmt"
	"net/http"
)

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List of admins")
}

func main() {
	http.HandleFunc("/list-users", listUsersHandler)
	http.ListenAndServe(":8082", nil)
}
