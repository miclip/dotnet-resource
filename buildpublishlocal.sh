#!/bin/bash

set -ex 

docker build . -t dotnet-resource -t localhost:5000/dotnet-resource:latest
docker push localhost:5000/dotnet-resource:latest