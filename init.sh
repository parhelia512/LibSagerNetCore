#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240716161057-1ad2df20a8b6
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240716161057-1ad2df20a8b6
gomobile init
