#!/bin/bash

while true; do
	echo -e "*3\r\n$4\r\nconfig\r\n$3\r\nget\r\n$3\r\ndir\r\n" | nc 0.0.0.0 6379

	sleep 1
done
