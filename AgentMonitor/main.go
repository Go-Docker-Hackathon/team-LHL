package main

import(
	"fmt"
//	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
//	"github.com/fsouza/go-dockerclient"
//	"github.com/Go-Docker-Hackathon/team-LHL/AgentHost/resource"
)

type Agent struct {
	Uuid string
	Agent_name string
	Status string
}

type Job struct {
	Result string
}

type History struct {
	Jobs []Job
}

func getAgents(goServerUrl string) []Agent {
	response, _ := http.Get("http://" + goServerUrl + ":8153/go/api/agents")
	bs, _ := ioutil.ReadAll(response.Body)
	var agents []Agent
	err := json.Unmarshal(bs, &agents)
	fmt.Println("------------------------")
	if err != nil {
		fmt.Println(err)
	}
	return agents
}

func setAgentResources(goServerUrl string) {
	client, _ := docker.NewClient("unix:///var/run/docker.sock")
	for _, agent := range getAgents(goServerUrl) {
		container, _ := client.InspectContainer(agent.Agent_name)
		resources := resource.GetImageTags(container.Image)
		for _, r := range resources {
			http.Post("http://" + goServerUrl + ":8153/go/api/agents-resources/" + agent.Uuid + "/" + r, string(""), nil)		
		}
	}
}

func enablePendingAgents(goServerUrl string) {
	for _, agent := range getAgents(goServerUrl) {
		if agent.Status == "Pending" {
			http.Post("http://" + goServerUrl + ":8153/go/api/agents/" + agent.Uuid + "/enable", string(""), nil)
		}
	}
}

func removeIdleAgents(goServerUrl string) {
	client, _ := docker.NewClient("unix:///var/run/docker.sock")
	for _, agent := range getAgents(goServerUrl) {
		r, _ := http.Get("http://" + goServerUrl + ":8153/go/api/agents/" + agent.Uuid + "/job_run_history/0")
		bs, _ := ioutil.ReadAll(r.Body)
		history := History{}
		json.Unmarshal(bs, &history)
		if agent.Status == "Idle" &&  len(history.Jobs) > 0{
			http.Post("http://" + goServerUrl + ":8153/go/api/agents/" + agent.Uuid + "/disable", string(""), nil)
			http.Post("http://" + goServerUrl + ":8153/go/api/agents/" + agent.Uuid + "/delete", string(""), nil)
			client.KillContainer(agent.Agent_name)
			client.RemoveContainer(agent.Agent_name)
		}
	}
}


func main (){
	var goServerUrl = "10.18.7.153"
	for {
		removeIdleAgents(goServerUrl)
		enablePendingAgents(goServerUrl)
		setAgentResources(goServerUrl)
		time.Sleep(1000)
	}
	fmt.Println(getAgents(goServerUrl))
//http.Post("http://10.18.5.26:8153/go/api/agents/")
}