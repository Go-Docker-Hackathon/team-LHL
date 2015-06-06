package resource

import(
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Image struct {
	Imageid string
	Tags []string
}

func getImagesList() []Image {
	var images []Image
	file,_ := ioutil.ReadFile("/home/xkli/go/src/github.com/Go-Docker-Hackathon/team-LHL/AgentHost/dockerimagesjsonfile/dockerimages.josn")	
	err := json.Unmarshal(file, &images)
	if err != nil {
		fmt.Println(err)
	}
	return images
}

func contains(item string, list []string) bool {
	var contain = false
    for _, b := range list {
        if b == item {
            contain = true
			break
        }
    }
    return contain
}

func tagsMatch(tags []string, resources []string) bool {
	for _, tag := range tags {
		if !contains(tag, resources) {
			return false
		}
	}
	return true
}

func GetImage(resources []string) string {
	images := getImagesList()
	
	for _, image := range images {
		if(tagsMatch(resources, image.Tags)){
			return image.Imageid
		}
	}
	
	return ""
}