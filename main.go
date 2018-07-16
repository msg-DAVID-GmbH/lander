package main

import (
	"github.com/fsouza/go-dockerclient"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Traefik string
	Exposed string
	Port    string
	Title   string
}

type Container struct {
	AppName string
	AppURL  string
}

type PayloadData struct {
	Title  string
	Groups map[string][]Container
}

func (payload PayloadData) Get() {
	// set endpoint of docker daemon api
	// TODO: parameterize docker endpoint URL
	endpoint := "unix:///var/run/docker.sock"

	// get new client
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	// get running containers
	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		panic(err)
	}

	// iterate through slice of containers and find "lander" labels
	for _, container := range containers {

		// check if map contains a key named "lander.enable"
		if _, found := container.Labels["lander.enable"]; found {
			log.Println("INFO: found lander labels on Container:", container.ID)

			// extract strings for easier use
			ContainerName := container.Labels["lander.name"]
			delimiterPosition := strings.LastIndex(container.Labels["traefik.frontend.rule"], ":")
			ContainerURL := container.Labels["traefik.frontend.rule"][delimiterPosition:]

			// check if lander.group is already present
			if _, found := payload.Groups[container.Labels["lander.group"]]; found {
				payload.Groups[container.Labels["lander.group"]] = append(payload.Groups[container.Labels["lander.group"]], Container{AppName: ContainerName, AppURL: ContainerURL})
			} else {
				payload.Groups[container.Labels["lander.group"]] = []Container{Container{AppName: ContainerName, AppURL: ContainerURL}}
			}
		}
	}
}

func RenderAndRespond(w http.ResponseWriter, r *http.Request) {
	// check if the request is exactly "/", otherwise stop the response
	if r.URL.String() != "/" {
		log.Println("ERROR:", r.RemoteAddr, r.URL, "not a valid request")
		return
	}

	// print request to log
	log.Println("INFO:", r.RemoteAddr, r.Method, r.URL)

	// initialize payload struct
	var payload = PayloadData{"", make(map[string][]Container)}
	// call method to get values
	payload.Get()
	// set page title
	// TODO: make this variable
	payload.Title = "devops01.david-bs.de"

	templ := template.Must(template.ParseFiles("template.html"))

	err := templ.Execute(w, payload)
	if err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	var config Config

	config.Traefik = os.Getenv("LANDER_TRAEFIK")
	if config.Traefik == "" {
		log.Println("INFO: environment variable LANDER_TRAEFIK not set, assuming: \"true\"")
		config.Traefik = "true"
	}

	config.Exposed = os.Getenv("LANDER_EXPOSED")
	if config.Exposed == "" {
		log.Println("INFO: environment variable LANDER_EXPOSED not set, assuming: \"false\"")
		config.Exposed = "false"
	}

	config.Port = os.Getenv("LANDER_PORT")
	if config.Port == "" {
		log.Println("INFO: environment variable LANDER_PORT not set, assuming: \"8080\"")
		config.Port = ":8080"
	}

	config.Title = os.Getenv("LANDER_TITLE")
	if config.Title == "" {
		log.Println("INFO: environment variable LANDER_TITLE not set, assuming: \"LANDER\"")
		config.Title = "LANDER"
	}

	return config
}

func main() {
	// get configuration
	config := GetConfig()

	// register handle function for root context
	http.HandleFunc("/", RenderAndRespond)

	// start listener
	log.Println("INFO: Starting Server on", config.Port)
	err := http.ListenAndServe(config.Port, nil)
	if err != nil {
		panic(err)
	}

}
