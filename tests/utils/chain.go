package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	relayerConfig "github.com/ava-labs/awm-relayer/config"
	"github.com/ava-labs/awm-relayer/peers"
	"github.com/ava-labs/awm-relayer/signature-aggregator/aggregator"
	sigAggConfig "github.com/ava-labs/awm-relayer/signature-aggregator/config"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/eth/tracers"
	"github.com/ava-labs/subnet-evm/ethclient"
	subnetEvmInterfaces "github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/teleporter/tests/interfaces"
	gasUtils "github.com/ava-labs/teleporter/utils/gas-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	CChainPathSpecifier = "C"
)

var NativeTransferGas uint64 = 21_000

var WarpEnabledChainConfig = tmpnet.FlagsMap{
	"log-level":         "debug",
	"warp-api-enabled":  true,
	"local-txs-enabled": true,
	"eth-apis": []string{
		"eth",
		"eth-filter",
		"net",
		"admin",
		"web3",
		"internal-eth",
		"internal-blockchain",
		"internal-transaction",
		"internal-debug",
		"internal-account",
		"internal-personal",
		"debug",
		"debug-tracer",
		"debug-file-tracer",
		"debug-handler",
	},
}

//
// URL utils
//

func HttpToWebsocketURI(uri string, blockchainID string) string {
	return fmt.Sprintf("ws://%s/ext/bc/%s/ws", strings.TrimPrefix(uri, "http://"), blockchainID)
}

func HttpToRPCURI(uri string, blockchainID string) string {
	return fmt.Sprintf("http://%s/ext/bc/%s/rpc", strings.TrimPrefix(uri, "http://"), blockchainID)
}

// Get the host and port from a URI. The URI should be in the format http://host:port or https://host:port
func GetURIHostAndPort(uri string) (string, uint32, error) {
	// At a minimum uri should have http:// of 7 characters
	Expect(len(uri)).Should(BeNumerically(">", 7))
	if uri[:7] == "http://" {
		uri = uri[7:]
	} else if uri[:8] == "https://" {
		uri = uri[8:]
	} else {
		return "", 0, fmt.Errorf("invalid uri: %s", uri)
	}

	// Split the uri into host and port
	hostAndPort := strings.Split(uri, ":")
	Expect(len(hostAndPort)).Should(Equal(2))

	// Parse the port
	port, err := strconv.ParseUint(hostAndPort[1], 10, 32)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse port: %w", err)
	}

	return hostAndPort[0], uint32(port), nil
}

//
// Transaction utils
//

func CreateNativeTransferTransaction(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Transaction {
	fromAddress := crypto.PubkeyToAddress(fromKey.PublicKey)
	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo, fromAddress)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   subnetInfo.EVMChainID,
		Nonce:     nonce,
		To:        &recipient,
		Gas:       NativeTransferGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     amount,
	})

	return SignTransaction(tx, fromKey, subnetInfo.EVMChainID)
}

func SendNativeTransfer(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Receipt {
	tx := CreateNativeTransferTransaction(ctx, subnetInfo, fromKey, recipient, amount)
	return SendTransactionAndWaitForSuccess(ctx, subnetInfo, tx)
}

// Sends a tx, and waits for it to be mined.
// Asserts Receipt.status equals success.
func sendAndWaitForTransaction(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	tx *types.Transaction,
	success bool,
) *types.Receipt {
	err := subnetInfo.RPCClient.SendTransaction(ctx, tx)
	Expect(err).Should(BeNil())

	return waitForTransaction(ctx, subnetInfo, tx.Hash(), success)
}

// Sends a tx, and waits for it to be mined.
// Asserts Receipt.status equals false.
func SendTransactionAndWaitForFailure(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	tx *types.Transaction,
) *types.Receipt {
	return sendAndWaitForTransaction(ctx, subnetInfo, tx, false)
}

// Sends a tx, and waits for it to be mined.
// Asserts Receipt.status equals true.
func SendTransactionAndWaitForSuccess(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	tx *types.Transaction,
) *types.Receipt {
	return sendAndWaitForTransaction(ctx, subnetInfo, tx, true)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals true.
func WaitForTransactionSuccess(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	txHash common.Hash,
) *types.Receipt {
	return waitForTransaction(ctx, subnetInfo, txHash, true)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals false.
func WaitForTransactionFailure(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	txHash common.Hash,
) *types.Receipt {
	return waitForTransaction(ctx, subnetInfo, txHash, false)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals success.
func waitForTransaction(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	txHash common.Hash,
	success bool,
) *types.Receipt {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	receipt, err := WaitMined(cctx, subnetInfo.RPCClient, txHash)
	Expect(err).Should(BeNil())

	if success {
		if receipt.Status == types.ReceiptStatusFailed {
			TraceTransactionAndExit(ctx, subnetInfo.RPCClient, receipt.TxHash)
		}
	} else {
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusFailed))
	}
	return receipt
}

// Polls for a transaction receipt of the given txHash on each queryTicker tick until
// either a transaction receipt returned, or the context is cancelled or expired.
func waitForTransactionReceipt(
	cctx context.Context,
	rpcClient ethclient.Client,
	txHash common.Hash,
) (*types.Receipt, error) {
	queryTicker := time.NewTicker(200 * time.Millisecond)
	defer queryTicker.Stop()
	for {
		receipt, err := rpcClient.TransactionReceipt(cctx, txHash)
		if err == nil {
			return receipt, nil
		}

		if errors.Is(err, subnetEvmInterfaces.NotFound) {
			log.Debug("Transaction not yet mined")
		} else {
			log.Error("Receipt retrieval failed", "err", err)
			return nil, err
		}

		// Wait for the next round.
		select {
		case <-cctx.Done():
			return nil, cctx.Err()
		case <-queryTicker.C:
		}
	}
}

// Signs a transaction using the provided key for the specified chainID
func SignTransaction(tx *types.Transaction, key *ecdsa.PrivateKey, chainID *big.Int) *types.Transaction {
	txSigner := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(tx, txSigner, key)
	Expect(err).Should(BeNil())

	return signedTx
}

// Returns the gasFeeCap, gasTipCap, and nonce the be used when constructing a transaction from fundedAddress
func CalculateTxParams(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	fundedAddress common.Address,
) (*big.Int, *big.Int, uint64) {
	baseFee, err := subnetInfo.RPCClient.EstimateBaseFee(ctx)
	Expect(err).Should(BeNil())

	gasTipCap, err := subnetInfo.RPCClient.SuggestGasTipCap(ctx)
	Expect(err).Should(BeNil())

	nonce, err := subnetInfo.RPCClient.NonceAt(ctx, fundedAddress, nil)
	Expect(err).Should(BeNil())

	gasFeeCap := baseFee.Mul(baseFee, big.NewInt(gasUtils.BaseFeeFactor))
	gasFeeCap.Add(gasFeeCap, big.NewInt(gasUtils.MaxPriorityFeePerGas))

	return gasFeeCap, gasTipCap, nonce
}

// Gomega will print the transaction trace and exit
func TraceTransactionAndExit(ctx context.Context, rpcClient ethclient.Client, txHash common.Hash) {
	Expect(TraceTransaction(ctx, rpcClient, txHash)).Should(Equal(""))
}

func TraceTransaction(ctx context.Context, rpcClient ethclient.Client, txHash common.Hash) string {
	var result interface{}
	ct := "callTracer"
	err := rpcClient.Client().Call(&result, "debug_traceTransaction", txHash.String(), tracers.TraceConfig{Tracer: &ct})
	Expect(err).Should(BeNil())

	jsonStr, err := json.Marshal(result)
	Expect(err).Should(BeNil())

	return string(jsonStr)
}

//
// Block utils
//

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
// Takes a tx hash instead of the full tx in the subnet-evm version of this function.
// Copied and modified from https://github.com/ava-labs/subnet-evm/blob/v0.6.0-fuji/accounts/abi/bind/util.go#L42
func WaitMined(ctx context.Context, rpcClient ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	cctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	receipt, err := waitForTransactionReceipt(cctx, rpcClient, txHash)
	if err != nil {
		return nil, err
	}

	// Check that the block height endpoint returns a block height as high as the block number that the transaction was
	// included in. This is to workaround the issue where multiple nodes behind a public RPC endpoint see
	// transactions/blocks at different points in time. Ideally, all nodes in the network should have seen this block
	// and transaction before returning from WaitMined. The block height endpoint of public RPC endpoints is
	// configured to return the lowest value currently returned by any node behind the load balancer, so waiting for
	// it to be at least as high as the block height specified in the receipt should provide a relatively strong
	// indication that the transaction has been seen widely throughout the network.
	err = waitForBlockHeight(cctx, rpcClient, receipt.BlockNumber.Uint64())
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

// Polls for the eth_blockNumber endpoint for the latest blockheight on each queryTicker tick until
// either the returned height is greater than or equal to the expectedBlockNumber, or the context
// is cancelled or expired.
func waitForBlockHeight(
	cctx context.Context,
	rpcClient ethclient.Client,
	expectedBlockNumber uint64,
) error {
	queryTicker := time.NewTicker(2 * time.Second)
	defer queryTicker.Stop()
	for {
		currentBlockNumber, err := rpcClient.BlockNumber(cctx)
		if err != nil {
			return err
		}

		if currentBlockNumber >= expectedBlockNumber {
			return nil
		} else {
			log.Info("Waiting for block height where transaction was included",
				"blockNumber", expectedBlockNumber)
		}

		// Wait for the next round.
		select {
		case <-cctx.Done():
			return cctx.Err()
		case <-queryTicker.C:
		}
	}
}

//
// Log utils
//

// Returns the first log in 'logs' that is successfully parsed by 'parser'
// Errors and prints a trace of the transaction if no log is found.
func GetEventFromLogsOrTrace[T any](
	ctx context.Context,
	receipt *types.Receipt,
	subnetInfo interfaces.SubnetTestInfo,
	parser func(log types.Log) (T, error),
) T {
	log, err := GetEventFromLogs(receipt.Logs, parser)
	if err != nil {
		TraceTransactionAndExit(ctx, subnetInfo.RPCClient, receipt.TxHash)
	}
	return log
}

// Returns the first log in 'logs' that is successfully parsed by 'parser'
func GetEventFromLogs[T any](logs []*types.Log, parser func(log types.Log) (T, error)) (T, error) {
	for _, log := range logs {
		event, err := parser(*log)
		if err == nil {
			return event, nil
		}
	}
	return *new(T), fmt.Errorf("failed to find %T event in receipt logs", *new(T))
}

//
// Account utils
//

func PrivateKeyToAddress(k *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(k.PublicKey)
}

// Throws a Gomega error if there is a mismatch
func CheckBalance(ctx context.Context, addr common.Address, expectedBalance *big.Int, rpcClient ethclient.Client) {
	bal, err := rpcClient.BalanceAt(ctx, addr, nil)
	Expect(err).Should(BeNil())
	ExpectBigEqual(bal, expectedBalance)
}

//
// Big int utils
//

func ExpectBigEqual(v1 *big.Int, v2 *big.Int) {
	// Compare strings, so gomega will print the numbers if they differ
	Expect(v1.String()).Should(Equal(v2.String()))
}

func BigIntSub(v1 *big.Int, v2 *big.Int) *big.Int {
	return big.NewInt(0).Sub(v1, v2)
}

func BigIntMul(v1 *big.Int, v2 *big.Int) *big.Int {
	return big.NewInt(0).Mul(v1, v2)
}

//
// Network utils
//

func GetPChainInfo(cChainInfo interfaces.SubnetTestInfo) interfaces.SubnetTestInfo {
	pChainBlockchainID, err := info.NewClient(cChainInfo.NodeURIs[0]).GetBlockchainID(context.Background(), "P")
	Expect(err).Should(BeNil())
	return interfaces.SubnetTestInfo{
		BlockchainID: pChainBlockchainID,
		SubnetID:     ids.Empty,
	}
}

func GetTwoSubnets(network interfaces.Network) (
	interfaces.SubnetTestInfo,
	interfaces.SubnetTestInfo,
) {
	subnets := network.GetSubnetsInfo()
	Expect(len(subnets)).Should(BeNumerically(">=", 2))
	return subnets[0], subnets[1]
}

type ChainConfigMap map[string]string

// Sets the chain config in customChainConfigs for the specified subnet
func (m ChainConfigMap) Add(subnet interfaces.SubnetTestInfo, chainConfig string) {
	if subnet.SubnetID == constants.PrimaryNetworkID {
		m[CChainPathSpecifier] = chainConfig
	} else {
		m[subnet.BlockchainID.String()] = chainConfig
	}
}

func GetChainConfigWithOffChainMessages(offChainMessages []avalancheWarp.UnsignedMessage) string {
	// Convert messages to hex
	hexOffChainMessages := []string{}
	for _, message := range offChainMessages {
		hexOffChainMessages = append(hexOffChainMessages, hexutil.Encode(message.Bytes()))
	}

	chainConfig := WarpEnabledChainConfig
	chainConfig["warp-off-chain-messages"] = hexOffChainMessages

	// Marshal the map to JSON
	offChainMessageJson, err := tmpnet.DefaultJSONMarshal(chainConfig)
	Expect(err).Should(BeNil())

	return string(offChainMessageJson)
}

// read in the template file, make the substitutions declared at the beginning
// of the function, write out the instantiation to a temp file, and then return
// the path to that temp file.
func InstantiateGenesisTemplate(
	templateFileName string,
	chainID uint64,
	teleporterContractAddress common.Address,
	teleporterDeployedBytecode string,
	teleporterDeployerAddress common.Address,
) string {
	substitutions := []struct {
		Target string
		Value  string
	}{
		{
			"<EVM_CHAIN_ID>",
			strconv.FormatUint(chainID, 10),
		},
		{
			"<TELEPORTER_MESSENGER_CONTRACT_ADDRESS>",
			teleporterContractAddress.Hex(),
		},
		{
			"<TELEPORTER_MESSENGER_DEPLOYED_BYTECODE>",
			teleporterDeployedBytecode,
		},
		{
			"<TELEPORTER_MESSENGER_DEPLOYER_ADDRESS>",
			teleporterDeployerAddress.Hex(),
		},
	}

	templateFileBytes, err := os.ReadFile(templateFileName)
	Expect(err).Should(BeNil())

	subnetGenesisFile, err := os.CreateTemp(os.TempDir(), "")
	Expect(err).Should(BeNil())

	defer subnetGenesisFile.Close()

	var replaced string = string(templateFileBytes[:])
	for _, s := range substitutions {
		replaced = strings.ReplaceAll(replaced, s.Target, s.Value)
	}

	subnetGenesisFile.WriteString(replaced)

	return subnetGenesisFile.Name()
}

//
// Aggregator utils
//

func NewSignatureAggregator(apiUri string, subnets []ids.ID) *aggregator.SignatureAggregator {
	logger := logging.NoLog{}
	cfg := sigAggConfig.Config{
		PChainAPI: &relayerConfig.APIConfig{
			BaseURL: apiUri,
		},
		InfoAPI: &relayerConfig.APIConfig{
			BaseURL: apiUri,
		},
	}
	trackedSubnets := set.NewSet[ids.ID](len(subnets))
	trackedSubnets.Add(subnets...)
	registry := prometheus.NewRegistry()
	appRequestNetwork, err := peers.NewNetwork(
		logging.Debug,
		registry,
		trackedSubnets,
		&cfg,
	)
	Expect(err).Should(BeNil())

	messageCreator, err := message.NewCreator(
		logger,
		registry,
		constants.DefaultNetworkCompressionType,
		constants.DefaultNetworkMaximumInboundTimeout,
	)
	Expect(err).Should(BeNil())
	return aggregator.NewSignatureAggregator(
		appRequestNetwork,
		logger,
		messageCreator,
	)
}
