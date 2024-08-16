// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {PoSValidatorManagerTest} from "./PoSValidatorManagerTests.t.sol";
import {ERC20TokenStakingManager} from "../ERC20TokenStakingManager.sol";
import {ValidatorManagerSettings} from "../interfaces/IValidatorManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
contract ERC20TokenStakingManagerTest is PoSValidatorManagerTest {
    using SafeERC20 for IERC20;

    ERC20TokenStakingManager public app;
    IERC20 public token;

    function setUp() public virtual {
        // Construct the object under test
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        token = new ExampleERC20();
        app.initialize(
            ValidatorManagerSettings({
                pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                subnetID: DEFAULT_SUBNET_ID,
                maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN
            }),
            DEFAULT_MINIMUM_STAKE,
            DEFAULT_MAXIMUM_STAKE,
            DEFAULT_MINIMUM_STAKE_DURATION,
            IRewardCalculator(address(0)),
            token
        );
        validatorManager = app;
        posValidatorManager = app;
    }

    function _initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature,
        uint256 stakeAmount
    ) internal virtual override returns (bytes32) {
        return
            app.initializeValidatorRegistration(stakeAmount, nodeID, registrationExpiry, signature);
    }

    function _beforeSend(uint64 weight) internal override {
        // ERC20 tokens need to be pre-approved
        token.safeIncreaseAllowance(address(app), app.weightToValue(weight));
    }
}
