# Getting Started with Docker Using Node:
1. Create Dockerfile in command prompt
echo. > Dockerfile

- Place the above Dockerfile in your project's root directory and execute the following command in the terminal to build the Docker image:

docker build -t golang_backend .

- After the build completes, you can run the following command to start the container:


docker run golang_backend

bash 
docker run -p 8083:8083 -v $PWD:/app --name golang_backend golang:alpine 


- Start the specified images
docker run -it --rm node:20.11.1-slim /bin/bash

docker run -it --rm node:20.11.1-slim Node.js 的交互式终端

