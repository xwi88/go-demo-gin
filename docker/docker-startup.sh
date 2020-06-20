#!/usr/bin/env bash


# absolute path
current_path=$(
    # shellcheck disable=SC2164
    cd $(dirname "${BASH_SOURCE[0]}")
    pwd
)

script_file=$(basename "$0")
echo "current script location: $current_path"
echo "current script filename: $script_file"

# shellcheck disable=SC1090
source "${current_path}"/docker-version.sh
# shellcheck disable=SC2154
echo "docker tag:${LOCATION}"

CMD=$*

echo "docker run with command:${CMD}"

docker run -it --name ${CONTAINER_NAME} --rm ${DAEMON} \
-v "${current_path}"/log:/data1/services/app/log ${PORT} \
"${LOCATION}" $CMD
