GODEP := $(shell command -v dep 2> /dev/null)

all: lander

lander:
	@echo "++ building lander executable" && \
	go build ./src/lander

run:
	@echo "++ run lander"
	LANDER_DOCKER=unix:///var/run/docker.sock LANDER_LOGLEVEL=debug LANDER_HOSTNAME=`hostname -f` go run ./src/lander/main.go

dep:
ifndef GODEP
	$("!! ERROR: go dep is either not installed or not in $PATH")
endif
	@echo "++ installing project's dependencies" && \
	cd ./src/lander && \
	dep ensure

image:
	@echo "++ buildng docker image local/lander" && \
	docker build -t local/lander ./src/

clean:
	@echo "++ cleaning workspace" && \
	rm -f ./lander
