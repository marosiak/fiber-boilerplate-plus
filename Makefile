project_name = weekly
image_name = gofiber:latest

air:
	# May not work if you don't have $GOPATH/bin in $PATH, we don't want to touch your envs
	go install github.com/cosmtrek/air@latest && air

run-local:
	go run cmd/main.go

requirements:
	go mod tidy

clean-packages:
	go clean -modcache

up: 
	make up-silent
	make shell

build:
	docker build -t $(image_name) .

build-no-cache:
	docker build --no-cache -t $(image_name) .

up-silent:
	make delete-container-if-exist
	docker run -d -p 8000:8000 --name $(project_name) $(image_name)

delete-container-if-exist:
	docker stop $(project_name) || true && docker rm $(project_name) || true

shell:
	docker exec -it $(project_name) /bin/sh

stop:
	docker stop $(project_name)

start:
	docker start $(project_name)
