#!/bin/bash

local_ip=$(hostname -I | awk '{print $1}')
server_id=$(echo "$local_ip" | awk -F'.' '{printf("%d%d%d\n", $2, $3, $4)}')

echo "Server ID: $server_id"
