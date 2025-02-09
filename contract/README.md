# Smart Contract Setup

## Prerequisites

```bash
# Install Solidity compiler
brew tap ethereum/ethereum
brew install solidity

# Install abigen
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Add GOPATH to your shell (for zsh)
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

## Compile Contract

```bash
# Create build directory if not exists
mkdir -p build

# Generate ABI
solc --abi contract/Store.sol -o build

# Generate bin (bytecode)
solc --bin contract/Store.sol -o build

# Generate Go bindings (with deploy methods)
abigen --bin=build/Store.bin --abi=build/Store.abi --pkg=store --out=Store.go
```

## Contract Details

The `Store.sol` contract is a simple key-value store that:

- Stores key-value pairs in a mapping
- Emits events when items are set
- Has a version string

## Alternative Installation Methods

### Ubuntu

```bash
# Using snap
sudo snap install solc --edge
```

### Build from Source

```bash
# Get go-ethereum
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

For other platforms, check the [official solidity installation guide](https://docs.soliditylang.org/en/v0.8.28/installing-solidity.html).

## Next Steps

- Deploy the smart contract
- Interact with the deployed contract
- Test contract functionality
