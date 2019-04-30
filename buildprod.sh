#! /bin/bash
export src_next="Streaming-Media-Service"

cd $GOPATH/src/$src_next/api
#env GOOS=linux GOARCH=amd64 go build -o ../bin/api
go build -o ../bin/api

cd $GOPATH/src/$src_next/scheduler
#env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler
go build -o ../bin/scheduler

cd $GOPATH/src/$src_next/streamserver
#env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver
go build -o ../bin/streamserver

cd $GOPATH/src/$src_next/web
#env GOOS=linux GOARCH=amd64 go build -o ../bin/web
go build -o ../bin/web