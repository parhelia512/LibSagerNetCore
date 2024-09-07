#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240905004112-7c4916698cc9
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240905004112-7c4916698cc9
gomobile init
