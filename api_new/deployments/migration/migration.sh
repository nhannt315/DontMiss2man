#!/bin/bash

BASE_DIR=$(cd $(dirname $0) && pwd)
GOOSE_BIN=${BASE_DIR}/../bin/goose
MIGRATION_DIR=${BASE_DIR}/migrations
DB_DRIVER=mysql

DB_USERNAME=${DB_USERNAME:-root}
DB_PASSWORD=${DB_PASSWORD:-''}
DB_HOST=${DB_HOST:-127.0.0.1}
DB_PORT=${DB_PORT:-3306}
DB_NAME=${DB_NAME:-real_estate}

SUB_CMD=$1
MIGRATION_NAME=$2

# parseTime=true は、gooseのために必要
DB_STRING="${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true"

FLAGS="-dir ${MIGRATION_DIR}"

if [ "$DB_CA_CERT_PATH" != "" ]; then
  FLAGS="${FLAGS} -certfile ${DB_CA_CERT_PATH}"
fi

function create_new_migration_file() {
    mkdir -p ${MIGRATION_DIR}
    local latest_file=`ls ${MIGRATION_DIR}/*.sql | sort | tail -1`
    local next_seq
    if [[ -z ${latest_file} ]] ; then
        next_seq=1
    elif [[ ${latest_file} =~ ([0-9]+)_[A-Za-z0-9]+(_[A-Za-z0-9]+)*.sql ]] ; then
        next_seq=$(( 10#${BASH_REMATCH[1]} +1))
    else
        echo "invalid migrations. ${latest_file}"
        exit 1
    fi
    local migration_name
    if test -z "$MIGRATION_NAME"
    then
      migration_name="migration"
    else
      migration_name=$MIGRATION_NAME
    fi
    local new_file=`printf %04d ${next_seq}`_${migration_name}.sql
    echo "create new migration file.  ${MIGRATION_DIR}/${new_file}"
    cat <<EOF >${MIGRATION_DIR}/${new_file}
-- +goose Up
-- SQL in this section is executed when the migration is applied.
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
EOF
}

case ${SUB_CMD} in
  "create")
    create_new_migration_file
    ;;
  "up")
    ${GOOSE_BIN} $FLAGS ${DB_DRIVER} "${DB_STRING}" up
    ;;
  "down")
    ${GOOSE_BIN} $FLAGS ${DB_DRIVER} "${DB_STRING}" down
    ;;
  "status")
    ${GOOSE_BIN} $FLAGS ${DB_DRIVER} "${DB_STRING}" status
    ;;
  *)
    echo "invalid sub command. ${SUB_CMD}"
    ${GOOSE_BIN} --version
    exit 1
    ;;
esac
