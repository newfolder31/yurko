#!/bin/sh

dbname="yurkodb"
rolename="yurkorole" # password 'secret'

echo "Drop DB..."
psql -U postgres -c "DROP DATABASE $dbname;"

echo "Drop role..."
psql -U postgres -c "DROP ROLE $rolename;"

echo "Create role..."
psql -U postgres -c "CREATE ROLE $rolename LOGIN
  ENCRYPTED PASSWORD 'md594a9b8678b087dffd5bf8b3d7cd02117'
  NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;"

echo "Create DB..."
psql -U postgres -c "CREATE DATABASE $dbname OWNER $rolename ENCODING 'UTF-8';"

echo "Init DB..."
psql -U postgres -d $dbname -f init-scheme.sql
psql -U postgres -d $dbname -f init-data.sql
