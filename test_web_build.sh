#! /bin/bash

# Build web UI
# 设置好GOPATH变量
export src_next="Streaming-Media-Service"

pwd #结果应该为$GOPATH/src/$src_next/
cd ./web
go install
#mkdir $GOPATH/bin/video_server_web_ui #目录不存在请先创建

mv $GOPATH/bin/web $GOPATH/bin/video_server_web_ui/web
cp -R $GOPATH/src/$src_next/web/templates $GOPATH/bin/video_server_web_ui/
