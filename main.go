package main

import (
	"io"
	/*"os"*/
	"log"
	"net/http"
)

type Config struct {
	traefik bool
	exposed bool
	port    string
}

func RenderAndRespond(w http.ResponseWriter, r *http.Request) {
	/*conf := getConfig()*/
	log.Println(r.RemoteAddr, r.Method, r.URL)
	io.WriteString(w, "Hallo Welt!")
}

func getConfig() Config {
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
	config := getConfig()

	// register handle function for root context
	http.HandleFunc("/", RenderAndRespond)

	log.Println("Starting Server on", config.port)
	err := http.ListenAndServe(config.port, nil)
	if err != nil {
		panic(err)
	}

}
