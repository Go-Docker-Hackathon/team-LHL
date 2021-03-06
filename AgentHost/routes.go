package main

import(
	"net/http"
	"github.com/dpapathanasiou/go-api"
	"github.com/Go-Docker-Hackathon/team-LHL/AgentHost/controllers"
)

type Route struct{
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

func route() Route {
	var handlers = map[string]func(http.ResponseWriter, *http.Request){}
    
    handlers["/create"] = func(w http.ResponseWriter, r *http.Request) {
        api.Respond("application/json", "utf-8", controllers.CreateContainer)(w, r)
    }
	
    return Route {
    	handlers: handlers,
    }
}