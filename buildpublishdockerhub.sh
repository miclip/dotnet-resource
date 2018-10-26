#!/bin/bash

set -ex 

docker build . -t dotnet-resource -t miclip/dotnet-resource:latest
docker push miclip/dotnet-resource:latest