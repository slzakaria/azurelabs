#!/bin/bash

echo "Testing code ..."
go test -v
sleep 1

read -rp "Enter your ip: " IP
echo "Testing endpoints on port $IP ....."
sleep 1

curl http://$IP:5000/photos
sleep 1

curl http://$IP:5000/users
sleep 1

curl http://$IP:5000/posts
sleep 1

echo -e "\nAll endpoints tested"
