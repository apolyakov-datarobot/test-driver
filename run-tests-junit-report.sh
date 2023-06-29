#!/bin/sh


go test -v -run $1 /app/*.go | go-junit-report -set-exit-code
