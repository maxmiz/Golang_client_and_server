#!/bin/bash


xterm -e "/home/rafael/Documents/Golang_client_and_server/server/server_run.sh; bash" & 
sleep 2
xterm -e "/home/rafael/Documents/Golang_client_and_server/client/client_run.sh; bash"
