#!/bin/bash

docker container exec -it manager \
    docker service ps todo_mysql_master \
    --no-trunc \
    --filter "desired-state=running" \
    --format "docker container exec -it {{.Node}} docker container exec -it {{.Name}}.{{.ID}} init-data.sh"
