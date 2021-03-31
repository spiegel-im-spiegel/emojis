#!/bin/bash
go run main.go > ../emoji-data.json
go run main.go -s > ../emoji-sequences.json
