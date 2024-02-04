DOCKER_APP_NAME=twitch-slower

run:
	docker run -it $(DOCKER_APP_NAME):latest "$@"

build:
	docker build -t $(DOCKER_APP_NAME):latest .