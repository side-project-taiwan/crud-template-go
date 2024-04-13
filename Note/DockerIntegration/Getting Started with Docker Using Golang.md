

# Getting Started with Docker Using Node:
1. Create Dockerfile in command prompt
echo. > Dockerfile

- Place the above Dockerfile in your project's root directory and execute the following command in the terminal to build the Docker image:
docker build -t golang_image .


- ----------------option:----------------
# A: Use docker-compose directly
docker-compose up -d

docker-compose -p node-container-group up -d



# B: Use Dockerfile Cli


- After the build completes, you can run the following command to start the container:
bash 
docker run -p 8083:8083 -v $PWD:/app --name golangContainer golang_image

docker run -p 8083:8083 -v $PWD:/app --name golangContainer golang_image



docker run golang_image
- Start the specified images
docker run -it --rm golang:alpine /bin/sh


