#!/bin/bash
#seed=`head -200 /dev/urandom | cksum | cut -f1 -d " "`
#echo $seed

go run client.go -nome Bob -ip 127.0.0.1 -port 8082
