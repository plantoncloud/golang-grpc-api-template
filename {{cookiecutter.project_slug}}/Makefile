.PHONY: build
build:
	GOOS=linux go build -o build/app-linux .
	go build -o build/app .

.PHONY: build-image
build-image: build
	docker build --platform linux/amd64 -t {{cookiecutter.project_slug}} .

.PHONY: run
run:
	go run main.go

.PHONY: run-image
run-image:
	docker run --platform linux/amd64 -p 8080:8080  {{cookiecutter.project_slug}}
