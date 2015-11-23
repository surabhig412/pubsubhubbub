package main
import (
	"fmt"
	"log"
	"net/http"
	"flag"
	"container/list"
	"encoding/json"
	sr "github.com/tuvistavie/securerandom"
)

var topic = make(map[string]*list.List)

func subscribe(w http.ResponseWriter, req *http.Request) {
	fmt.Println("*******in subscribe**********")
	req.ParseForm()
	random,_ := sr.Base64(10, true)
	if req.PostFormValue("mode") == "Subscribe" {
		if _, ok := topic[req.PostFormValue("topic")]; ok {
			topic[req.PostFormValue("topic")].PushBack(req.PostFormValue("callback"))
			fmt.Println("Subscribed...")
			response := "SUBSCRIPTION ACCEPTED FOR VERIFICATION"
			out,_ :=json.MarshalIndent(response, "202", "    ")
			w.Write([]byte(out))
			http.Get("http://"+ req.PostFormValue("callback") +"/verify?topic=" + req.PostFormValue("topic") +"&challenge=" + random)
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

			func verify(w http.ResponseWriter, req *http.Request) {
				fmt.Println("In subscription server verify method")
				fmt.Println(req.FormValue("topic"))
				http.Get("http://localhost:8080/verify_intent?topic=" + req.FormValue("topic") +"&challenge=" + req.FormValue("challenge"))
			}

			func verify_intent(w http.ResponseWriter, req *http.Request) {
				fmt.Println("In server verify intent method")
				fmt.Println(req.FormValue("topic"))
				fmt.Println(req.FormValue("challenge"))
			 	http.Redirect(w, req, "http://localhost:8080", 200)
			}

			func all(w http.ResponseWriter, req *http.Request) {
				fmt.Println("in redirection")
					w.Write([]byte("Verified"))
			}


			func HTTPListener(host string, port int) {
				addr := fmt.Sprintf("%s:%d", host, port)
				http.HandleFunc("/subscribe", subscribe)
				http.HandleFunc("/", all)
				http.HandleFunc("/verify_intent", verify_intent)
				log.Println("[HTTP] Starting HTTP interface on", addr)
				err := http.ListenAndServe(":8080", nil)
				if err != nil {
					log.Fatal("[HTTP] ERROR:", err)
				}
			}
			func SubscriptionServer(host string, port int){
				addr := fmt.Sprintf("%s:%d", host, port)
				http.HandleFunc("/verify", verify)
				log.Println("[HTTP] Starting Subscriber interface on", addr)
				err := http.ListenAndServe(":3000", nil)
				if err != nil {
					log.Fatal("[HTTP] ERROR:", err)
				}
			}

			var (
				host = flag.String("b", "0.0.0.0", "listen on HOST")
				httpPort = flag.Int("p", 8080, "use PORT for HTTP")
				subscriberPort = flag.Int("s", 3000, "use PORT for Subscriber")
			)

			func main() {
				flag.Parse()
				topic["ePayments"] = list.New()
				topic["RS"] = list.New()
				go HTTPListener(*host, *httpPort)
				SubscriptionServer(*host, *subscriberPort)
			}
