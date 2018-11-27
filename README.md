# lander - automate the landing page for your standalone docker host

|Maintainer| David Daehne <david.daehne@msg-david.de>|
|---|---|
|**Version**|[0.4.0](https://operations.gba.msg.team/ao/gitlab/DevOps/lander/tags)|
|**Status**|~~planned~~ -> **in progress** -> evaluation -> ready|
|**written in**|go|
|**docker registry repo**|docker.msg.team/automotive/lander|

## How does it work?
Lander is basically a webserver implementation which pulls information about running containers from the docker daemon every time it gets a http GET request at '/'.
Based on it's configuration, lander will decide which and how to generate links to running containers on this host.

To do so, it needs the following labels at your docker containers:
- lander.enable: if set (e.g. to 'true'), lander will start to parse additional lanbels
- lander.group: indicates the group, under which lander will put the application inside the index.html
- lander.name: gives the name to be shown as a link in the index.html

## How to build lander:
To build lander, you need a up-to-date installation of the [Go](https://golang.org) programming tools (and eventually 'make').

- clone repository:
```
git clone https://operations.gba.msg.team/ao/gitlab/DevOps/lander.git
```

- install dependencies:
```
# via make:
make dep
# directly via go-dep:
dep ensure
```

- build lander:
```
# via make: 
make
# directly via go:
go build
```

- build lander docker image:
```
make image
```
This step will build a docker image with the tag 'local/lander:latest' locally

## Configuration:
Lander receives it's configuration completely via environment variables. At the moment you can set:

- LANDER_DOCKER: absolute path (including protocoll) to the docker daemon socket (e.g. "unix:///var/run/docker.sock" - which is the standard path)
- LANDER_TRAEFIK: if true, lander will search for [Traefik](https://traefik.io/) specific labels and use them to create http application paths. Possible values:
    - true: standard
    - false
- LANDER_EXPOSED: if true, lander will search for exposed container ports. possible values:
    - true
    - false: standard
- LANDER_LISTEN: the ip address and port on which lander will listen for requests (e.g. 192.168.1.1:9000, default: ":8080")
- LANDER_TITLE: string which will deals as headline in the index.html (standard: "LANDER")
- LANDER_HOSTNAME: should be the hostname of the docker host. used to generate hyperlinks. (Attention: lander won't work correctly without this variable set!**)

## State of development:
At the moment, lander can in fact find publicly exposed ports of docker containers, but will assume that the applications answer via http, not https. When searching for hyperlinks via traefik labels, 
lander will assume a https connections, since traefik is mostly used as an ssl endpoint. In addition to that parsing of a traefik hostname parameter is not possible at the moment.

We also have a maintest.go file, but the included tests are neither good nor currently maintained...

## Contribute
You wanna contribute? Well.. the best idea would be to create a new branch, code on there and open up a merge request.
The go dependencies here are maintained with [dep](https://github.com/golang/dep).
