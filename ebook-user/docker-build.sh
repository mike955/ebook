#!bin/bash

 
set -ex
 
IMAGE_NAME="ebook/ebook-user"
VERSION=0.0.0.1
docker build -t $IMAGE_NAME:$VERSION -f Dockerfile .