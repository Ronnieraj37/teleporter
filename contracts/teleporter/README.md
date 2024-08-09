# Teleporter Protocol

## Overview

Teleporter is a messaging protocol built on top of [Avalanche Warp Messaging (AWM)](https://docs.avax.network/learn/avalanche/awm) that provides a developer-friendly interface for sending and receiving cross-chain messages from within the EVM.

The `ITeleporterMessenger` interface provides two primary methods:

- `sendCrossChainMessage`: called by contracts on the origin chain to initiate the sending of a message to a contract on another EVM instance.
- `receiveCrossChainMessage`: called by cross-chain relayers on the destination chain to deliver signed messages to the destination EVM instance.

The `ITeleporterReceiver` interface provides a single method. All contracts that wish to receive Teleporter messages on the destination chain must implement this interface:

- `receiveTeleporterMessage`: called by the Teleporter contract on the destination chain to deliver a message to the destination contract.

> Note: If a contract does not implement `ITeleporterReceiver`, but instead implements [fallback](https://docs.soliditylang.org/en/latest/contracts.html#fallback-function), the fallback function will be called when Teleporter attempts to perform message execution. The message execution is marked as failed if the fallback function reverts, otherwise it is marked as successfully executed.

## Data Flow

<div align="center">
  <img src="../../resources/TeleporterDataFlowDiagram.png?raw=true"/>
</div>

## Properties

Teleporter provides a handful of useful properties to cross-chain applications that Avalanche Warp Messages do not provide by default. These include:

1. Replay protection: Teleporter ensures that a cross-chain message is not delivered multiple times.
2. Retries: In certain edge cases when there is significant validator churn, it is possible for an Avalanche Warp Message to be dropped before a valid aggregate signature is created for it. Teleporter ensures that messages can still be delivered even in this event by allowing for retries of previously submitted messages.
3. Relay incentivization: Teleporter provides a mechanism for messages to optionally incentivize relayers to perform the necessary signature aggregation and pay the transaction fee to broadcast the signed message on the destination chain.
4. Allowed relayers: Teleporter allows users to specify a list of `allowedRelayerAddresses`, where only the specified addresses can relay and deliver the Teleporter message. Leaving this list empty allows all relayers to deliver.
5. Message execution: Teleporter enables cross-chain messages to have direct effect on their destination chain by using `evm.Call()` to invoke the `receiveTeleporterMessage` function of destination contracts that implement the `ITeleporterReceiver` interface.

## Fees

Fees can be paid on a per message basis by specifing the ERC20 asset and amount to be used to incentivize a relayer to deliver the message in the call to `sendCrossChainMessage`. The fee amount is transferred into the control of the Teleporter contract (i.e. locked) before the Warp message is sent. The Teleporter contract tracks the fee amount for each message ID it creates. When it subsequently receives a message back from the destination chain of the original message, the new message will have a list of receipts identifying the relayer that delivered the given message ID. At this point, the fee amount originally locked by Teleporter for the given message will be redeemable by the relayer identified in the receipt. If the initial fee amount was not sufficient to incentivize a relayer, it can be added to by using `addFeeAmount`.

### Message Receipts and Fee Redemption

In order to confirm delivery of a Teleporter message from a source chain to a destination chain, a receipt is included in the next Teleporter message sent in the opposite direction, from the destination chain back to the source chain. This receipt contains the message ID of the original message, as well as the reward address that the delivering relayer specified. That reward address is then able to redeem the corresponding reward on the original chain by calling `redeemRelayerRewards`. The following example illustrates this flow:

- A Teleporter message is sent from Chain A to Chain B, with a relayer incentive of `10` `USDC`. This message is assigned the ID `1` by the Teleporter contract on Chain A.
  - On Chain A, this is done by calling `sendCrossChainMessage`, and providing the `USDC` contract address and amount in the function call.
- A relayer delivers the message on Chain B by calling `receiveCrossChainMessage` and providing its address, `0x123...`
- The Teleporter contract on Chain B stores the relayer address in a receipt for the message ID.
- Some time later, a separate Teleporter message is sent from Chain B to Chain A. The Teleporter contract on Chain B includes the receipt for the original message in this new message.
- When this new message is delivered on Chain A, the Teleporter contract on Chain A reads the receipt and attributes the rewards for delivering the original message (message ID `1`) to the address `0x123...`.
- Address `0x123...` may now call `redeemRelayerRewards` on Chain A, which transfers the `10` `USDC` to its address. If it tries to do this before the receipt is received on Chain A, the call will fail.

It is possible for receipts to get "stuck" on the destination chain in the event that Teleporter traffic between two chains is skewed in one direction. In such a scenario, incoming messages on one chain may cause the rate at which receipts are generated to outpace the rate at which they are sent back to the other chain. To mitigate this, the method `sendSpecifiedReceipts` can be called to immediately send the receipts associated with the given message IDs back to the original chain.

## Required Interface

Teleporter messages are delivered by calling the `receiveTeleporterMessage` function defined by the `ITeleporterReceiver` interface. Contracts must implement this interface in order to be able to receive messages. The first two parameters of `receiveTeleporterMessage` identify the original sender of the given message on the origin chain and are set by the `TeleporterMessenger`. The third parameter to `receiveTeleporterMessage`, is the raw message payload. Applications using Teleporter are responsible for defining the exact format of this payload in a way that can be decoded on the receiving end. For example, applications may encode an action enum value along with the target method parameters on the sending side, then decode this data and route to the target method within `receiveTeleporterMessage`. See `ERC20Bridge.sol` for an example of this approach.

## Message Delivery and Execution

Teleporter is able to ensure that messages are considered delivered even if their execution fails (i.e. reverts) by using `evm.Call()` with a pre-defined gas limit to execute the message payload. This gas limit is specified by each message in the call to `sendCrossChainMessage`. Relayers must provide at least enough gas for the sub-call in addition to the standard gas used by a call to `receiveCrossChainMessage`. In the event that a message execution runs out of gas or reverts for any other reason, the hash of the message payload is stored by the receiving Teleporter contract instance. This allows for the message execution to be retried in the future, with possibly a higher gas limit by calling `retryMessageExecution`. Importantly, a message is still considered delivered on its destination chain even if its execution fails. This allows the relayer of the message to redeem their reward for deliverying the message, because they have no control on whether or not its execution will succeed or not so long as they provide sufficient gas to meet the specified `requiredGasLimit`.

Note that due to [EIP-150](https://eips.ethereum.org/EIPS/eip-150), the lesser of 63/64<sup>ths</sup> of the remaining gas and the `requiredGasLimit` will be provided to the code executed using `evm.Call()`. This creates an edge case where sufficient gas is provided by the relayer at time of the `requiredGasLimit` check, but less than the `requiredGasLimit` is provided for the message execution. In such a case, the message execution may fail due to having less than the `requiredGasLimit` available, but the message would still be considered received. Such a case is only possible if the remaining 1/64<sup>th</sup> of the `requiredGasLimit` is sufficient for executing the remaining logic of `receiveCrossChainMessage` so that the top level transaction does not also revert. Based on the current implementation, a message must have a `requiredGasLimit` of over 1,200,000 gas for this to be possible. In order to avoid this case entirely, it is recommended for applications sending Teleporter messages to add a buffer to the `requiredGasLimit` such that 63/64<sup>ths</sup> of the value passed is sufficient for the message execution.

## Resending a Message

If the sending Subnet's validator set changes, then it's possible for the receiving Subnet to reject the underlying Warp message due to insufficient signing stake. For example, suppose Subnet A has 5 validators with equal stake weight who all sign a Teleporter message sent to Subnet B. 100% of Subnet A's stake has signed the message. Also suppose Subnet B requires 67% of the sending Subnet's stake to have signed a given Warp message in order for it to be accepted. Before the message can be delivered, however, 5 _more_ validators are added to Subnet A's validator set (all with the same stake weight as the original validators), meaning that the Teleporter message was signed by _only 50%_ of Subnet A's stake. Subnet B will reject this message.

Once sent on chain, Warp messages cannot be re-signed by a new validator set in such a scenario. Teleporter, however, does support re-signing via the function `retrySendCrossChainMessage`, which can be called for any message that has not been acknowledged as delivered to its destination. Under the hood, this packages the Teleporter message into a brand new Warp message that is re-signed by the current validator set. 

## Teleporter Messenger Contract Deployment

**Do not deploy the `TeleporterMessenger` contract using `forge create`**. The `TeleporterMessenger` contract must be deployed to the same contract address on every chain. To achieve this, the contract can be deployed using a static transaction that uses Nick's method as documented in [this guide](../..//utils/contract-deployment/README.md). Alternatively, if creating a new Subnet, the contract can be pre-allocated with the proper address and state in the new chain's [genesis file](https://docs.avax.network/build/subnet/upgrade/customize-a-subnet#setting-the-genesis-allocation).

As an example, to include `TeleporterMessenger` `v1.0.0` in the genesis file, include the following values in the `alloc` settings, as documented at the link above. The `storage` values included below correspond to the two contract values that are initialized as part of the default constructor of `TeleporterMessenger`. These are the `ReentrancyGuard` values set in this [abstract contract](../utilities/ReentrancyGuards.sol). Future versions of `TeleporterMessenger` may require different storage value intializations.
```json
"alloc": {
    "0x253b2784c75e510dD0fF1da844684a1aC0aa5fcf": {
        "balance": "0x0",
        "code": "0x608060405234801561001057600080fd5b506004361061014d5760003560e01c8063a8898181116100c3578063df20e8bc1161007c578063df20e8bc1461033b578063e69d606a1461034e578063e6e67bd5146103b6578063ebc3b1ba146103f2578063ecc7042814610415578063fc2d61971461041e57600080fd5b8063a8898181146102b2578063a9a85614146102c5578063b771b3bc146102d8578063c473eef8146102e6578063ccb5f8091461031f578063d127dc9b1461033257600080fd5b8063399b77da11610115578063399b77da1461021957806362448850146102395780638245a1b01461024c578063860a3b061461025f578063892bf4121461027f5780638ac0fd041461029f57600080fd5b80630af5b4ff1461015257806322296c3a1461016d5780632bc8b0bf146101825780632ca40f55146101955780632e27c223146101ee575b600080fd5b61015a610431565b6040519081526020015b60405180910390f35b61018061017b366004612251565b610503565b005b61015a61019036600461226e565b6105f8565b6101e06101a336600461226e565b6005602090815260009182526040918290208054835180850190945260018201546001600160a01b03168452600290910154918301919091529082565b604051610164929190612287565b6102016101fc36600461226e565b610615565b6040516001600160a01b039091168152602001610164565b61015a61022736600461226e565b60009081526005602052604090205490565b61015a6102473660046122ae565b61069e565b61018061025a366004612301565b6106fc565b61015a61026d36600461226e565b60066020526000908152604090205481565b61029261028d366004612335565b6108a7565b6040516101649190612357565b6101806102ad366004612377565b6108da565b61015a6102c03660046123af565b610b19565b61015a6102d3366004612426565b610b5c565b6102016005600160991b0181565b61015a6102f43660046124be565b6001600160a01b03918216600090815260096020908152604080832093909416825291909152205490565b61018061032d3660046124f7565b610e03565b61015a60025481565b61015a61034936600461226e565b61123d565b61039761035c36600461226e565b600090815260056020908152604091829020825180840190935260018101546001600160a01b03168084526002909101549290910182905291565b604080516001600160a01b039093168352602083019190915201610164565b6103dd6103c436600461226e565b6004602052600090815260409020805460019091015482565b60408051928352602083019190915201610164565b61040561040036600461226e565b611286565b6040519015158152602001610164565b61015a60035481565b61018061042c36600461251e565b61129c565b600254600090806104fe576005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015610481573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a59190612564565b9050806104cd5760405162461bcd60e51b81526004016104c49061257d565b60405180910390fd5b600281905560405181907f1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b10490600090a25b919050565b3360009081526009602090815260408083206001600160a01b0385168452909152902054806105855760405162461bcd60e51b815260206004820152602860248201527f54656c65706f727465724d657373656e6765723a206e6f2072657761726420746044820152676f2072656465656d60c01b60648201526084016104c4565b3360008181526009602090815260408083206001600160a01b03871680855290835281842093909355518481529192917f3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88910160405180910390a36105f46001600160a01b03831633836114f7565b5050565b600081815260046020526040812061060f9061155f565b92915050565b6000818152600760205260408120546106825760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f74604482015268081c9958d95a5d995960ba1b60648201526084016104c4565b506000908152600860205260409020546001600160a01b031690565b60006001600054146106c25760405162461bcd60e51b81526004016104c4906125c4565b60026000556106f16106d383612804565b833560009081526004602052604090206106ec90611572565b61167c565b600160005592915050565b60016000541461071e5760405162461bcd60e51b81526004016104c4906125c4565b6002600081815590546107379060408401358435610b19565b6000818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b03168452600290910154838301529081019190915280519192509061079f5760405162461bcd60e51b81526004016104c4906128a7565b6000836040516020016107b29190612b42565b60408051601f19818403018152919052825181516020830120919250146107eb5760405162461bcd60e51b81526004016104c490612b55565b8360400135837f2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8868560200151604051610826929190612b9e565b60405180910390a360405163ee5b48eb60e01b81526005600160991b019063ee5b48eb90610858908490600401612c23565b6020604051808303816000875af1158015610877573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089b9190612564565b50506001600055505050565b604080518082019091526000808252602082015260008381526004602052604090206108d390836118bc565b9392505050565b6001600054146108fc5760405162461bcd60e51b81526004016104c4906125c4565b600260005560018054146109225760405162461bcd60e51b81526004016104c490612c36565b60026001558061098c5760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f727465724d657373656e6765723a207a65726f2061646469746960448201526e1bdb985b0819995948185b5bdd5b9d608a1b60648201526084016104c4565b6001600160a01b0382166109b25760405162461bcd60e51b81526004016104c490612c7b565b6000838152600560205260409020546109dd5760405162461bcd60e51b81526004016104c4906128a7565b6000838152600560205260409020600101546001600160a01b03838116911614610a6f5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642066656560448201527f20617373657420636f6e7472616374206164647265737300000000000000000060648201526084016104c4565b6000610a7b8383611981565b600085815260056020526040812060020180549293508392909190610aa1908490612ce5565b909155505060008481526005602052604090819020905185917fc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e191610b03916001019081546001600160a01b0316815260019190910154602082015260400190565b60405180910390a2505060018080556000555050565b60408051306020820152908101849052606081018390526080810182905260009060a0016040516020818303038152906040528051906020012090509392505050565b6000600160005414610b805760405162461bcd60e51b81526004016104c4906125c4565b60026000818155905490866001600160401b03811115610ba257610ba2612607565b604051908082528060200260200182016040528015610be757816020015b6040805180820190915260008082526020820152815260200190600190039081610bc05790505b5090508660005b81811015610d6c5760008a8a83818110610c0a57610c0a612cf8565b90506020020135905060006007600083815260200190815260200160002054905080600003610c8a5760405162461bcd60e51b815260206004820152602660248201527f54656c65706f727465724d657373656e6765723a2072656365697074206e6f7460448201526508199bdd5b9960d21b60648201526084016104c4565b610c958d8783610b19565b8214610d095760405162461bcd60e51b815260206004820152603a60248201527f54656c65706f727465724d657373656e6765723a206d6573736167652049442060448201527f6e6f742066726f6d20736f7572636520626c6f636b636861696e00000000000060648201526084016104c4565b6000828152600860209081526040918290205482518084019093528383526001600160a01b03169082018190528651909190879086908110610d4d57610d4d612cf8565b602002602001018190525050505080610d6590612d0e565b9050610bee565b506040805160c0810182528b815260006020820152610df0918101610d96368b90038b018b612d27565b8152602001600081526020018888808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250938552505060408051928352602080840190915290920152508361167c565b60016000559a9950505050505050505050565b6001805414610e245760405162461bcd60e51b81526004016104c490612c36565b60026001556040516306f8253560e41b815263ffffffff8316600482015260009081906005600160991b0190636f82535090602401600060405180830381865afa158015610e76573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610e9e9190810190612da3565b9150915080610f015760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642077617260448201526870206d65737361676560b81b60648201526084016104c4565b60208201516001600160a01b03163014610f785760405162461bcd60e51b815260206004820152603260248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206f726960448201527167696e2073656e646572206164647265737360701b60648201526084016104c4565b60008260400151806020019051810190610f929190612f40565b90506000610f9e610431565b90508082604001511461100d5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206465736044820152701d1a5b985d1a5bdb8818da185a5b881251607a1b60648201526084016104c4565b8351825160009161101f918490610b19565b600081815260076020526040902054909150156110945760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465724d657373656e6765723a206d65737361676520616c7260448201526c1958591e481c9958d95a5d9959609a1b60648201526084016104c4565b6110a2338460a00151611ae9565b6111005760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20756e617574686f72697a6560448201526832103932b630bcb2b960b91b60648201526084016104c4565b61110e818460000151611b61565b6001600160a01b0386161561114557600081815260086020526040902080546001600160a01b0319166001600160a01b0388161790555b60c08301515160005b81811015611192576111828488600001518760c00151848151811061117557611175612cf8565b6020026020010151611bd3565b61118b81612d0e565b905061114e565b50604080518082018252855181526001600160a01b038916602080830191909152885160009081526004909152919091206111cc91611cfb565b336001600160a01b03168660000151837f292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e348a8860405161120d929190613150565b60405180910390a460e0840151511561122f5761122f82876000015186611d57565b505060018055505050505050565b600254600090806112605760405162461bcd60e51b81526004016104c49061257d565b600060035460016112719190612ce5565b905061127e828583610b19565b949350505050565b600081815260076020526040812054151561060f565b60018054146112bd5760405162461bcd60e51b81526004016104c490612c36565b60026001819055546000906112d59084908435610b19565b600081815260066020526040902054909150806113045760405162461bcd60e51b81526004016104c4906128a7565b80836040516020016113169190612b42565b60405160208183030381529060405280519060200120146113495760405162461bcd60e51b81526004016104c490612b55565b600061135b6080850160608601612251565b6001600160a01b03163b116113cf5760405162461bcd60e51b815260206004820152603460248201527f54656c65706f727465724d657373656e6765723a2064657374696e6174696f6e604482015273206164647265737320686173206e6f20636f646560601b60648201526084016104c4565b604051849083907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c90600090a360008281526006602090815260408083208390558691611420918701908701612251565b61142d60e0870187613174565b60405160240161144094939291906131ba565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b179052905060006114886114816080870160608801612251565b5a84611e8a565b9050806114eb5760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f727465724d657373656e6765723a20726574727920657865637560448201526a1d1a5bdb8819985a5b195960aa1b60648201526084016104c4565b50506001805550505050565b6040516001600160a01b03831660248201526044810182905261155a90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611ea4565b505050565b8054600182015460009161060f916131e5565b6060600061158960056115848561155f565b611f76565b9050806000036115d85760408051600080825260208201909252906115d0565b60408051808201909152600080825260208201528152602001906001900390816115a95790505b509392505050565b6000816001600160401b038111156115f2576115f2612607565b60405190808252806020026020018201604052801561163757816020015b60408051808201909152600080825260208201528152602001906001900390816116105790505b50905060005b828110156115d05761164e85611f8c565b82828151811061166057611660612cf8565b60200260200101819052508061167590612d0e565b905061163d565b600080611687610431565b9050600060036000815461169a90612d0e565b919050819055905060006116b383876000015184610b19565b90506000604051806101000160405280848152602001336001600160a01b031681526020018860000151815260200188602001516001600160a01b0316815260200188606001518152602001886080015181526020018781526020018860a00151815250905060008160405160200161172c91906131f8565b60405160208183030381529060405290506000808960400151602001511115611794576040890151516001600160a01b031661177a5760405162461bcd60e51b81526004016104c490612c7b565b604089015180516020909101516117919190611981565b90505b6040805180820182528a820151516001600160a01b039081168252602080830185905283518085018552865187830120815280820184815260008a815260058452869020915182555180516001830180546001600160a01b03191691909516179093559101516002909101558a51915190919086907f2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df890611838908890869061320b565b60405180910390a360405163ee5b48eb60e01b81526005600160991b019063ee5b48eb9061186a908690600401612c23565b6020604051808303816000875af1158015611889573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118ad9190612564565b50939998505050505050505050565b60408051808201909152600080825260208201526118d98361155f565b82106119315760405162461bcd60e51b815260206004820152602160248201527f5265636569707451756575653a20696e646578206f7574206f6620626f756e646044820152607360f81b60648201526084016104c4565b8260020160008385600001546119479190612ce5565b81526020808201929092526040908101600020815180830190925280548252600101546001600160a01b0316918101919091529392505050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa1580156119ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119ee9190612564565b9050611a056001600160a01b038516333086612058565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015611a4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a709190612564565b9050818111611ad65760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016104c4565b611ae082826131e5565b95945050505050565b60008151600003611afc5750600161060f565b815160005b81811015611b5657846001600160a01b0316848281518110611b2557611b25612cf8565b60200260200101516001600160a01b031603611b465760019250505061060f565b611b4f81612d0e565b9050611b01565b506000949350505050565b80600003611bc15760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f206d657373616760448201526665206e6f6e636560c81b60648201526084016104c4565b60009182526007602052604090912055565b6000611be484848460000151610b19565b6000818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b031684526002909101548383015290810191909152805191925090611c3b575050505050565b60008281526005602090815260408083208381556001810180546001600160a01b03191690556002018390558382018051830151878401516001600160a01b0390811686526009855283862092515116855292528220805491929091611ca2908490612ce5565b9250508190555082602001516001600160a01b031684837fd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf578460200151604051611cec919061321e565b60405180910390a45050505050565b6001820180548291600285019160009182611d1583612d0e565b90915550815260208082019290925260400160002082518155910151600190910180546001600160a01b0319166001600160a01b039092169190911790555050565b80608001515a1015611db95760405162461bcd60e51b815260206004820152602560248201527f54656c65706f727465724d657373656e6765723a20696e73756666696369656e604482015264742067617360d81b60648201526084016104c4565b80606001516001600160a01b03163b600003611dda5761155a838383612096565b602081015160e0820151604051600092611df892869260240161323e565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b17905260608301516080840151919250600091611e3d919084611e8a565b905080611e5657611e4f858585612096565b5050505050565b604051849086907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c90600090a35050505050565b60008060008084516020860160008989f195945050505050565b6000611ef9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661210b9092919063ffffffff16565b80519091501561155a5780806020019051810190611f179190613268565b61155a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016104c4565b6000818310611f8557816108d3565b5090919050565b604080518082019091526000808252602082015281546001830154819003611ff65760405162461bcd60e51b815260206004820152601960248201527f5265636569707451756575653a20656d7074792071756575650000000000000060448201526064016104c4565b60008181526002840160208181526040808420815180830190925280548252600180820180546001600160a01b03811685870152888852959094529490556001600160a01b031990921690559061204e908390612ce5565b9093555090919050565b6040516001600160a01b03808516602483015283166044820152606481018290526120909085906323b872dd60e01b90608401611523565b50505050565b806040516020016120a791906131f8565b60408051601f1981840301815282825280516020918201206000878152600690925291902055829084907f4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985906120fe9085906131f8565b60405180910390a3505050565b606061127e848460008585600080866001600160a01b031685876040516121329190613283565b60006040518083038185875af1925050503d806000811461216f576040519150601f19603f3d011682016040523d82523d6000602084013e612174565b606091505b509150915061218587838387612190565b979650505050505050565b606083156121ff5782516000036121f8576001600160a01b0385163b6121f85760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016104c4565b508161127e565b61127e83838151156122145781518083602001fd5b8060405162461bcd60e51b81526004016104c49190612c23565b6001600160a01b038116811461224357600080fd5b50565b80356104fe8161222e565b60006020828403121561226357600080fd5b81356108d38161222e565b60006020828403121561228057600080fd5b5035919050565b828152606081016108d3602083018480516001600160a01b03168252602090810151910152565b6000602082840312156122c057600080fd5b81356001600160401b038111156122d657600080fd5b820160e081850312156108d357600080fd5b600061010082840312156122fb57600080fd5b50919050565b60006020828403121561231357600080fd5b81356001600160401b0381111561232957600080fd5b61127e848285016122e8565b6000806040838503121561234857600080fd5b50508035926020909101359150565b815181526020808301516001600160a01b0316908201526040810161060f565b60008060006060848603121561238c57600080fd5b83359250602084013561239e8161222e565b929592945050506040919091013590565b6000806000606084860312156123c457600080fd5b505081359360208301359350604090920135919050565b60008083601f8401126123ed57600080fd5b5081356001600160401b0381111561240457600080fd5b6020830191508360208260051b850101111561241f57600080fd5b9250929050565b60008060008060008086880360a081121561244057600080fd5b8735965060208801356001600160401b038082111561245e57600080fd5b61246a8b838c016123db565b90985096508691506040603f198401121561248457600080fd5b60408a01955060808a013592508083111561249e57600080fd5b50506124ac89828a016123db565b979a9699509497509295939492505050565b600080604083850312156124d157600080fd5b82356124dc8161222e565b915060208301356124ec8161222e565b809150509250929050565b6000806040838503121561250a57600080fd5b823563ffffffff811681146124dc57600080fd5b6000806040838503121561253157600080fd5b8235915060208301356001600160401b0381111561254e57600080fd5b61255a858286016122e8565b9150509250929050565b60006020828403121561257657600080fd5b5051919050565b60208082526027908201527f54656c65706f727465724d657373656e6765723a207a65726f20626c6f636b636040820152661a185a5b88125160ca1b606082015260800190565b60208082526023908201527f5265656e7472616e63794775617264733a2073656e646572207265656e7472616040820152626e637960e81b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561263f5761263f612607565b60405290565b60405160c081016001600160401b038111828210171561263f5761263f612607565b60405161010081016001600160401b038111828210171561263f5761263f612607565b604051601f8201601f191681016001600160401b03811182821017156126b2576126b2612607565b604052919050565b6000604082840312156126cc57600080fd5b6126d461261d565b905081356126e18161222e565b808252506020820135602082015292915050565b60006001600160401b0382111561270e5761270e612607565b5060051b60200190565b600082601f83011261272957600080fd5b8135602061273e612739836126f5565b61268a565b82815260059290921b8401810191818101908684111561275d57600080fd5b8286015b848110156127815780356127748161222e565b8352918301918301612761565b509695505050505050565b60006001600160401b038211156127a5576127a5612607565b50601f01601f191660200190565b600082601f8301126127c457600080fd5b81356127d26127398261278c565b8181528460208386010111156127e757600080fd5b816020850160208301376000918101602001919091529392505050565b600060e0823603121561281657600080fd5b61281e612645565b8235815261282e60208401612246565b602082015261284036604085016126ba565b60408201526080830135606082015260a08301356001600160401b038082111561286957600080fd5b61287536838701612718565b608084015260c085013591508082111561288e57600080fd5b5061289b368286016127b3565b60a08301525092915050565b60208082526026908201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f7460408201526508199bdd5b9960d21b606082015260800190565b6000808335601e1984360301811261290457600080fd5b83016020810192503590506001600160401b0381111561292357600080fd5b8060051b360382131561241f57600080fd5b8183526000602080850194508260005b858110156129735781356129588161222e565b6001600160a01b031687529582019590820190600101612945565b509495945050505050565b6000808335601e1984360301811261299557600080fd5b83016020810192503590506001600160401b038111156129b457600080fd5b8060061b360382131561241f57600080fd5b8183526000602080850194508260005b858110156129735781358752828201356129ef8161222e565b6001600160a01b03168784015260409687019691909101906001016129d6565b6000808335601e19843603018112612a2657600080fd5b83016020810192503590506001600160401b03811115612a4557600080fd5b80360382131561241f57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000610100823584526020830135612a948161222e565b6001600160a01b0316602085015260408381013590850152612ab860608401612246565b6001600160a01b0316606085015260808381013590850152612add60a08401846128ed565b8260a0870152612af08387018284612935565b92505050612b0160c084018461297e565b85830360c0870152612b148382846129c6565b92505050612b2560e0840184612a0f565b85830360e0870152612b38838284612a54565b9695505050505050565b6020815260006108d36020830184612a7d565b60208082526029908201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206d65736040820152680e6c2ceca40d0c2e6d60bb1b606082015260800190565b606081526000612bb16060830185612a7d565b90506108d3602083018480516001600160a01b03168252602090810151910152565b60005b83811015612bee578181015183820152602001612bd6565b50506000910152565b60008151808452612c0f816020860160208601612bd3565b601f01601f19169290920160200192915050565b6020815260006108d36020830184612bf7565b60208082526025908201527f5265656e7472616e63794775617264733a207265636569766572207265656e7460408201526472616e637960d81b606082015260800190565b60208082526034908201527f54656c65706f727465724d657373656e6765723a207a65726f2066656520617360408201527373657420636f6e7472616374206164647265737360601b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b8082018082111561060f5761060f612ccf565b634e487b7160e01b600052603260045260246000fd5b600060018201612d2057612d20612ccf565b5060010190565b600060408284031215612d3957600080fd5b6108d383836126ba565b80516104fe8161222e565b600082601f830112612d5f57600080fd5b8151612d6d6127398261278c565b818152846020838601011115612d8257600080fd5b61127e826020830160208701612bd3565b805180151581146104fe57600080fd5b60008060408385031215612db657600080fd5b82516001600160401b0380821115612dcd57600080fd5b9084019060608287031215612de157600080fd5b604051606081018181108382111715612dfc57612dfc612607565b604052825181526020830151612e118161222e565b6020820152604083015182811115612e2857600080fd5b612e3488828601612d4e565b6040830152509350612e4b91505060208401612d93565b90509250929050565b600082601f830112612e6557600080fd5b81516020612e75612739836126f5565b82815260059290921b84018101918181019086841115612e9457600080fd5b8286015b84811015612781578051612eab8161222e565b8352918301918301612e98565b600082601f830112612ec957600080fd5b81516020612ed9612739836126f5565b82815260069290921b84018101918181019086841115612ef857600080fd5b8286015b848110156127815760408189031215612f155760008081fd5b612f1d61261d565b8151815284820151612f2e8161222e565b81860152835291830191604001612efc565b600060208284031215612f5257600080fd5b81516001600160401b0380821115612f6957600080fd5b908301906101008286031215612f7e57600080fd5b612f86612667565b82518152612f9660208401612d43565b602082015260408301516040820152612fb160608401612d43565b60608201526080830151608082015260a083015182811115612fd257600080fd5b612fde87828601612e54565b60a08301525060c083015182811115612ff657600080fd5b61300287828601612eb8565b60c08301525060e08301518281111561301a57600080fd5b61302687828601612d4e565b60e08301525095945050505050565b600081518084526020808501945080840160005b838110156129735781516001600160a01b031687529582019590820190600101613049565b600081518084526020808501945080840160005b83811015612973576130a8878351805182526020908101516001600160a01b0316910152565b6040969096019590820190600101613082565b60006101008251845260018060a01b0360208401511660208501526040830151604085015260608301516130fa60608601826001600160a01b03169052565b506080830151608085015260a08301518160a086015261311c82860182613035565b91505060c083015184820360c0860152613136828261306e565b91505060e083015184820360e0860152611ae08282612bf7565b6001600160a01b038316815260406020820181905260009061127e908301846130bb565b6000808335601e1984360301811261318b57600080fd5b8301803591506001600160401b038211156131a557600080fd5b60200191503681900382131561241f57600080fd5b8481526001600160a01b0384166020820152606060408201819052600090612b389083018486612a54565b8181038181111561060f5761060f612ccf565b6020815260006108d360208301846130bb565b606081526000612bb160608301856130bb565b81516001600160a01b03168152602080830151908201526040810161060f565b8381526001600160a01b0383166020820152606060408201819052600090611ae090830184612bf7565b60006020828403121561327a57600080fd5b6108d382612d93565b60008251613295818460208701612bd3565b919091019291505056fea2646970667358221220586881dd1413fe17197100ceb55646481dae802ef65d37df603c3915f51a4b6364736f6c63430008120033",
        "storage": {
            "0x0000000000000000000000000000000000000000000000000000000000000000": "0x0000000000000000000000000000000000000000000000000000000000000001",
            "0x0000000000000000000000000000000000000000000000000000000000000001": "0x0000000000000000000000000000000000000000000000000000000000000001"
        },
        "nonce": 1
    },
    "0x618FEdD9A45a8C456812ecAAE70C671c6249DfaC": {
        "balance": "0x0",
        "nonce": 1
    }
}
```

The values above are taken from the `v1.0.0` [release artifacts](https://github.com/ava-labs/teleporter/releases/tag/v1.0.0). The contract address, deployed bytecode, and deployer address are unique per major release. All of the other values should remain the same.