package staking

//go:generate go run ../../utils/unpack/unpackgen --out=./Unpack.sol --test_out=./tests/Unpack.t.sol --solc_version="0.8.25" --byte_sizes=4,32,32,8,8,dyn;4,32,1;4,32,8,8;4,32,8;32,32,8,8,dyn