#!/bin/bash

if [ "$#" -eq 1 ]; then
	curl -v -H "Authorization: Bearer $@" http://localhost:1323/api/pokemon
fi
exit 1
