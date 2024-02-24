#!/bin/bash

while true; do
	echo -e "*2\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\nbar\r\n" | nc 0.0.0.0 6379
	echo -e "*2\r\n$3\r\nget\r\n$3\r\nfoo\r\n" | nc 0.0.0.0 6379

	sleep 1
done
