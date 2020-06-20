#!/usr/bin/env bash

set -e
set -o pipefail

[[ $DEBUG == true ]] && set -x

APP_TYPE="go"
APP_PATH="/data1/services/app"
APP_FULL_PATH="_________"

# shellcheck disable=SC2034
OPERATE=$1
# shellcheck disable=SC2034
RUNNING="false"

if [[ ${APP_PATH:0:1} != "/" ]]; then
	APP_PATH="$(pwd)/${APP_PATH}"
fi

mkdir -p "${APP_PATH}"/log
APP_NAME=${APP_PATH##*/}
if [[ ${APP_TYPE} == "java" ]]; then
	APP_FULL_PATH=${APP_PATH}/lib/${APP_NAME}
elif [[ ${APP_TYPE} == "go" ]]; then
	APP_FULL_PATH=${APP_PATH}/${APP_NAME}
fi


case ${1} in
  app:start|app:version|app:status|app:stop)
    # shellcheck disable=SC2005
    echo "$(date "+%FT%T%z"), command:[${1}]"

    case ${1} in
      app:start)
        exec "${APP_FULL_PATH}" start -c conf/app.yml
        ;;
      app:version)
        exec "${APP_FULL_PATH}" version
        ;;
      app:status)
        # shellcheck disable=SC2005
        echo "$(ps -ef | grep "${APP_FULL_PATH}" | grep -v grep | grep -v kill)"
        ;;
      app:stop)
        APP_PID=$(ps -ef | grep -n "${APP_FULL_PATH}" | grep -v grep | grep -v kill | awk '{print $2}')
        kill -15 "${APP_PID}"
        if ps h -p "${APP_PID}" > /dev/null ; then
        wait "${APP_PID}" || true
        fi
        echo "app stop success"
        ;;
    esac
    ;;
  app:help)
    echo "Available options:"
    echo " app:start        - Starts the app service (default)"
    echo " app:version      - Display the app version."
    echo " app:status       - Display the app status."
    echo " app:stop         - Execute a rake task."
    echo " app:help         - Displays the help"
    echo " [command]        - Execute the specified command, eg. bash."
    ;;
  *)
    exec "$@"
    ;;
esac
