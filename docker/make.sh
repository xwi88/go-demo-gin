#!/usr/bin/env bash

# absolute path
current_path=$(
    # shellcheck disable=SC2046
    cd $(dirname "${BASH_SOURCE[0]}") || exit
    pwd
)

script_file=$(basename "$0")
echo "current script location: $current_path"
echo "current script filename: $script_file"

#设置运行环境
env_support_array=("dev" "uat" "prod" "mod")

ENV=$1
DEFAULT_ENV="dev"
# shellcheck disable=SC2164
# shellcheck disable=SC2006
current_path=$(cd `dirname "$0"`; pwd)

# echo -e "choose env, you want to use: dev|uat|production default(dev)"
# read -p "Please input a env:" INPUT_ENV
if [[ ${ENV} == "" ]]; then
    ENV=${DEFAULT_ENV}
fi

validCount=0
# shellcheck disable=SC2068
for element in ${env_support_array[@]}
do
if [[ "${element}" ==  "${ENV}" ]]; then
    ((validCount++))
fi
done

if [[ ${validCount} == "0" ]]; then
    # shellcheck disable=SC2145
    echo "env输入错误: ${ENV} not exist in ${env_support_array[@]}"
    exit 1
fi

echo -e "choose env: ${ENV}"

parent_dir="${current_path}/.."

conf_dir="${parent_dir}/profiles/${ENV}"
if [[ ! -d "${conf_dir}" ]]; then
	echo "conf_dir 不存在"
  exit 1
fi

rm -rf "${parent_dir}"/dist/*
mkdir -p "${parent_dir}"/dist/{conf,log}

if [[ "$(ls -A "${conf_dir}")" = "" ]]; then
  echo "${conf_dir} 不存在配置文件"
  exit 1
else
  cp -r "${parent_dir}"/profiles/${ENV}/* "${parent_dir}"/dist/conf
  cp "${parent_dir}"/entrypoint.sh "${parent_dir}"/dist
  chmod +x "${parent_dir}"/entrypoint.sh
fi
