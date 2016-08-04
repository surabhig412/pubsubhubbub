package utils

import "container/list"

var (
	topic  = make(map[string]*list.List)
	random string
)

func initialise() {
	topic["ePayments"] = list.New()
	topic["RS"] = list.New()
}
