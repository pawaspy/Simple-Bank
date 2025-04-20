#!/bin/sh

set -e

echo "Waiting for postgres..."
sleep 3

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up


echo "start the app"
exec "$@"