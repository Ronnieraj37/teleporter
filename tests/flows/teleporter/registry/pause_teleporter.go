package registry

import (
	"context"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func PauseTeleporter(network interfaces.Network, teleporter utils.TeleporterTestInfo) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy TestMessenger to Subnets A and B
	//
	ctx := context.Background()
	teleporterAddress := teleporter.TeleporterMessengerAddress(subnetAInfo)
	_, testMessengerA := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(subnetAInfo),
		subnetAInfo,
	)
	testMessengerAddressB, testMessengerB := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(subnetBInfo),
		subnetBInfo,
	)

	// Pause Teleporter on subnet B
	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetBInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := testMessengerB.PauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, subnetBInfo, tx.Hash())
	pauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessengerB.ParseTeleporterAddressPaused)
	Expect(err).Should(BeNil())
	Expect(pauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err := testMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Send a message from subnet A to subnet B, which should fail
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		teleporter,
		subnetAInfo,
		testMessengerA,
		subnetBInfo,
		testMessengerAddressB,
		testMessengerB,
		fundedKey,
		"message_1",
		false)

	// Unpause Teleporter on subnet B
	tx, err = testMessengerB.UnpauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt = utils.WaitForTransactionSuccess(ctx, subnetBInfo, tx.Hash())
	unpauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessengerB.ParseTeleporterAddressUnpaused)
	Expect(err).Should(BeNil())
	Expect(unpauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err = testMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())

	// Send a message from subnet A to subnet B again, which should now succeed
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		teleporter,
		subnetAInfo,
		testMessengerA,
		subnetBInfo,
		testMessengerAddressB,
		testMessengerB,
		fundedKey,
		"message_2",
		true)
}
