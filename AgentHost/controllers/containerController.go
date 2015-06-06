package controllers

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/fsouza/go-dockerclient"
)

func getPostData(r *http.Request) map[string]string {
	postData := make(map[string]string)
	if "POST" == r.Method {
		r.ParseForm()
		fmt.Println(r.PostForm)
		for k, v := range r.PostForm {
			postData[string(k)] = strings.Join(v,",")
		}
	}
	return postData
}

func CreateContainer(w http.ResponseWriter, r *http.Request) string {
	data := getPostData(r)
	client, _ := docker.NewClient("unix:///var/run/docker.sock")
	config := &docker.Config{
            Image: data["image"],
    }
    containerOptions := docker.CreateContainerOptions{
        Name: data["name"],
        Config: config,
        HostConfig: nil,
    }
    exec, err := client.CreateContainer(containerOptions)
    s, _ := json.Marshal(exec)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}