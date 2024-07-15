#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240707233753-b765e5d5218f
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240707233753-b765e5d5218f
gomobile init
