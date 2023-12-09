#!/bin/sh

container_name="epseed-db"

if [ "$( docker container inspect -f '{{.State.Running}}' $container_name )" = "true" ]; then
    echo "Importing database..."
    docker compose exec database sh -c 'mariadb -uroot -ple_mdp_de_ouf < dump.sql'
    echo "Import done"
else
    echo "Container epseed-db is not running, please run it first (docker compose up -d database)" >&2
fi
