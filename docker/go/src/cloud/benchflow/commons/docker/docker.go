package docker

import (
	"github.com/fsouza/go-dockerclient"
	"log"
)

//Function to create a Docker client using the Docker socket (shared with the container)
func CreateDockerClient() docker.Client {
	endpoint := "unix:///var/run/docker.sock"
    client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	return *client
}