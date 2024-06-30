#!/bin/zsh

CURRENT_DIR=$1

rm -rf ${CURRENT_DIR}/genproto

mkdir -p ${CURRENT_DIR}/genproto

for x in $(find ${CURRENT_DIR}/protos -name "*.proto"); do
  # Extract the directory name relative to protos
  REL_DIR=$(dirname ${x} | sed "s|${CURRENT_DIR}/protos/||")
  
  # Create the corresponding directory in genproto
  OUTPUT_DIR=${CURRENT_DIR}/genproto/${REL_DIR}
  mkdir -p ${OUTPUT_DIR}
  
  # Run protoc command for each proto file
  protoc -I=${CURRENT_DIR}/protos -I /usr/local/go --go_out=${OUTPUT_DIR} \
  --go-grpc_out=${OUTPUT_DIR} ${x}
done
