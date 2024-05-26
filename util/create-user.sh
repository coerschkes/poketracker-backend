#!/bin/bash
#-H "Content-Type: application/json"
if [ "$#" -eq 1 ]; then
	curl -v -X POST -H "Authorization: Bearer $1" -H "Content-Type: application/json" --data '{"email":"test@test2.com","firebaseUid":"whatever"}' http://localhost:1323/api/user
fi
exit 1