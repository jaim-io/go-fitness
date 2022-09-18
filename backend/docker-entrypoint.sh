#!/bin/sh

# Abort ony any error (including if wait-for-it.sh fails)
set -e

if [ -n "$MYSQL_HOST" ]; then
  /app/wait-for-it.sh "$MYSQL_HOST:${MYSQL_PORT:-3306}"
fi

exec "$@"