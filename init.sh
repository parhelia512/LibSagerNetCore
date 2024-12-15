#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20241213221354-a87c1cf6cf46
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20241213221354-a87c1cf6cf46
gomobile init
