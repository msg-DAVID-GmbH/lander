package main

import (
	"github.com/fsouza/go-dockerclient"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Config struct {
	traefik bool
	exposed bool
	port    string
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

	// set page title
	// TODO: make this variable
	payload.Title = "devops01.david-bs.de"

	// iterate through slice of containers and find "lander" labels
	for _, container := range containers {

		// check if map contains a key named "lander.enable"
		if _, found := container.Labels["lander.enable"]; found {
			log.Println("found lander labels on Container:", container.ID)

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
		log.Println(r.RemoteAddr, r.URL, "not a valid request")
		return
	}

	// print request to log
	log.Println(r.RemoteAddr, r.Method, r.URL)

	// initialize payload struct
	var payload = PayloadData{"", make(map[string][]Container)}
	// call method to get values
	payload.Get()

	templ := template.Must(template.ParseFiles("index.html"))

	err := templ.Execute(w, payload)
	if err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	// TODO: implement routine to parse command line arguments
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
