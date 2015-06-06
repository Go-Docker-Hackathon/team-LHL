package main

import(
	"github.com/dpapathanasiou/go-api"
        "os"
)

func main() {
     hostServer := "10.18.2.10"
     if len(os.Args) > 1 { hostServer = os.Args[1] }
     api.NewServer(hostServer, api.DefaultServerTransport, 9001, api.DefaultServerReadTimeout, false, route().handlers)
}