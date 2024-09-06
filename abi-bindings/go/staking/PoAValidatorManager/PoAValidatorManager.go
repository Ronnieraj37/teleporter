// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poavalidatormanager

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type ValidatorManagerSettings struct {
	PChainBlockchainID [32]byte
	SubnetID           [32]byte
	MaximumHourlyChurn uint8
}

// PoAValidatorManagerMetaData contains all meta data concerning the PoAValidatorManager contract.
var PoAValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeValidators\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeEndValidation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getWeight\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structValidatorManagerSettings\",\"components\":[{\"name\":\"pChainBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maximumHourlyChurn\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorRegistration\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nodeID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendEndValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendRegisterValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodEnded\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodRegistered\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemovalInitialized\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"IncorrectInputLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516126d03803806126d083398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b6125838061014d5f395ff3fe608060405234801561000f575f80fd5b50600436106100cb575f3560e01c80638da5cb5b11610088578063b771b3bc11610063578063b771b3bc146101f8578063bee0a03f14610206578063c151c0b914610219578063f2fde38b1461022c575f80fd5b80638da5cb5b1461018e57806397fb70d4146101d2578063a3a65e48146101e5575f80fd5b80630322ed98146100cf5780630cdd0985146100e45780633aaa9f2514610135578063467ef06f1461014857806366435abf1461015b578063715018a614610186575b5f80fd5b6100e26100dd366004611e4e565b61023f565b005b6101226100f2366004611e4e565b5f9081527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb08602052604090205490565b6040519081526020015b60405180910390f35b610122610143366004611f12565b6103a0565b6100e2610156366004611fb5565b6103c0565b61016e610169366004611e4e565b610743565b6040516001600160401b03909116815260200161012c565b6100e2610757565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b03909116815260200161012c565b6100e26101e0366004611e4e565b61076a565b6100e26101f3366004611fb5565b610776565b6101ba6005600160991b0181565b6100e2610214366004611e4e565b610933565b6100e2610227366004611ff3565b610a02565b6100e261023a36600461202f565b610b10565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb066020526040812080545f805160206125578339815191529291906102879061204a565b90501180156102b9575060035f83815260078301602052604090205460ff1660058111156102b7576102b7612082565b145b6103225760405162461bcd60e51b815260206004820152602f60248201527f56616c696461746f724d616e616765723a2056616c696461746f72206e6f742060448201526e1c195b991a5b99c81c995b5bdd985b608a1b60648201526084015b60405180910390fd5b5f82815260068201602052604090819020905163ee5b48eb60e01b81526005600160991b019163ee5b48eb9161035b9190600401612096565b6020604051808303815f875af1158015610377573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061039b9190612120565b505050565b5f6103a9610b4a565b6103b584868585610ba5565b90505b949350505050565b5f805160206125578339815191525f6103d88361100d565b90505f806103e983604001516111ee565b91509150801561044e5760405162461bcd60e51b815260206004820152602a60248201527f56616c696461746f724d616e616765723a20726567697374726174696f6e20736044820152691d1a5b1b081d985b1a5960b21b6064820152608401610319565b5f828152600785016020526040808220815161010081019092528054829060ff16600581111561048057610480612082565b600581111561049157610491612082565b81526001820154602082015260028201546001600160a01b03811660408301526001600160401b03600160a01b909104811660608301526003928301548082166080840152600160401b8104821660a0840152600160801b8104821660c0840152600160c01b90041660e0909101529091505f908251600581111561051857610518612082565b1480610536575060018251600581111561053457610534612082565b145b6105955760405162461bcd60e51b815260206004820152602a60248201527f56616c696461746f724d616e616765723a20696e76616c69642076616c696461604482015269746f722073746174757360b01b6064820152608401610319565b6003825160058111156105aa576105aa612082565b036105d157505f83815260068601602052604081206004916105cc9190611e04565b6105d5565b5060055b6020808301515f908152600888019091526040812055818160058111156105fe576105fe612082565b9081600581111561061157610611612082565b9052505f84815260078701602052604090208251815484929190829060ff1916600183600581111561064557610645612082565b021790555060208201516001820155604082015160028201805460608501516001600160401b03908116600160a01b026001600160e01b03199092166001600160a01b039094169390931717905560808301516003909201805460a085015160c086015160e0909601518416600160c01b026001600160c01b03968516600160801b02969096166001600160801b03918516600160401b026001600160801b031990931695909416949094171792909216179190911790558151600581111561071057610710612082565b60405185907f1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16905f90a350505050505050565b5f61074d8261133f565b60a0015192915050565b61075f610b4a565b6107685f611463565b565b610773816114d3565b50565b5f805160206125578339815191525f61078e8361100d565b90505f8061079f83604001516111ee565b91509150806108015760405162461bcd60e51b815260206004820152602860248201527f56616c696461746f724d616e616765723a20526567697374726174696f6e206e6044820152671bdd081d985b1a5960c21b6064820152608401610319565b5f8281526005850160205260408120805461081b9061204a565b905011801561084d575060015f83815260078601602052604090205460ff16600581111561084b5761084b612082565b145b6108695760405162461bcd60e51b815260040161031990612137565b5f828152600585016020526040812061088191611e04565b5f8281526007850160208181526040808420805460ff191660021781556003810180546001600160401b0342818116600160801b0267ffffffffffffffff60801b19909316929092178355600190930154875260088b0185528387208990559588905293835292548151600160401b90910490931683529082019290925283917ff8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568910160405180910390a25050505050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb056020526040812080545f8051602061255783398151915292919061097b9061204a565b90501180156109ad575060015f83815260078301602052604090205460ff1660058111156109ab576109ab612082565b145b6109c95760405162461bcd60e51b815260040161031990612137565b5f82815260058201602052604090819020905163ee5b48eb60e01b81526005600160991b019163ee5b48eb9161035b9190600401612096565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610a465750825b90505f826001600160401b03166001148015610a615750303b155b905081158015610a6f575080155b15610a8d5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610ab757845460ff60401b1916600160401b1785555b610ac187876118e9565b8315610b0757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b610b18610b4a565b6001600160a01b038116610b4157604051631e4fbdf760e01b81525f6004820152602401610319565b61077381611463565b33610b7c7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107685760405163118cdaa760e01b8152336004820152602401610319565b5f610bae611907565b5f80516020612557833981519152426001600160401b038516118015610be857506001600160401b038416610be6426202a300612192565b115b610c4a5760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f724d616e616765723a20696e76616c69642072656769737460448201526c726174696f6e2065787069727960981b6064820152608401610319565b85610ca15760405162461bcd60e51b815260206004820152602160248201527f56616c696461746f724d616e616765723a20696e76616c6964206e6f646520496044820152601160fa1b6064820152608401610319565b5f86815260088201602052604090205415610d0f5760405162461bcd60e51b815260206004820152602860248201527f56616c696461746f724d616e616765723a206e6f646520494420616c72656164604482015267792061637469766560c01b6064820152608401610319565b8251603014610d765760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f724d616e616765723a20696e76616c696420626c7350756260448201526c0d8d2c696caf240d8cadccee8d609b1b6064820152608401610319565b610d7f85611951565b5f80610dc66040518060a00160405280856001015481526020018a8152602001896001600160401b03168152602001886001600160401b0316815260200187815250611af2565b5f82815260058601602052604090209193509150610de482826121f6565b5060405163ee5b48eb60e01b81525f906005600160991b019063ee5b48eb90610e119085906004016122d7565b6020604051808303815f875af1158015610e2d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e519190612120565b604080516101008101825260018152602081018c90529192508101336001600160a01b031681526001600160401b038a1660208083018290525f604080850182905260608501939093526080840181905260a0909301839052868352600788019052902081518154829060ff19166001836005811115610ed357610ed3612082565b021790555060208201516001820155604082015160028201805460608501516001600160401b03908116600160a01b026001600160e01b03199092166001600160a01b039094169390931717905560808301516003909201805460a085015160c086015160e0909601518416600160c01b026001600160c01b03968516600160801b02969096166001600160801b03918516600160401b026001600160801b03199093169590941694909417179290921617919091179055610f9483611bf5565b50604080516001600160401b03808b1682528916602082015282918b9186917f79b81620b81daf2c08cd5bb3dbb79e75d2d7a87f52171fde5aadc8c47823026e910160405180910390a45090925050506103b860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60408051606080820183525f8083526020830181905292820152905f805160206125578339815191526040516306f8253560e41b815263ffffffff851660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa158015611082573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526110a99190810190612318565b91509150806111105760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f724d616e616765723a20496e76696e76616c6964616c696460448201526c2077617270206d65737361676560981b6064820152608401610319565b82548251146111735760405162461bcd60e51b815260206004820152602960248201527f56616c696461746f724d616e616765723a20696e76616c696420736f757263656044820152680818da185a5b88125160ba1b6064820152608401610319565b60208201516001600160a01b0316156111e65760405162461bcd60e51b815260206004820152602f60248201527f56616c696461746f724d616e616765723a20696e76616c6964206f726967696e60448201526e2073656e646572206164647265737360881b6064820152608401610319565b509392505050565b5f8082516027146112535760405162461bcd60e51b815260206004820152602960248201527f56616c696461746f724d657373616765733a20696e76616c6964206d657373616044820152680ceca40d8cadccee8d60bb1b6064820152608401610319565b5f805f8061126087611ca2565b9296509094509250905060f084901c156112c85760405162461bcd60e51b815260206004820152602360248201527f56616c696461746f724d657373616765733a20696e76616c696420636f64656360448201526208125160ea1b6064820152608401610319565b60e083901c60021461132c5760405162461bcd60e51b815260206004820152602760248201527f56616c696461746f724d657373616765733a20696e76616c6964206d657373616044820152666765207479706560c81b6064820152608401610319565b909660f89190911c151595509350505050565b6040805161010080820183525f8083526020808401829052838501829052606084018290526080840182905260a0840182905260c0840182905260e084018290528582527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0790528390208351918201909352825491925f8051602061255783398151915292829060ff1660058111156113da576113da612082565b60058111156113eb576113eb612082565b81526001820154602082015260028201546001600160a01b03811660408301526001600160401b03600160a01b909104811660608301526003909201548083166080830152600160401b8104831660a0830152600160801b8104831660c0830152600160c01b900490911660e0909101529392505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0760205260408082208151610100810190925280545f8051602061255783398151915293929190829060ff16600581111561153457611534612082565b600581111561154557611545612082565b8152600182015460208201526002808301546001600160a01b03811660408401526001600160401b03600160a01b909104811660608401526003909301548084166080840152600160401b8104841660a0840152600160801b8104841660c0840152600160c01b900490921660e090910152909150815160058111156115cd576115cd612082565b146116295760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f724d616e616765723a2076616c696461746f72206e6f742060448201526561637469766560d01b6064820152608401610319565b60408101516001600160a01b0316336001600160a01b0316146116a35760405162461bcd60e51b815260206004820152602c60248201527f56616c696461746f724d616e616765723a2073656e646572206e6f742076616c60448201526b34b230ba37b91037bbb732b960a11b6064820152608401610319565b6116b08160a00151611951565b60038152426001600160401b031660e08201525f83815260078301602052604090208151815483929190829060ff191660018360058111156116f4576116f4612082565b021790555060208201516001820155604082015160028201805460608501516001600160401b03908116600160a01b026001600160e01b03199092166001600160a01b039094169390931717905560808301516003909201805460a085015160c086015160e0909601518416600160c01b026001600160c01b03968516600160801b02969096166001600160801b03918516600160401b026001600160801b031990931695909416949094171792909216179190911790555f61180b846117ba81611bf5565b5f604080515f6020820152600160e01b6022820152602681019490945260c092831b6001600160c01b031990811660468601529190921b16604e830152805160368184030181526056909201905290565b5f858152600685016020526040902090915061182782826121f6565b5060405163ee5b48eb60e01b81525f906005600160991b019063ee5b48eb906118549085906004016122d7565b6020604051808303815f875af1158015611870573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118949190612120565b60a0840151604080516001600160401b039092168252426020830152919250829187917f13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67910160405180910390a35050505050565b6118f1611cf5565b6118fa82611d3e565b61190381611d5f565b5050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161194b57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb02545f805160206125578339815191529060ff165f0361198f575050565b60408051606081018252600383015480825260048401546001600160401b038082166020850152600160401b9091041692820192909252904290610e10906119d790836123ea565b106119f3576001600160401b0384166040830152808252611a12565b8382604001818151611a0591906123fd565b6001600160401b03169052505b5f826020015183604001516064611a299190612424565b611a33919061244f565b600285015490915060ff9081169082161115611aae5760405162461bcd60e51b815260206004820152603460248201527f56616c696461746f724d616e616765723a206d6178696d756d20686f75726c796044820152730818da1d5c9b881c985d1948195e18d95959195960621b6064820152608401610319565b5050805160038301556020810151600490920180546040909201516001600160401b03908116600160401b026001600160801b031990931693169290921717905550565b5f6060826080015151603014611b5c5760405162461bcd60e51b815260206004820152602960248201527f5374616b696e674d657373616765733a20696e76616c6964207369676e6174756044820152680e4ca40d8cadccee8d60bb1b6064820152608401610319565b5f805f855f01518660200151876040015188608001518960600151604051602001611b8d9796959493929190612480565b6040516020818303038152906040529050600281604051611bae91906124f6565b602060405180830381855afa158015611bc9573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190611bec9190612120565b94909350915050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb076020526040812060030180545f80516020612557833981519152916001600160401b0390911690819084611c4d83612511565b91906101000a8154816001600160401b0302191690836001600160401b03160217905550508092505050919050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f805f808451602714611cd45784516040516241835b60e31b8152600481019190915260276024820152604401610319565b50505050602081015160228201516026830151604690930151919390929190565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661076857604051631afcd79f60e31b815260040160405180910390fd5b611d46611cf5565b611d4e611d70565b611d56611d80565b61077381611d88565b611d67611cf5565b61077381611df4565b611d78611cf5565b610768611dfc565b610768611cf5565b611d90611cf5565b80355f8051602061255783398151915290815560208201357fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0155611dda6060830160408401612536565b600291909101805460ff191660ff90921691909117905550565b610b18611cf5565b611c7c611cf5565b508054611e109061204a565b5f825580601f10611e1f575050565b601f0160209004905f5260205f209081019061077391905b80821115611e4a575f8155600101611e37565b5090565b5f60208284031215611e5e575f80fd5b5035919050565b80356001600160401b0381168114611e7b575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715611eb657611eb6611e80565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611ee457611ee4611e80565b604052919050565b5f6001600160401b03821115611f0457611f04611e80565b50601f01601f191660200190565b5f805f8060808587031215611f25575f80fd5b611f2e85611e65565b935060208501359250611f4360408601611e65565b915060608501356001600160401b03811115611f5d575f80fd5b8501601f81018713611f6d575f80fd5b8035611f80611f7b82611eec565b611ebc565b818152886020838501011115611f94575f80fd5b816020840160208301375f6020838301015280935050505092959194509250565b5f60208284031215611fc5575f80fd5b813563ffffffff81168114611fd8575f80fd5b9392505050565b6001600160a01b0381168114610773575f80fd5b5f808284036080811215612005575f80fd5b6060811215612012575f80fd5b50829150606083013561202481611fdf565b809150509250929050565b5f6020828403121561203f575f80fd5b8135611fd881611fdf565b600181811c9082168061205e57607f821691505b60208210810361207c57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b5f60208083525f84546120a88161204a565b806020870152604060018084165f81146120c957600181146120e557612112565b60ff19851660408a0152604084151560051b8a01019550612112565b895f5260205f205f5b858110156121095781548b82018601529083019088016120ee565b8a016040019650505b509398975050505050505050565b5f60208284031215612130575f80fd5b5051919050565b60208082526027908201527f56616c696461746f724d616e616765723a20696e76616c69642076616c6964616040820152661d1a5bdb88125160ca1b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156121a5576121a561217e565b92915050565b601f82111561039b57805f5260205f20601f840160051c810160208510156121d05750805b601f840160051c820191505b818110156121ef575f81556001016121dc565b5050505050565b81516001600160401b0381111561220f5761220f611e80565b6122238161221d845461204a565b846121ab565b602080601f831160018114612256575f841561223f5750858301515b5f19600386901b1c1916600185901b1785556122ad565b5f85815260208120601f198616915b8281101561228457888601518255948401946001909101908401612265565b50858210156122a157878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f5b838110156122cf5781810151838201526020016122b7565b50505f910152565b602081525f82518060208401526122f58160408501602087016122b5565b601f01601f19169190910160400192915050565b80518015158114611e7b575f80fd5b5f8060408385031215612329575f80fd5b82516001600160401b038082111561233f575f80fd5b9084019060608287031215612352575f80fd5b61235a611e94565b8251815260208084015161236d81611fdf565b82820152604084015183811115612382575f80fd5b80850194505087601f850112612396575f80fd5b835192506123a6611f7b84611eec565b83815288828587010111156123b9575f80fd5b6123c8848383018488016122b5565b806040840152508195506123dd818801612309565b9450505050509250929050565b818103818111156121a5576121a561217e565b6001600160401b0381811683821601908082111561241d5761241d61217e565b5092915050565b6001600160401b038181168382160280821691908281146124475761244761217e565b505092915050565b5f6001600160401b038084168061247457634e487b7160e01b5f52601260045260245ffd5b92169190910492915050565b61ffff60f01b8860f01b16815263ffffffff60e01b8760e01b1660028201528560068201528460268201525f6001600160401b0360c01b808660c01b16604684015284516124d581604e8601602089016122b5565b60c09490941b1691909201604e810191909152605601979650505050505050565b5f82516125078184602087016122b5565b9190910192915050565b5f6001600160401b0380831681810361252c5761252c61217e565b6001019392505050565b5f60208284031215612546575f80fd5b813560ff81168114611fd8575f80fdfee92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb00a164736f6c6343000819000a",
}

// PoAValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PoAValidatorManagerMetaData.ABI instead.
var PoAValidatorManagerABI = PoAValidatorManagerMetaData.ABI

// PoAValidatorManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoAValidatorManagerMetaData.Bin instead.
var PoAValidatorManagerBin = PoAValidatorManagerMetaData.Bin

// DeployPoAValidatorManager deploys a new Ethereum contract, binding an instance of PoAValidatorManager to it.
func DeployPoAValidatorManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *PoAValidatorManager, error) {
	parsed, err := PoAValidatorManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoAValidatorManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoAValidatorManager{PoAValidatorManagerCaller: PoAValidatorManagerCaller{contract: contract}, PoAValidatorManagerTransactor: PoAValidatorManagerTransactor{contract: contract}, PoAValidatorManagerFilterer: PoAValidatorManagerFilterer{contract: contract}}, nil
}

// PoAValidatorManager is an auto generated Go binding around an Ethereum contract.
type PoAValidatorManager struct {
	PoAValidatorManagerCaller     // Read-only binding to the contract
	PoAValidatorManagerTransactor // Write-only binding to the contract
	PoAValidatorManagerFilterer   // Log filterer for contract events
}

// PoAValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoAValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoAValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoAValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoAValidatorManagerSession struct {
	Contract     *PoAValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PoAValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoAValidatorManagerCallerSession struct {
	Contract *PoAValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PoAValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoAValidatorManagerTransactorSession struct {
	Contract     *PoAValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PoAValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoAValidatorManagerRaw struct {
	Contract *PoAValidatorManager // Generic contract binding to access the raw methods on
}

// PoAValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoAValidatorManagerCallerRaw struct {
	Contract *PoAValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PoAValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoAValidatorManagerTransactorRaw struct {
	Contract *PoAValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoAValidatorManager creates a new instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManager(address common.Address, backend bind.ContractBackend) (*PoAValidatorManager, error) {
	contract, err := bindPoAValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManager{PoAValidatorManagerCaller: PoAValidatorManagerCaller{contract: contract}, PoAValidatorManagerTransactor: PoAValidatorManagerTransactor{contract: contract}, PoAValidatorManagerFilterer: PoAValidatorManagerFilterer{contract: contract}}, nil
}

// NewPoAValidatorManagerCaller creates a new read-only instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*PoAValidatorManagerCaller, error) {
	contract, err := bindPoAValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerCaller{contract: contract}, nil
}

// NewPoAValidatorManagerTransactor creates a new write-only instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PoAValidatorManagerTransactor, error) {
	contract, err := bindPoAValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerTransactor{contract: contract}, nil
}

// NewPoAValidatorManagerFilterer creates a new log filterer instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PoAValidatorManagerFilterer, error) {
	contract, err := bindPoAValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerFilterer{contract: contract}, nil
}

// bindPoAValidatorManager binds a generic wrapper to an already deployed contract.
func bindPoAValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoAValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoAValidatorManager *PoAValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoAValidatorManager.Contract.PoAValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoAValidatorManager *PoAValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.PoAValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoAValidatorManager *PoAValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.PoAValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoAValidatorManager *PoAValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoAValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoAValidatorManager *PoAValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoAValidatorManager *PoAValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoAValidatorManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerSession) WARPMESSENGER() (common.Address, error) {
	return _PoAValidatorManager.Contract.WARPMESSENGER(&_PoAValidatorManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _PoAValidatorManager.Contract.WARPMESSENGER(&_PoAValidatorManager.CallOpts)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x0cdd0985.
//
// Solidity: function activeValidators(bytes32 nodeID) view returns(bytes32)
func (_PoAValidatorManager *PoAValidatorManagerCaller) ActiveValidators(opts *bind.CallOpts, nodeID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PoAValidatorManager.contract.Call(opts, &out, "activeValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActiveValidators is a free data retrieval call binding the contract method 0x0cdd0985.
//
// Solidity: function activeValidators(bytes32 nodeID) view returns(bytes32)
func (_PoAValidatorManager *PoAValidatorManagerSession) ActiveValidators(nodeID [32]byte) ([32]byte, error) {
	return _PoAValidatorManager.Contract.ActiveValidators(&_PoAValidatorManager.CallOpts, nodeID)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x0cdd0985.
//
// Solidity: function activeValidators(bytes32 nodeID) view returns(bytes32)
func (_PoAValidatorManager *PoAValidatorManagerCallerSession) ActiveValidators(nodeID [32]byte) ([32]byte, error) {
	return _PoAValidatorManager.Contract.ActiveValidators(&_PoAValidatorManager.CallOpts, nodeID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_PoAValidatorManager *PoAValidatorManagerCaller) GetWeight(opts *bind.CallOpts, validationID [32]byte) (uint64, error) {
	var out []interface{}
	err := _PoAValidatorManager.contract.Call(opts, &out, "getWeight", validationID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_PoAValidatorManager *PoAValidatorManagerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _PoAValidatorManager.Contract.GetWeight(&_PoAValidatorManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_PoAValidatorManager *PoAValidatorManagerCallerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _PoAValidatorManager.Contract.GetWeight(&_PoAValidatorManager.CallOpts, validationID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoAValidatorManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerSession) Owner() (common.Address, error) {
	return _PoAValidatorManager.Contract.Owner(&_PoAValidatorManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCallerSession) Owner() (common.Address, error) {
	return _PoAValidatorManager.Contract.Owner(&_PoAValidatorManager.CallOpts)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.CompleteEndValidation(&_PoAValidatorManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.CompleteEndValidation(&_PoAValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.CompleteValidatorRegistration(&_PoAValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.CompleteValidatorRegistration(&_PoAValidatorManager.TransactOpts, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xc151c0b9.
//
// Solidity: function initialize((bytes32,bytes32,uint8) settings, address initialOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) Initialize(opts *bind.TransactOpts, settings ValidatorManagerSettings, initialOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "initialize", settings, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc151c0b9.
//
// Solidity: function initialize((bytes32,bytes32,uint8) settings, address initialOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) Initialize(settings ValidatorManagerSettings, initialOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.Initialize(&_PoAValidatorManager.TransactOpts, settings, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc151c0b9.
//
// Solidity: function initialize((bytes32,bytes32,uint8) settings, address initialOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) Initialize(settings ValidatorManagerSettings, initialOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.Initialize(&_PoAValidatorManager.TransactOpts, settings, initialOwner)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x97fb70d4.
//
// Solidity: function initializeEndValidation(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "initializeEndValidation", validationID)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x97fb70d4.
//
// Solidity: function initializeEndValidation(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) InitializeEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.InitializeEndValidation(&_PoAValidatorManager.TransactOpts, validationID)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x97fb70d4.
//
// Solidity: function initializeEndValidation(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) InitializeEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.InitializeEndValidation(&_PoAValidatorManager.TransactOpts, validationID)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x3aaa9f25.
//
// Solidity: function initializeValidatorRegistration(uint64 weight, bytes32 nodeID, uint64 registrationExpiry, bytes signature) returns(bytes32 validationID)
func (_PoAValidatorManager *PoAValidatorManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, weight uint64, nodeID [32]byte, registrationExpiry uint64, signature []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "initializeValidatorRegistration", weight, nodeID, registrationExpiry, signature)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x3aaa9f25.
//
// Solidity: function initializeValidatorRegistration(uint64 weight, bytes32 nodeID, uint64 registrationExpiry, bytes signature) returns(bytes32 validationID)
func (_PoAValidatorManager *PoAValidatorManagerSession) InitializeValidatorRegistration(weight uint64, nodeID [32]byte, registrationExpiry uint64, signature []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.InitializeValidatorRegistration(&_PoAValidatorManager.TransactOpts, weight, nodeID, registrationExpiry, signature)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x3aaa9f25.
//
// Solidity: function initializeValidatorRegistration(uint64 weight, bytes32 nodeID, uint64 registrationExpiry, bytes signature) returns(bytes32 validationID)
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) InitializeValidatorRegistration(weight uint64, nodeID [32]byte, registrationExpiry uint64, signature []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.InitializeValidatorRegistration(&_PoAValidatorManager.TransactOpts, weight, nodeID, registrationExpiry, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.RenounceOwnership(&_PoAValidatorManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.RenounceOwnership(&_PoAValidatorManager.TransactOpts)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.ResendEndValidatorMessage(&_PoAValidatorManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.ResendEndValidatorMessage(&_PoAValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.ResendRegisterValidatorMessage(&_PoAValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.ResendRegisterValidatorMessage(&_PoAValidatorManager.TransactOpts, validationID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.TransferOwnership(&_PoAValidatorManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.TransferOwnership(&_PoAValidatorManager.TransactOpts, newOwner)
}

// PoAValidatorManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoAValidatorManager contract.
type PoAValidatorManagerInitializedIterator struct {
	Event *PoAValidatorManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoAValidatorManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoAValidatorManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoAValidatorManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerInitialized represents a Initialized event raised by the PoAValidatorManager contract.
type PoAValidatorManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoAValidatorManagerInitializedIterator, error) {

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerInitializedIterator{contract: _PoAValidatorManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerInitialized)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseInitialized(log types.Log) (*PoAValidatorManagerInitialized, error) {
	event := new(PoAValidatorManagerInitialized)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PoAValidatorManager contract.
type PoAValidatorManagerOwnershipTransferredIterator struct {
	Event *PoAValidatorManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoAValidatorManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoAValidatorManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoAValidatorManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PoAValidatorManager contract.
type PoAValidatorManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoAValidatorManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerOwnershipTransferredIterator{contract: _PoAValidatorManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerOwnershipTransferred)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PoAValidatorManagerOwnershipTransferred, error) {
	event := new(PoAValidatorManagerOwnershipTransferred)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodCreatedIterator struct {
	Event *PoAValidatorManagerValidationPeriodCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoAValidatorManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerValidationPeriodCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoAValidatorManagerValidationPeriodCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoAValidatorManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      [32]byte
	RegisterValidationMessageID [32]byte
	Weight                      *big.Int
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0x79b81620b81daf2c08cd5bb3dbb79e75d2d7a87f52171fde5aadc8c47823026e.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes32 indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][32]byte, registerValidationMessageID [][32]byte) (*PoAValidatorManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerValidationPeriodCreatedIterator{contract: _PoAValidatorManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0x79b81620b81daf2c08cd5bb3dbb79e75d2d7a87f52171fde5aadc8c47823026e.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes32 indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][32]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerValidationPeriodCreated)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0x79b81620b81daf2c08cd5bb3dbb79e75d2d7a87f52171fde5aadc8c47823026e.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes32 indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*PoAValidatorManagerValidationPeriodCreated, error) {
	event := new(PoAValidatorManagerValidationPeriodCreated)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodEndedIterator struct {
	Event *PoAValidatorManagerValidationPeriodEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoAValidatorManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerValidationPeriodEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoAValidatorManagerValidationPeriodEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoAValidatorManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*PoAValidatorManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerValidationPeriodEndedIterator{contract: _PoAValidatorManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerValidationPeriodEnded)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*PoAValidatorManagerValidationPeriodEnded, error) {
	event := new(PoAValidatorManagerValidationPeriodEnded)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodRegisteredIterator struct {
	Event *PoAValidatorManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoAValidatorManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerValidationPeriodRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoAValidatorManagerValidationPeriodRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoAValidatorManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*PoAValidatorManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerValidationPeriodRegisteredIterator{contract: _PoAValidatorManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerValidationPeriodRegistered)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*PoAValidatorManagerValidationPeriodRegistered, error) {
	event := new(PoAValidatorManagerValidationPeriodRegistered)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidatorRemovalInitializedIterator struct {
	Event *PoAValidatorManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoAValidatorManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerValidatorRemovalInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoAValidatorManagerValidatorRemovalInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoAValidatorManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*PoAValidatorManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerValidatorRemovalInitializedIterator{contract: _PoAValidatorManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerValidatorRemovalInitialized)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*PoAValidatorManagerValidatorRemovalInitialized, error) {
	event := new(PoAValidatorManagerValidatorRemovalInitialized)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
