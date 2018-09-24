#!/bin/bash

docker-compose up -d

docker container exec -it manager docker network create --driver=overlay --attachable todoapp
docker container exec -it manager docker swarm init
SWMTKN=$(docker container exec -it manager docker swarm join-token -q worker)

for no in 01 02 03
do
    docker container exec -it worker$no docker swarm join --token $SWMTKN "manager:2377"
done
