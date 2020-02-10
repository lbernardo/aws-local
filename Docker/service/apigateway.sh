#!/bin/sh
make build

P_PORT=""
P_HOST=""
P_NET=""
P_ENV=""
P_SERVERLESS=""

if [ $PORT != "" ] ; then
  P_PORT="--port $PORT"
fi

if [ $HOST != "" ] ; then
  P_HOST="--host $HOST"
fi

if [ $NETWORK != "" ] ; then
  P_NET="--network $NETWORK"
fi

if [ $ENVFILE != "" ] ; then
  P_ENV="--env $ENVFILE"
fi

if [ $SERVERLESS != "" ] ; then
  P_SERVERLESS="--yaml $SERVERLESS"
fi

awslocal api-gateway $P_PORT $P_HOST --volume $VOLUME $P_NET $P_ENV $P_SERVERLESS