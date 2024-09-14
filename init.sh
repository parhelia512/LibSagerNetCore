#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240909163608-642950227fb3
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240909163608-642950227fb3
gomobile init
