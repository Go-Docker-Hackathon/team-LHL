package controllers

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/fsouza/go-dockerclient"
	"github.com/Go-Docker-Hackathon/team-LHL/AgentHost/resource"
)

func getPostData(r *http.Request) map[string]string {
	postData := make(map[string]string)
	if "POST" == r.Method {
		r.ParseForm()
		for k, v := range r.PostForm {
			postData[string(k)] = strings.Join(v,",")
		}
	}
	return postData
}

func createContainer(image string, goServer string) string {
	client, _ := docker.NewClient("unix:///var/run/docker.sock")
	var commands = []string{"/startGoAgent.sh", goServer}
	config := &docker.Config{
            Image: image,
			Cmd: commands,
    }
    containerOptions := docker.CreateContainerOptions{
        Config: config,
        HostConfig: nil,
		
    }
	fmt.Println("--------creating---image:--", image, "--------")
    exec, _ := client.CreateContainer(containerOptions)
	client.StartContainer(exec.ID, nil)
	return exec.ID
}

func destroyContainer(containers []string) {
	client, _ := docker.NewClient("unix:///var/run/docker.sock")
	for _, contianer := range containers {
		go client.RemoveContainer(docker.RemoveContainerOptions{
			ID: contianer,
			Force: true,
		})
	}
}

func CreateContainer(w http.ResponseWriter, r *http.Request) string {
	data := getPostData(r)
	goServer := strings.Split(r.RemoteAddr,":")[0]
	resources := strings.Split(data["resources"], "|")
	image := resource.GetImage(resources)
	return createContainer(image, goServer)
}