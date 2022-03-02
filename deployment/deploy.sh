#!/bin/bash

echo "Cloning Server Repo from GitHub.com..."
git clone git@github.com:asciiflix/server.git

echo "Changing Directory to Server"
cd ./server

echo "Swichting to Master-Branch"
git switch master

echo "Starting builded containers"
VERSION=$1 docker-compose up -d

echo "Cleaning Up Source Files"
rm -rf ~/server

exit 0
