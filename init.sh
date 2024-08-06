#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240806205939-81131f6468ab
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240806205939-81131f6468ab
gomobile init
