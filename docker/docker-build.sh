#!/usr/bin/env bash

# absolute path
current_path=$(
    # shellcheck disable=SC2046
    cd $(dirname "${BASH_SOURCE[0]}") || exit
    pwd
)

function echo_red(){
    # shellcheck disable=SC2124
    local str_info=$@
    echo -e "\033[31m $str_info \033[0m"
    return 0
}

# shellcheck disable=SC1090
source "${current_path}"/docker-version.sh

ENV=$1
# shellcheck disable=SC2154
echo -e "docker build info: ENV: ${ENV}, Image: ${LOCATION}"
if [[ ${ENV} == "" ]]; then
    echo_red "please specify the env for image: ${LOCATION}"
    exit 0
fi

# shellcheck disable=SC1090
source "${current_path}"/make.sh

# project root dir
parent_dir="${current_path}/.."

# shellcheck disable=SC2164
if [[ ${ENV} == "mod" ]]; then
  docker build -t "${LOCATION}"  -f "${parent_dir}"/Dockerfile.mod .
else
  docker build -t "${LOCATION}"  -f "${parent_dir}"/Dockerfile .
fi
