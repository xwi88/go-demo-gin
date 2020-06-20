#!/usr/bin/env bash

# absolute path
current_path=$(
    # shellcheck disable=SC2164
    cd $(dirname "${BASH_SOURCE[0]}")
    pwd
)

# shellcheck disable=SC1090
source "${current_path}"/docker-version.sh

# shellcheck disable=SC2154
docker push "${LOCATION}"
