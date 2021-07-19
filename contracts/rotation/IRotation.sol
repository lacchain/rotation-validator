// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.0 <0.9.0;

/**
 * @dev Interface for `Rotation`,
 *
 */
interface IRotation {

    enum State {
        ACTIVE,
        INACTIVE_GOOD,
        INACTIVE_NEW,
        INACTIVE_REVIEW
    }

    struct Validator {
        address id;
        State status;
    }

    /**
     * @dev Returns validators which will be removed
     */
    function getRemovedValidators() external view returns(Validator[] memory);

    /**
     * @dev Returns new validators which will be included
     */
    function getNewValidators() external view returns(address[] memory);

    /**
     * @dev Emitted when a rotation was set at specific future block
     */
    event RotationStarted(uint256 blockNumber);

    /**
     * @dev Emitted when a rotation was set at specific future block
     */
    event ValidatorsRemoved(address validator);

    event InactiveRemoved(address validator);
}