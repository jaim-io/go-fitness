#!/bin/sh

# Abort ony any error (including if wait-for-it.sh fails)
set -e

if [ -n "$POSTGRES_HOST" ]; then
  for i in {1..5}; do 
    /app/wait-for-it.sh "$POSTGRES_HOST:${POSTGRES_PORT:-5432}" && break || echo "docker-entrypoint.sh: sleeping 15 seconds" && sleep 15;
  done
fi

exec "$@"