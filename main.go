package main

import (
	"ascii-web-app/src/server"
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", server.HandleHome)
	mux.HandleFunc("/art", server.Handleascii)
	mux.HandleFunc("/ascii-art-switch", server.Handleswitch)

	fmt.Println("running server..http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("error running server..")
		return
	}
}
