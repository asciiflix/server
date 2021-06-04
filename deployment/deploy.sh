#!/bin/bash

echo "Cloning Server Repo from GitHub.com..."
git clone git@github.com:asciiflix/server.git
echo "Repo Cloned"

echo "Starting Docker-Compose"
cd ./server

echo "Swichting to Master Branch"
git switch master

echo "Getting current tag"
tag=$(git describe --tags `git rev-list --tags --max-count=1`)

echo "Build Server from source files"
docker-compose build --build-arg VERSION=$tag

echo "Starting builded containers"
docker-compose up -d

echo "Started Containers..."
echo "Cleaning Up Source Files"
rm -rf ~/server

exit 0
