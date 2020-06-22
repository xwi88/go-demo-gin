#!/usr/bin/env bash

# absolute path
# shellcheck disable=SC2046
# shellcheck disable=SC2164
CURRENT_PATH=$(cd $(dirname "$0"); pwd)

echo "CURRENT_PATH: ${CURRENT_PATH}"
make docker-mod -f "${CURRENT_PATH}"/Makefile

echo "docker-compose stop"
docker-compose -f "${CURRENT_PATH}"/docker-compose.yml stop

echo "docker-compose rm"
docker-compose -f "${CURRENT_PATH}"/docker-compose.yml rm -f

echo "docker-compose up"
docker-compose -f "${CURRENT_PATH}"/docker-compose.yml  up -d

docker-compose ps
