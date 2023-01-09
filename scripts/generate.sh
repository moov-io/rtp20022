#!/bin/bash
set -e

files=($(find ./xsd -name '*.xsd'))

for file in "${files[@]}"
do
    name=$(basename "$file")
    name=${name%".xsd"}
    name=$(echo "$name" | tr '.' '_')
    echo "creating $name from $file"

    mkdir -p gen/"$name"

    xsdgen -pkg "$name" \
           -r '"type Max(\d*)Text string" -> "type Max${1}Text common.Max${1}Text"' \
           -o gen/"$name"/"$name".go \
           "$file"
done
