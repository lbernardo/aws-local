#!/bin/sh
make build

P_PORT=""
P_HOST=""

if [ $PORT != "" ] ; then
  P_PORT="--port $PORT"
fi

if [ $HOST != "" ] ; then
  P_HOST="--host $HOST"
fi

awslocal s3 $P_PORT $P_HOST --volume $VOLUME