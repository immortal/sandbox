#!/bin/sh

rm -f pipe
mkfifo -m 0666 pipe

sleep 10000 > pipe
