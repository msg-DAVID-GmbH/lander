package main

import (
	"io"
	/*"os"*/
	"github.com/fsouza/go-dockerclient"
	"log"
	"net/http"
)

type Config struct {
	traefik bool
	exposed bool
	port    string
}

type Container struct {
	appName  string
	appURL   string
	appGroup string
}

func getRoutes() []Container {
	// set endpoint of docker daemon api
	endpoint := "unix:///var/run/docker.sock"

	// get new client
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	// initialise containers struct slice
	var landerRoutes []Container

	// get running containers
	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		panic(err)
	}
	// iterate through slice of containers and find "lander" labels
	for _, container := range containers {
		if _, found := container.Labels["lander.enable"]; found {
			log.Println("found lander labels on Container:", container.ID)
			landerRoutes = append(landerRoutes, Container{appName: container.Labels["lander.name"], appURL: container.Labels["traefik.frontend.rule"], appGroup: container.Labels["lander.group"]})
		}
	}
	return landerRoutes
}

func RenderAndRespond(w http.ResponseWriter, r *http.Request) {
	// check if the request is exactly "/", otherwise stop the response
	if r.URL.String() != "/" {
		log.Println(r.RemoteAddr, r.URL, "not a valid request")
		return
	}

	// print request to log
	log.Println(r.RemoteAddr, r.Method, r.URL)
	// call getRoutes
	routes := getRoutes()
	log.Println(routes)
	// answer the request
	io.WriteString(w, "Hallo Welt!")
}

func GetConfig() Config {
	// get command line arguments
	// commandLineArgs := os.Args[1:]

	// parse command line arguments
	// for _, param := range commandLineArgs {
	// 	switch param {
	// 	case "traefik=true":
	// 		conf.traefik = true
	// 	case "exposed=true":
	// 		config.exposed = true
	// 	default:

	// 		break
	// 	}
	// }
	// return struct containing configuration parameter

	return Config{traefik: true, exposed: false, port: ":8080"}
}

func main() {
	// get configuration
	config := GetConfig()

	// register handle function for root context
	http.HandleFunc("/", RenderAndRespond)

	log.Println("Starting Server on", config.port)
	err := http.ListenAndServe(config.port, nil)
	if err != nil {
		panic(err)
	}

}
