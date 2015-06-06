package main

import(
	"github.com/dpapathanasiou/go-api"
)

func main() {
    api.NewServer("localhost", api.DefaultServerTransport, 9001, api.DefaultServerReadTimeout, false, route().handlers)
}