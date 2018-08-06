GODEP := $(shell command -v dep 2> /dev/null)

all: lander

lander:
	@echo "++ building lander executable" && \
	go build

run:
	@echo "++ run lander"
	LANDER_DOCKER=unix:///var/run/docker.sock go run main.go

dep:
ifndef GODEP
	$("!! ERROR: go dep is either not installed or not in $PATH")
endif
	@echo "++ installing project's dependencies" && \
	dep ensure

clean:
	@echo "++ cleaning workspace" && \
	rm -f ./lander
