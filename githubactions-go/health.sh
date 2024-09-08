#!/bin/bash

echo "Testing code ..."
go test -v
sleep 1

echo "Testing endpoints on port 5000 ....."
sleep 1

curl http://13.53.35.206:5000/photos
sleep 1

curl http://13.53.35.206:5000/users
sleep 1

curl http://13.53.35.206:5000/posts
sleep 1

echo -e "\nAll endpoints tested"
