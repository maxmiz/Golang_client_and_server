#!/bin/bash


xterm -e "~/Documents/Golang_client_and_server/Teste1/server_run.sh; bash" & 
sleep 2
xterm -e "~/Documents/Golang_client_and_server/Teste1/client_run.sh; bash"
