#!/bin/sh
dockerize -wait tcp://postgres:5432 -timeout 300s
# Run migrations first
./migrate
# Then start the main service
./services
