all: lander

lander:
	@echo "++ building lander executable" && \
	go build

run:
	@echo "++ run lander"
	go run main.go

clean:
	@echo "++ cleaning workspace" && \
	rm -f ./lander
