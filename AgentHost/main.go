package main

import(
	"github.com/dpapathanasiou/go-api"
)

func main() {
    api.NewServer("10.18.2.10", api.DefaultServerTransport, 9001, api.DefaultServerReadTimeout, false, route().handlers)
}