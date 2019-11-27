#!/bin/bash


xterm -e "~/Documents/Golang_client_and_server/Teste1/server/server_run.sh; bash" & 
sleep 2
xterm -e "~/Documents/Golang_client_and_server/Teste1/client/client_run.sh; bash"
