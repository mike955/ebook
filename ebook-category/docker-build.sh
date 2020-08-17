set -ex
 
IMAGE_NAME="ebook/ebook-category"
VERSION=0.0.0.1
docker build -t $IMAGE_NAME:$VERSION -f Dockerfile .