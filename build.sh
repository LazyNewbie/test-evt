#!/bin/bash

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )";

outputGO="$SCRIPTPATH/generated/go";
outputPHP="$SCRIPTPATH/generated/php";
tools="$SCRIPTPATH/tools";


if [ "$(ls $outputGO | wc -l)" -gt 0 ]; then
    rm -r $outputGO/*;
fi

if [ "$(ls $outputPHP | wc -l)" -gt 0 ]; then
    rm -r $outputPHP/*;
fi


find ./proto_files/ -iname "*.proto" | while read file
do
 "$SCRIPTPATH/protoc/bin/protoc" --php_out="$outputPHP" --go_out="$outputGO" "$file";
# ./protoc/bin/protoc --php_out=$outputPHP --go_out=$outputGO --avro_out=./generated/avro/ --avro_opt=namespace_map=main:exads.avro "$file";
done


php "$tools/constants2enum.php" "$outputPHP" "Exads/Common/GPBMetadata";