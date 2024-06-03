#!/bin/sh
wait_database()
{
  HOST=$1
  PORT=$2

  echo "Waiting for database..."
  echo "$HOST:$PORT"
  while ! nc -z $HOST $PORT; do
    sleep 1
  done

  echo "database started"
}

wait_database $REDIS_HOST $REDIS_PORT

./exchange_service -config=./deploy/env/.env.prod