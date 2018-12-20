#!/usr/bin/env bash

export GOPATH=$PWD
go build ./...
$GOPATH/cmd/backendRaftD/backendRaftD -node localhost:46000 -id node0 ~/node0 &
sleep 3
$GOPATH/cmd/backendRaftD/backendRaftD -id node1 -grpc 45001 -node localhost:46001 -leader localhost:45000 ~/node1 &
$GOPATH/cmd/backendRaftD/backendRaftD -id node2 -grpc 45002 -node localhost:46002 -leader localhost:45000 ~/node2 &
$GOPATH/cmd/backendCntD/backendCntD &
$GOPATH/cmd/backendUsrD/backendUsrD &
cd $GOPATH/cmd/web/
sudo ./web
