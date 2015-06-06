package resource

import(
	"fmt"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Image struct {
	Imageid string
	Tags []string
}

func getImagesList() []Image {
	var images []Image
	absPath, _ := filepath.Abs("./dockerimagesjsonfile/dockerimages.json")
	file,_ := ioutil.ReadFile(absPath)	
	err := json.Unmarshal(file, &images)
	if err != nil {
		fmt.Println(err)
	}
	return images
}



func contains(item string, list []string) bool {
    for _, b := range list {
        if b == item {
            return true
        }
    }
    return false
}

func tagsMatch(tags []string, resources []string) bool {
	for _, tag := range tags {
	    if !contains(tag, resources) {
		return false
	    }
	}
	return true
}

func GetImageTags(imageId string) []string {
	for _, image := range getImagesList() {
	    if strings.Contains(imageId, image.Imageid) {
		return image.Tags
	    }
	}
	return []string{}
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