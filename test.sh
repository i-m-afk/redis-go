#!/bin/bash

while true; do
	echo -e "*2\r\n$4\r\necho\r\n$3\r\nhey\r\n" | nc 0.0.0.0 6379
	sleep 1
done
