#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240604190613-2782386b8afd
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240604190613-2782386b8afd
gomobile init
