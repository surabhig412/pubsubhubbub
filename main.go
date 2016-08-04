package main

import (
	"flag"
	"pubsubhubbub/servers"
	"pubsubhubbub/utils"
)

var (
	adminHost  = flag.String("s", "0.0.0.0", "listen on Host for admin interface")
	adminPort  = flag.Int("q", 9090, "use port for admin interface")
	host       = flag.String("b", "0.0.0.0", "listen on Host")
	port       = flag.Int("p", 8080, "use PORT for HTTP")
	ConfigFile = flag.String("c", "config.toml", "use Configfile")
)

func main() {
	flag.Parse()
	utils.SetConfigFile(*ConfigFile)
	go servers.StartAdminInterface(*adminHost, *adminPort)
	servers.HTTPListener(*host, *port)
}
