#!/bin/sh


SCRIPT_DIR=$(cd -- "$(dirname -- "$0")" && pwd -P)
WORK_DIR="${SCRIPT_DIR}/.."


# check if .env exists
if [ ! -f "${WORK_DIR}"/.env ];then
    printf "\033[31m.env file not found\033[0m\n"
    exit 3
fi

# load .env
. "${WORK_DIR}"/.env


migrate -path migrations-dev -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose up