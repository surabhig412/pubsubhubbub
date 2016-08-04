package handlers

import (
	"log"
	"net/http"
	"pubsubhubbub/utils"
	"pubsubhubbub/views"
)

var topicsAndCallBacksList = make(map[string][]string)

func New(w http.ResponseWriter, req *http.Request) {
	Data := utils.GetData()
	var DataMap = make(map[string]interface{})
	DataMap["data"] = Data
	views.Subscription(w, views.NewPage, views.PartialTemp, DataMap)
}

func Pshb(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	topicList := req.Form["topic"]
	callback := req.PostFormValue("subscription_callback_url")
	if req.PostFormValue("subscribe_option") == "1" {
		subscription(topicList, callback)
	} else if req.PostFormValue("unsubscribe_option") == "1" {
		unsubscribe(topicList, callback, topicsAndCallBacksList)
	}
}

func subscription(topics []string, callback string) {
	for _, j := range topics {
		topicsAndCallBacksList[j] = append(topicsAndCallBacksList[j], callback)
	}
	log.Println("SUBSCRIBED", topicsAndCallBacksList)
}

func unsubscribe(topic []string, callback string, topicsAndCallBacksList map[string][]string) {
	for _, value := range topic {
		if callbackList, ok := topicsAndCallBacksList[value]; ok {
			for index, item := range callbackList {
				if item == callback {
					callbackList = append(callbackList[:index], callbackList[index+1:]...)
					break
				}
			}
			topicsAndCallBacksList[value] = callbackList
		}
	}
}

// func Publish(w http.ResponseWriter, req *http.Request) {
// 	if val, ok := topic[req.PostFormValue("topic")]; ok {
// 		for e := val.Front(); e != nil; e = e.Next() {
// 			fmt.Println("************************")
// 			apiUrl := "http://localhost:4567"
// 			resource := "/publish"
// 			data := make(map[string]string)
// 			data["topic"] = req.PostFormValue("topic")
// 			data["content"] = req.PostFormValue("content")
// 			json_data, _ := json.Marshal(data)
// 			fmt.Println(string(json_data))
// 			fmt.Println("*********@@@@@@@@***************")
// 			text := string(json_data)
// 			u, _ := url.ParseRequestURI(apiUrl)
// 			u.Path = resource
// 			urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/"
// 			client := &http.Client{}
// 			r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(text)) // <-- URL-encoded payload
// 			r.Header.Add("Content-Type", "application/json")
// 			//  r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
// 			resp, _ := client.Do(r)
// 			fmt.Println(resp)
// 		}
// 	}
// }
//
// func Subscription(w http.ResponseWriter, req *http.Request) {
// 	req.ParseForm()
// 	random, _ := sr.Base64(10, true)
// 	if req.PostFormValue("mode") == "Subscribe" {
// 		if _, ok := topic[req.PostFormValue("topic")]; ok {
// 			topic[req.PostFormValue("topic")].PushBack(req.PostFormValue("callback"))
// 			fmt.Println("Subscribed...")
// 			response := "SUBSCRIPTION ACCEPTED FOR VERIFICATION"
// 			out, _ := json.MarshalIndent(response, "202", " ")
// 			w.Write([]byte(out))
// 			//sinatra subscriber server running at port 4567
// 			resp, err := http.Get("http://" + req.PostFormValue("callback") + "/verify?topic=" + req.PostFormValue("topic") + "&challenge=" + random)
// 			if err != nil {
// 				//
// 			}
// 			data := map[string]interface{}{}
// 			defer resp.Body.Close()
// 			body, _ := ioutil.ReadAll(resp.Body)
// 			json.Unmarshal(body, &data)
// 			// fmt.Println("The json response received is: %s \n", data["challenge"])
// 			// fmt.Println("topic:%s \n", data["topic"])
// 			if data["challenge"] == random {
// 				fmt.Printf("The Subscriber has been authorized\n")
// 			} else {
// 				fmt.Printf("The Subscriber is not authorized")
// 			}
// 		} else {
// 			response := "BAD REQUEST"
// 			out, _ := json.MarshalIndent(response, "400", " ")
// 			w.Write([]byte(out))
// 		}
// 	} else if req.PostFormValue("mode") == "Unsubscribe" {
// 		flag := false
// 		if val, ok := topic[req.PostFormValue("topic")]; ok {
// 			for e := val.Front(); e != nil; e = e.Next() {
// 				if e.Value == req.PostFormValue("callback") {
// 					fmt.Println("Found the callback url")
// 					val.Remove(e)
// 					response := "UNSUBSCRIPTION REQUEST ACCEPTED"
// 					out, _ := json.MarshalIndent(response, "200", " ")
// 					w.Write([]byte(out))
// 					flag = true
// 					break
// 				}
// 			}
// 		}
// 		if flag == false {
// 			response := "BAD REQUEST"
// 			out, _ := json.MarshalIndent(response, "400", " ")
// 			w.Write([]byte(out))
// 		}
// 	} else {
// 		response := "Bad Request"
// 		out, _ := json.MarshalIndent(response, "400", " ")
// 		w.Write([]byte(out))
// 	}
//
// }
