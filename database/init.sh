#!/bin/bash
set -e

echo "Initializing the database with partitioned tables..."

# Resolve the directory of this script, whether run locally or in Docker
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$DIR/test_db"

# Execute the SQL script
if [ -n "$MYSQL_ROOT_PASSWORD" ]; then
    mysql -u root -p"$MYSQL_ROOT_PASSWORD" "$MYSQL_DATABASE" < employees_partitioned.sql
else
    echo "MYSQL_ROOT_PASSWORD is not set. You will be prompted for the MySQL root password."
    mysql -u root -p "$MYSQL_DATABASE" < employees_partitioned.sql
fi

echo "Initialization finished successfully."
