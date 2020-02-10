#!/bin/sh
P_PORT=""
P_HOST=""

if [ $PORT != "" ] ; then
  P_PORT="--port $PORT"
fi

if [ $HOST != "" ] ; then
  P_HOST="--host $HOST"
fi

awslocal ssm $P_PORT $P_HOST --values $VALUES