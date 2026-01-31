#!/bin/bash

if [ $# -ne 2 ]; then
    echo "Usage: $0 <module_number> <module_name>"
    echo "Example: $0 011 my_new_module"
    exit 1
fi

MODULE_NUMBER=$1
MODULE_NAME=$2
MODULE_DIR="${MODULE_NUMBER}_${MODULE_NAME}"

to_pascal_case() {
    local name=$1
    local result=""
    local IFS='_'
    read -ra parts <<< "$name"
    for part in "${parts[@]}"; do
        if [ -n "$part" ]; then
            result="${result}$(echo "${part:0:1}" | tr '[:lower:]' '[:upper:]')${part:1}"
        fi
    done
    echo "$result"
}

if [ -d "$MODULE_DIR" ]; then
    echo "Error: Directory '$MODULE_DIR' already exists"
    exit 1
fi

echo "Creating module: $MODULE_DIR"

mkdir -p "$MODULE_DIR"
cd "$MODULE_DIR" || exit 1

cat <<'EOF' > "${MODULE_NAME}.go"
package __MODULE_NAME__

func __PASCAL_CASE__() string {
	return "__MODULE_NAME__"
}
EOF

sed -i '' "s/__MODULE_NAME__/${MODULE_NAME}/g" "${MODULE_NAME}.go"
PASCAL_CASE_NAME=$(to_pascal_case "$MODULE_NAME")
sed -i '' "s/__PASCAL_CASE__/${PASCAL_CASE_NAME}/g" "${MODULE_NAME}.go"

cat <<'EOF' > "${MODULE_NAME}_test.go"
package __MODULE_NAME__

import (
	"testing"
)

func Test__PASCAL_CASE__(t *testing.T) {
	result := __PASCAL_CASE__()
	expected := "__MODULE_NAME__"
	
	if result != expected {
		t.Errorf("Expected: %q, got: %q", expected, result)
	}
}
EOF

sed -i '' "s/__MODULE_NAME__/${MODULE_NAME}/g" "${MODULE_NAME}_test.go"
sed -i '' "s/__PASCAL_CASE__/${PASCAL_CASE_NAME}/g" "${MODULE_NAME}_test.go"

mkdir -p cmd
cat <<'EOF' > cmd/main.go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, __MODULE_NAME__!\n")
}
EOF
sed -i '' "s/__MODULE_NAME__/${MODULE_NAME}/g" cmd/main.go

go mod init "$MODULE_NAME"

cd ..

if grep -q "\.\/${MODULE_DIR}" go.work; then
    echo "Module already exists in go.work"
else
    awk -v module="./$MODULE_DIR" '/use \(/ { use=1 } use && /^\)/ { print "\t" module; use=0 } 1' go.work > go.work.tmp && mv go.work.tmp go.work
fi

echo "Adding make target to Makefile..."
cat >> Makefile <<EOF

.PHONY: ${MODULE_NAME}
${MODULE_NAME}:
	@go version
	@cd ${MODULE_DIR} && go test -v -bench=. -benchmem -cover
EOF

echo "Module $MODULE_DIR created successfully!"