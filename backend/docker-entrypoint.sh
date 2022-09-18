#!/bin/sh

# Abort ony any error (including if wait-for-it.sh fails)
set -e

if [ -n "$MYSQL_HOST" ]; then
  for i in {1..5}; do 
    /app/wait-for-it.sh "$MYSQL_HOST:${MYSQL_PORT:-3306}" && break || echo "docker-entrypoint.sh: sleeping 15 seconds" && sleep 15;
  done
fi

exec "$@"