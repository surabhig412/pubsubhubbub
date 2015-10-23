package main
import (
	"fmt"
	"log"
	"net/http"
	"flag"
  "container/list"
)

var topic = make(map[string]*list.List)

func subscribe(w http.ResponseWriter, req *http.Request) {
  fmt.Println("*******in subscribe**********")
  if req.FormValue("mode") == "Subscribe" {
    if _, ok := topic[req.FormValue("topic")]; ok {
      topic[req.FormValue("topic")].PushBack(req.FormValue("callback"))
      fmt.Println("Subscribed...")
    }
  }else if req.FormValue("mode") == "Unsubscribe"{
    if val, ok := topic[req.FormValue("topic")]; ok {
      for e := val.Front(); e != nil; e = e.Next() {
	       if e.Value == req.FormValue("callback"){
           fmt.Println("Found the callback url")
           fmt.Println(e.Value)
           val.Remove(e)
         }
      }
    }
  }
  // fmt.Println(topic["ePayments"].Front().Value)
}

func HTTPListener(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	http.HandleFunc("/subscribe", subscribe)
	log.Println("[HTTP] Starting HTTP interface on", addr)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("[HTTP] ERROR:", err)
	}
}

var (
	host = flag.String("b", "0.0.0.0", "listen on HOST")
	httpPort = flag.Int("p", 8080, "use PORT for HTTP")
)

func main() {
	flag.Parse()
  topic["ePayments"] = list.New()
  topic["RS"] = list.New()
	HTTPListener(*host, *httpPort)
}
