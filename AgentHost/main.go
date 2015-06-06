package main

import(
	"net/http"
	"github.com/dpapathanasiou/go-api"
)

type Route struct{
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

func routes() Route {
	var handlers = map[string]func(http.ResponseWriter, *http.Request){}
    
    handlers["/create/"] = func(w http.ResponseWriter, r *http.Request) {
        api.Respond("application/json", "utf-8", createContainer)(w, r)
    }

    return Route {
    	handlers: handlers,
    }
}

func main() {
    api.NewServer("localhost", api.DefaultServerTransport, 9001, api.DefaultServerReadTimeout, false, routes().handlers)
}

