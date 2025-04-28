# Protobuf documantation:
https://protobuf.dev/programming-guides/proto3/


# How to set-up
- Download protobuf prebuilt compiler from https://github.com/protocolbuffers/protobuf/releases
- Unzip it to the `./protoc` directory
- make sure that `./protoc/bin/protoc` is executable 
- run `./build_grpc_plugin.sh` to enable grpc support (this may take about 10 minutes)

# How to use
- add/modify files in `./proto_files` directory
- run `./generate.sh`
- commit the changes (`git add generated/*` will be executed automatically)


# Direcory structure:
- `./proto_files` - the place where all *.proto files should be stored
- `./generated` - the protoc keeps generated php/go files here
- `./protoc` - the directory where the prebuilt protoc compiler is stored
- `./tools` - the directory where the additional tools are stored. Currently only `constants2enum.php` which converts 
  protobuf generated constants to php 8.x enums