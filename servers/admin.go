package servers

import (
	"fmt"
	"log"
	"net/http"
	"pubsubhubbub/handlers"
)

func StartAdminInterface(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("starting admin interface at http://%s\n", addr)
	pshb := http.NewServeMux()
	pshb.HandleFunc("/admin", handlers.PSHBHandler)
	pshb.HandleFunc("/addTopic/", handlers.AddTopic)
	err := http.ListenAndServe(addr, pshb)
	if err != nil {
		log.Println("Error: ", err)
	}
}
