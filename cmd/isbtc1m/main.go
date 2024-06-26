package main

import (
	"github.com/RaghavSood/isbtc1m/web"
)

func main() {
	webServer := web.NewServer()
	webServer.Serve()
}
