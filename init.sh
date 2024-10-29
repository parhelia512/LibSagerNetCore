#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20241016134751-7ff83004ec2c
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20241016134751-7ff83004ec2c
gomobile init
