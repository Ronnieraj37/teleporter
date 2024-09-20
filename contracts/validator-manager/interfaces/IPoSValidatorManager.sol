// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IValidatorManager, ValidatorManagerSettings} from "./IValidatorManager.sol";
import {IRewardCalculator} from "./IRewardCalculator.sol";

enum DelegatorStatus {
    Unknown,
    PendingAdded,
    Active,
    PendingRemoved
}

// TODO: visit types of these fields, for example uint64 might be too big for stake duration seconds.
struct PoSValidatorManagerSettings {
    ValidatorManagerSettings baseSettings;
    uint256 minimumStakeAmount;
    uint256 maximumStakeAmount;
    uint64 minimumStakeDuration;
    uint16 minimumDelegationFeeBips;
    uint8 maximumStakeMultiplier;
    IRewardCalculator rewardCalculator;
}

struct Delegator {
    DelegatorStatus status;
    address owner;
    bytes32 validationID;
    uint64 weight;
    uint64 startedAt;
    uint64 endedAt;
    uint64 startingNonce;
    uint64 endingNonce;
}

struct PoSValidatorInfo {
    address owner;
    uint16 delegationFeeBips;
    uint64 minStakeDuration;
}

interface IPoSValidatorManager is IValidatorManager {
    /**
     * @notice Event emitted when a delegator registration is initiated
     * @param delegationID The ID of the delegation
     * @param validationID The ID of the validation period
     * @param delegatorAddress The address of the delegator
     * @param nonce The message nonce used to update the validator weight
     * @param validatorWeight The updated validator weight that is sent to the P-Chain
     * @param delegatorWeight The weight of the delegator
     * @param setWeightMessageID The ID of the Warp message that updates the validator's weight on the P-Chain
     */
    event DelegatorAdded(
        bytes32 indexed delegationID,
        bytes32 indexed validationID,
        address indexed delegatorAddress,
        uint64 nonce,
        uint64 validatorWeight,
        uint64 delegatorWeight,
        bytes32 setWeightMessageID
    );

    /**
     * @notice Event emitted when a delegator registration is completed
     * @param delegationID The ID of the delegation
     * @param validationID The ID of the validation period
     * @param nonce The message nonce used to update the validator weight, as returned by the P-Chain
     * @param startTime The time at which the registration was completed
     */
    event DelegatorRegistered(
        bytes32 indexed delegationID,
        bytes32 indexed validationID,
        uint64 indexed nonce,
        uint256 startTime
    );

    /**
     * @notice Event emitted when delegator removal is initiated
     * @param delegationID The ID of the delegation
     * @param validationID The ID of the validation period
     * @param endTime The time at which the removal was initiated
     */
    event DelegatorRemovalInitialized(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256 endTime
    );

    /**
     * @notice Event emitted when delegator removal is completed
     * @param delegationID The ID of the delegation
     * @param validationID The ID of the validator the delegator was staked to
     * @param rewards The rewards given to the delegator
     * @param fees The portion of the delegator's rewards paid to the validator
     */
    event DelegationEnded(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees
    );

    /**
     * @notice Begins the process of ending an active validation period. The validation period must have been previously
     * started by a successful call to {completeValidatorRegistration} with the given validationID.
     * Any rewards for this validation period will stop accruing when this function is called.
     * @param validationID The ID of the validation being ended.
     * @param includeUptimeProof Whether or not an uptime proof is provided for the validation period.
     * If no uptime proof is provided, the validation uptime will be assumed to be 0.
     * @param messageIndex If {includeUptimeProof} is true, the index of the Warp message to be received providing the
     * uptime proof.
     */
    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external;

    /**
     * @notice Completes the delegator registration process by returning an acknowledgement of the registration of a
     * validationID from the P-Chain. After this function is called, the validator's weight is updated in the contract state.
     * Any P-Chain acknowledgement with a nonce greater than or equal to the nonce used to initialize registration of the
     * delegator is valid, as long as that nonce has been sent by the contract. For the purposes of computing delegation rewards,
     * the delegation is considered active after this function is called.
     * Note: only the specified delegation will be marked as registered, even if the validator weight update
     * message implicitly includes multiple weight changes.
     * @param messageIndex The index of the Warp message to be received providing the acknowledgement.
     * @param delegationID The ID of the delegation being registered.
     */
    function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) external;

    /**
     * @notice Removes a delegator from a completed validation period. The delegator can be in either the active or pending removed
     * state. No uptime proof is required in this case, because it will have been provided by the validator upon their exit.
     * Note that this function can be called by any address to clean up the delegation.
     * @param delegationID The ID of the delegation being removed.
     */
    function endDelegationCompletedValidator(bytes32 delegationID) external;

    /**
     * @notice Begins the process of removing a delegator from a validation period. The delegator must have been previously
     * registered with the given validationID. For the purposes of computing delegation rewards, the delegation period is
     * considered ended when this function is called. In order to be eligible for rewards, an uptime proof must be provided.
     * Note that this function can only be called by the address that registered the delegation.
     * @param delegationID The ID of the delegation being removed.
     * @param includeUptimeProof Whether or not an uptime proof is provided for the validation period.
     * If the validator has completed its validation period, it has already provided an uptime proof, so {includeUptimeProof}
     * will be ignored and can be set to false. If the validator has not completed its validation period and no uptime proof
     * is provided, the validation uptime for the delegation period will be assumed to be 0.
     * @param messageIndex If {includeUptimeProof} is true, the index of the Warp message to be received providing the
     * uptime proof.
     */
    function initializeEndDelegation(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external;

    /**
     * @notice Resubmits a delegator registration or delegator end message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param delegationID The ID of the delegation.
     */
    function resendUpdateDelegation(bytes32 delegationID) external;

    /**
     * @notice Completes the process of ending a delegation by receiving an acknowledgement from the P-Chain.
     * After this function is called, the validator's weight is updated in the contract state.
     * Any P-Chain acknowledgement with a nonce greater than or equal to the nonce used to initialize the end of the
     * delegator's delegation is valid, as long as that nonce has been sent by the contract. This is because the validator
     * weight change pertaining to the delegation ending is included in any subsequent validator weight update messages.
     * Note: only the specified delegation will be marked as completed, even if the validator weight update
     * message implicitly includes multiple weight changes.
     * @param messageIndex The index of the Warp message to be received providing the acknowledgement.
     * @param delegationID The ID of the delegation being removed.
     */
    function completeEndDelegation(uint32 messageIndex, bytes32 delegationID) external;

    /**
     * @notice Withdraws the delegation fees from completed delegations to the owner of the validator.
     * Can currently only be called once the validator has completed its validation period.
     * @param validationID The ID of the validation being ended.
     */
    function claimDelegationFees(bytes32 validationID) external;
}
