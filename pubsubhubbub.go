package main
import (
  "fmt"
  "log"
  "net/http"
  "flag"
  "container/list"
  "encoding/json"
)

var topic = make(map[string]*list.List)

func subscribe(w http.ResponseWriter, req *http.Request) {
  fmt.Println("*******in subscribe**********")
    req.ParseForm()
    if req.PostFormValue("mode") == "Subscribe" {
      if _, ok := topic[req.PostFormValue("topic")]; ok {
        topic[req.PostFormValue("topic")].PushBack(req.PostFormValue("callback"))
        fmt.Println("Subscribed...")
        response := "SUBSCRIPTION ACCEPTED"
        out,_ :=json.MarshalIndent(response, "202", "    ")
        w.Write([]byte(out))

       }else{
         response := "BAD REQUEST"
         out,_ :=json.MarshalIndent(response, "400", "    ")
         w.Write([]byte(out))
       }
      }else if req.PostFormValue("mode") == "Unsubscribe"{
				flag:= false
        if val, ok := topic[req.PostFormValue("topic")]; ok {
          for e := val.Front(); e != nil; e = e.Next() {
            if e.Value == req.PostFormValue("callback"){
              fmt.Println("Found the callback url")
              fmt.Println(e.Value)
              val.Remove(e)
							response := "UNSUBSCRIPTION REQUEST ACCEPTED"
              out,_ :=json.MarshalIndent(response, "200", "    ")
              w.Write([]byte(out))
							flag = true
							break
            }
          }
        }
        if(flag==false) {
					response := "BAD REQUEST"
          out,_ :=json.MarshalIndent(response, "400", "    ")
          w.Write([]byte(out))
				}

      }else{
        response := "Bad Request"
        out,_ :=json.MarshalIndent(response, "400", "    ")
        w.Write([]byte(out))
      }
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
