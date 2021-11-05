#!/bin/sh

PROJECT_DIR=/go/src/app

SCRIPT_DIR=$(cd -- "$(dirname -- "$0")" && pwd -P)
WORK_DIR="${SCRIPT_DIR}/.."


# check if .env exists
if [ ! -f "${WORK_DIR}"/.env ];then
    printf "\033[31m.env file not found\033[0m\n"
    exit 3
fi

# load .env
. "${WORK_DIR}"/.env

echo "================================================================================"
printf "\033[36mВыполнение команд в контейнере\033[0m\n"
echo "================================================================================"
printf "\033[36mУстановка Go migrate(https://github.com/golang-migrate/migrate)\033[0m\n"
echo "================================================================================"
apt-get install -y lsb-release && apt-get clean all
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
apt-get install -y migrate
