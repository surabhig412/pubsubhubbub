package servers

import (
	"fmt"
	"log"
	"net/http"
	"pubsubhubbub/handlers"
)

func HTTPListener(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	http.HandleFunc("/new", handlers.New)
	http.HandleFunc("/subscribe", handlers.Pshb)
	// http.HandleFunc("/publish", handlers.Publish)
	http.HandleFunc("/assets/", handlers.AssetHandler)
	log.Println("[HTTP] Starting HTTP interface on", addr)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("[HTTP] ERROR:", err)
	}
}
