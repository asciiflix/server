#!/bin/bash

echo "Starting newman tests..."

newman run ./Server-API.postman_collection.json -e ./Server-API.local.postman_environment.json