#!/bin/bash

if [ "$#" -eq 2 ]; then
	curl -v -X DELETE -H "Authorization: Bearer $1" http://localhost:1323/api/pokemon/$2
fi
exit 1
