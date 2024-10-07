#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20241004191011-08a83c5af9f8
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20241004191011-08a83c5af9f8
gomobile init
