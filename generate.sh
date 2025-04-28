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


#find ./proto_files/ -iname "*.proto" | while read file
#do


 "$SCRIPTPATH/protoc/bin/protoc" \
  --proto_path="$SCRIPTPATH/proto_files" \
  --php_out="$outputPHP" \
  --grpc_out="$outputPHP" \
  --go_out="$outputGO" \
  --go-grpc_out="$outputGO" \
  --plugin=protoc-gen-grpc="$SCRIPTPATH/protoc/grpc_php_plugin" \
  --plugin=protoc-gen-go-grpc="$SCRIPTPATH/protoc/protoc_gen_go_grpc" \
  "$SCRIPTPATH/proto_files/adevent.proto"
# ./protoc/bin/protoc --php_out=$outputPHP --go_out=$outputGO --avro_out=./generated/avro/ --avro_opt=namespace_map=main:exads.avro "$file";
#done


php "$tools/constants2enum.php" "$outputPHP" "Exads/Grpc/GPBMetadata" && \
cd "$outputGO/main" && \
go mod init github.com/LazyNewbie/test-evt/generated/go/main && \
go mod tidy && \
cd "$SCRIPTPATH"  && \
git add "$SCRIPTPATH/generated/*"