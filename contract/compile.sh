#!/bin/bash

# Compile contract
solc --optimize --evm-version shanghai \
     --abi --bin contract/store.sol \
     -o contract/store --overwrite

# Generate Go bindings
abigen --bin=contract/store/Store.bin \
       --abi=contract/store/Store.abi \
       --pkg=store \
       --out=contract/store/Store.go

echo "Contract compiled and bindings generated successfully!"