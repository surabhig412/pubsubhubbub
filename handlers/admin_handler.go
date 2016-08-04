package handlers

import (
	"net/http"
	"pubsubhubbub/utils"
	"pubsubhubbub/views"
)

func AddTopic(w http.ResponseWriter, req *http.Request) {
	conf := utils.GetStruct()
	conf.Publisher = req.FormValue("publisher")
	topic1 := req.FormValue("topic1")
	topic2 := req.FormValue("topic2")
	topic3 := req.FormValue("topic3")

	conf.Topics = append(conf.Topics, topic1, topic2, topic3)
	conf.Flush()
	views.RenderSubscription(w, views.SuccessPage, nil)
}

// func deleteTopic(index string) {
// 	conf := utils.GetData()
// 	i, err := strconv.Atoi(index)
// 	if err != nil {
// 		log.Println("ERROR:", "strconv failed for ", index, err.Error())
// 	}
// 	conf.Topics = append(conf.Topics[:i], conf.Topics[i+1:]...)
// 	conf.Flush()
// }

func PSHBHandler(w http.ResponseWriter, req *http.Request) {
	conf := utils.GetData()
	views.RenderSubscription(w, views.AdminPage, conf)
}
