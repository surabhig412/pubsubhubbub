pubsubhubbub: main.go assets.go
	go build
assets.go: assets/
	go-bindata -pkg handlers -o handlers/assets.go assets/...
