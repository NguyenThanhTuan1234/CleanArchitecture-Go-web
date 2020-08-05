package main

import (
	"CleanArchitecture-Go-web/driver"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/", driver.Index)
	http.HandleFunc("/user", driver.User)
	http.HandleFunc("/login", driver.Login)
	http.ListenAndServe(":8080", nil)
}
