package views

import (
	"net/http"
	"text/template"
)

var topicList []string

func RenderSubscription(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.New("view").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Subscription(w http.ResponseWriter, tmpl, partial string, data map[string]interface{}) {
	t := template.New("view")
	template.Must(t.Parse(tmpl))
	template.Must(t.Parse(partial))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddTopic(topic string) []string {
	topicList = append(topicList, topic)
	return topicList
}
