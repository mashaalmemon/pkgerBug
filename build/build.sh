#!/usr/bin/env bash

# Working directory should be <PROJECT_ROOT>/client/

# run pkger for /libEmbed, including specific files explicitly and outputting to package that will be included
cd ../libEmbed
pkger -include /sub/HelloWorld.txt -o /sub
cd ../executable

# Build executable
go build -o ../build/main main.go

# delete pkged.go file produced for /lib
cd ../libEmbed/sub
rm pkged.go
cd ../../build