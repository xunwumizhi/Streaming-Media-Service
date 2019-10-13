#! /bin/bash
export src_next="Streaming-Media-Service"

cd $GOPATH
export install_path="windowsEXE"
mkdir $install_path
cd $install_path
mkdir videos

cd $GOPATH\src\\$src_next\\api
#env GOOS=linux GOARCH=amd64 go build -o ../bin/api
go build -o $GOPATH$install_path/api.exe

cd $GOPATH\src\\$src_next\\scheduler
#env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler
go build -o $GOPATH$install_path/scheduler.exe

cd $GOPATH\src\\$src_next\\streamserver
#env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver
go build -o $GOPATH$install_path/streamserver.exe

cd $GOPATH\src\\$src_next\\web
#env GOOS=linux GOARCH=amd64 go build -o ../bin/web
go build -o $GOPATH$install_path/web.exe

cp -R templates/ $GOPATH$install_path