#!/bin/bash

if [[ "$TRAVIS_BRANCH" != "master" ]]; then
  echo "We're not on the master branch."
  exit 0
fi

if [[ "$TRAVIS_TAG" == "" ]]; then
  echo "We're not in a release."
  exit 0
fi

export REPO=gocaio/goca
export TAG=$TRAVIS_TAG
export COMMIT=${TRAVIS_COMMIT::8}

docker login -e $DOCKER_EMAIL -u gocaio -p $DOCKER_PASS
docker build -f Dockerfile -t $REPO:$COMMIT .
docker tag $REPO:$COMMIT $REPO:latest
docker tag $REPO:$COMMIT $REPO:$TAG
docker push $REPO
