# k8s-in-action-example

A "Hello, World!" style golang application that exposes `localhost:8080/hello` for use while working through *Kubernetes in Action*. A Dockerfile is included to build a docker image with the application binary.

## Run the app

```
go run cmd/k8s-in-action-example/main.go
```

## Build the docker image

```
docker build -t percygrunwald/k8s-in-action-example .
```

## Run the application with docker

```
docker run -p 8080:8080 percygrunwald/k8s-in-action-example

curl locahost:8080
# => 404 page not found

curl localhost 8080/hello
# => Hello, World!
```

## Push the application to Docker Hub

```
docker login
docker push percygrunwald/k8s-in-action-example
```
