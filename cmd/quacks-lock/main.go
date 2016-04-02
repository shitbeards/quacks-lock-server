package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8081", "http service address")

func main() {
	flag.Parse()
	go h.run()
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
