#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240326195318-268e6c3a80d1
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240326195318-268e6c3a80d1
gomobile init
