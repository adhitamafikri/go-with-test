#!/bin/bash

if [ $# -ne 2 ]; then
    echo "Usage: $0 <module_number> <module_name>"
    echo "Example: $0 011 my_new_module"
    exit 1
fi

MODULE_NUMBER=$1
MODULE_NAME=$2
MODULE_DIR="${MODULE_NUMBER}_${MODULE_NAME}"

if [ -d "$MODULE_DIR" ]; then
    echo "Error: Directory '$MODULE_DIR' already exists"
    exit 1
fi

echo "Creating module: $MODULE_DIR"

mkdir -p "$MODULE_DIR"
cd "$MODULE_DIR" || exit 1

touch "${MODULE_NAME}.go"
touch "${MODULE_NAME}_test.go"

mkdir -p cmd
touch cmd/main.go

go mod init "$MODULE_NAME"

cd ..

if grep -q "\.\/${MODULE_DIR}" go.work; then
    echo "Module already exists in go.work"
else
    awk -v module="./$MODULE_DIR" '/use \(/ { use=1 } use && /^\)/ { print "\t" module; use=0 } 1' go.work > go.work.tmp && mv go.work.tmp go.work
fi

echo "Module $MODULE_DIR created successfully!"