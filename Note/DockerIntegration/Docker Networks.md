- docker network
docker network create my-network
Commands:
  connect     Connect a container to a network
  create      Create a network
  disconnect  Disconnect a container from a network
  inspect     Display detailed information on one or more networks
  ls          List networks
  prune       Remove all unused networks
  rm          Remove one or more networks

  docker network inspect host

  
# Execute the following command to view all networks:
docker network ls

# inspect the network 
docker network inspect afd6dc41d1a9

# Add the container to the network
docker network connect afd6dc41d1a9 <container_id>