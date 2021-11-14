#!/bin/sh
# wait_service_up.sh  Dockerサービス名  grep文字列
COMPOSE_OPTS=""
SERVICE=""
TEXT=""
PRE_ARGS="none"
SLEEP=1
END_SLEEP=2
LOG_TAIL_LINES=100

for X in "$@" ; do
    case ${PRE_ARGS} in
    --compose-file)
        COMPOSE_OPTS+=" -f $X"
    ;;
    --service)
        SERVICE=$X
    ;;
    --text)
        TEXT=$X
    ;;
    --tail)
        LOG_TAIL_LINES=$X
    ;;
    esac
    PRE_ARGS=$X
done

echo "${SERVICE} wait : ${TEXT}"
RET=1
while [ ${RET} -ne 0 ] ; do
  sleep ${SLEEP}
  docker-compose $COMPOSE_OPTS ps ${SERVICE} | grep "Exit" >/dev/null
  if [ $? -eq 0 ] ; then
    docker-compose $COMPOSE_OPTS logs --tail ${LOG_TAIL_LINES} ${SERVICE}
    echo "Exit service : ${SERVICE}"
    exit 1
  fi
  docker-compose $COMPOSE_OPTS logs --tail ${LOG_TAIL_LINES} ${SERVICE} | grep "${TEXT}" >/dev/null
  RET=$?
done
sleep ${END_SLEEP}

# 最後にも起動されているか確認
docker-compose $COMPOSE_OPTS ps ${SERVICE} | grep "Exit" >/dev/null
if [ $? -eq 0 ] ; then
  docker-compose $COMPOSE_OPTS logs ${SERVICE}  | tail -10
  echo "Exit service : ${SERVICE}"
  exit 1
fi
