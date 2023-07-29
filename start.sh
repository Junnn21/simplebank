#!/bin/sh

set -e

echo "run db migration"
. /app/app.env
echo "after source"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"