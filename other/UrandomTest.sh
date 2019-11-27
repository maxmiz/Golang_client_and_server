#!/bin/bash

seed=`head -200 /dev/urandom | cksum | cut -f1 -d " "`
echo $seed