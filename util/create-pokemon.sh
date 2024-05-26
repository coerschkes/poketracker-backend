#!/bin/bash
#-H "Content-Type: application/json"
if [ "$#" -eq 1 ]; then
	curl -v -X POST -H "Authorization: Bearer $1" -H "Content-Type: application/json" --data '{"dex":5,"name":"test", "types":["Wasser", "Feuer"], "shiny":true, "normal":true, "universal":true, "regional":true, "editions":["Smaragd"]}' http://localhost:1323/api/pokemon
fi
exit 1

#		Dex       int      `json:"dex" validate:"required"`
 #		Name      string   `json:"name" validate:"required"`
 #		Types     []string `json:"types" validate:"required"`
 #		Shiny     bool     `json:"shiny" validate:"required"`
 #		Normal    bool     `json:"normal" validate:"required"`
 #		Universal bool     `json:"universal" validate:"required"`
 #		Regional  bool     `json:"regional" validate:"required"`
 #		Editions  []string `json:"editions" validate:"required"`
 #		UserId    int      `json:"userId" validate:"required"`