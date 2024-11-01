#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source $TELEPORTER_PATH/scripts/constants.sh
source $TELEPORTER_PATH/scripts/versions.sh

export ARCH=$(uname -m)
[ $ARCH = x86_64 ] && ARCH=amd64
echo "ARCH set to $ARCH"

DEFAULT_CONTRACT_LIST="TeleporterMessenger TeleporterRegistry ExampleERC20 ExampleRewardCalculator TestMessenger ValidatorSetSig NativeTokenStakingManager ERC20TokenStakingManager PoAValidatorManager"

PROXY_LIST="TransparentUpgradeableProxy ProxyAdmin"

SUBNET_EVM_LIST="INativeMinter"

CONTRACT_LIST=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -c | --contract) CONTRACT_LIST=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/abi_bindings.sh [OPTIONS]"
    echo "Build contracts and generate Go bindings"
    echo ""
    echo "Options:"
    echo "  -c, --contract <contract_name>          Generate Go bindings for the contract. If empty, generate Go bindings for a default list of contracts"
    echo "  -c, --contract "contract1 contract2"    Generate Go bindings for multiple contracts"
    echo "  -h, --help                              Print this help message"
    exit 0
fi

if ! command -v forge &> /dev/null; then
    echo "forge not found. You can install by calling $TELEPORTER_PATH/scripts/install_foundry.sh" && exit 1
fi

echo "Building subnet-evm abigen"
go install github.com/ava-labs/subnet-evm/cmd/abigen@${SUBNET_EVM_VERSION}

# Force recompile of all contracts to prevent against using previous
# compilations that did not generate new ABI files.
echo "Building Contracts"
cd $TELEPORTER_PATH
forge build --skip test --optimizer-runs 100 --force --extra-output-files abi bin

function convertToLower() {
    if [ "$ARCH" = 'arm64' ]; then
        echo $1 | perl -ne 'print lc'
    else
        echo $1 | sed -e 's/\(.*\)/\L\1/'
    fi
}

function generate_bindings() {
    local contract_names=("$@")
    for contract_name in "${contract_names[@]}"
    do
        path=$(find . -name $contract_name.sol)
        dir=$(dirname $path)
        abi_file=$TELEPORTER_PATH/out/$contract_name.sol/$contract_name.abi.json
        if ! [ -f $abi_file ]; then
            echo "Error: Contract $contract_name abi file not found"
            exit 1
        fi

        echo "Generating Go bindings for $contract_name..."
        gen_path=$TELEPORTER_PATH/abi-bindings/go/$dir/$contract_name
        mkdir -p $gen_path
        $GOPATH/bin/abigen --abi $abi_file \
                        --pkg $(convertToLower $contract_name) \
                        --bin $TELEPORTER_PATH/out/$contract_name.sol/$contract_name.bin \
                        --type $contract_name \
                        --out $gen_path/$contract_name.go
        echo "Done generating Go bindings for $contract_name."
    done
}

contract_names=($CONTRACT_LIST)

# If CONTRACT_LIST is empty, use DEFAULT_CONTRACT_LIST
if [[ -z "${CONTRACT_LIST}" ]]; then
    contract_names=($DEFAULT_CONTRACT_LIST)
fi

cd $TELEPORTER_PATH/contracts
generate_bindings "${contract_names[@]}"

contract_names=($PROXY_LIST)
cd $TELEPORTER_PATH/
forge build --skip test --force --extra-output-files abi bin --contracts lib/openzeppelin-contracts/contracts/proxy/transparent

cd $TELEPORTER_PATH/lib/openzeppelin-contracts/contracts/proxy/transparent
generate_bindings "${contract_names[@]}"

contract_names=($SUBNET_EVM_LIST)
cd $TELEPORTER_PATH/
forge build --skip test --force --extra-output-files abi bin --contracts lib/subnet-evm/contracts/contracts/interfaces

cd $TELEPORTER_PATH/lib/subnet-evm/contracts/contracts/interfaces
generate_bindings "${contract_names[@]}"

exit 0
