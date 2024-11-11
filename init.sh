#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20241108191957-fa514ef75a0f
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20241108191957-fa514ef75a0f
gomobile init
