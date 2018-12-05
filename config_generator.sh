#!/usr/bin/env bash
#/bin/bash

echo '
application:
  host: "$APPLICATION_HOST"
  http:
    enabled: $APPLICATION_HTTP_ENABLED
    port: $APPLICATION_HTTP_PORT
  grpc:
    enabled: $APPLICATION_GRPC_ENABLED
    port: $APPLICATION_GRPC_PORT

database:
  address: "$DATABASE_HOST:$DATABASE_PORT"
  user: "$DATABASE_USER"
  password: "$DATABASE_PASSWORD"
  database: "DATABASE_NAME"

consul:
  host: "$CONSUL_HOST"
  port: $CONSUL_PORT
'
