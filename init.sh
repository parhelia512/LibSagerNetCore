#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240404231514-09dbf07665ed
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240404231514-09dbf07665ed
gomobile init
