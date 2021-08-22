## This application with two services:

###### my docker repository: https://hub.docker.com/repository/docker/akvise/goscratch

* one of them (1. main_server) - server with a small html page to make a POST request (on localhost:8080)
* the other one (2. server_store) - server that stores user tokens from the previos server (on localhost:8081)

(1) set a coockie and make a token for this rule: "name:address". And next step, (1) send token to (2), 
and last of them contains tokens in map and display all tokens in JSON format

I am using Ubuntu 20.04 for this task. also i am using 2 docker files to make 2 containers.
the official golang image on the docker hub weighs many megabytes. That's not cool. Therefore, I use a scratch image.

``` bash
docker pull scratch
```

An important point, since I use scratch, it is necessary to make a static assembly of the application so that it does not search for libraries in the system. 
Since it will be in a container, it will not find them. Check this [article](https://habr.com/ru/post/460535/ "https://habr.com/ru/post/460535/").

instead of doing a dynamic build:
``` bash
go build -o main
```

need to run following command for static build:
``` bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main
```

And after that, we just build a docker image to check the container and push it to our docker hub:
``` bash
docker build -t <new_image_name> -f <dockerfile_name>
```

To correctly push an image, its name must conform to the following format([this link](https://dker.ru/docs/docker-engine/get-started-with-docker/tag-push-pull-your-image/ "https://dker.ru/docs/docker-engine/get-started-with-docker/tag-push-pull-your-image/")):
YOUR_DOCKERHUB_NAME/docker-image

And then just commit and push it like in git

