#!/bin/bash

echo -n "Cleaning generated templ files..."
find . -type f -name '*_templ.go' -delete 
echo "Done"
echo "Generating templ files..."
templ generate pkg/ui
echo "Running server..."
go run cmd/server/main.go

