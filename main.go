package main

import (
	"flag"
	"log"
	"net/http"
    "os"
)

var addr = flag.String("addr", "8081", "http service address")

func main() {
    port := os.Getenv("PORT")
    if port == "" {
		port = *addr
	}
	flag.Parse()
	go h.run()
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
