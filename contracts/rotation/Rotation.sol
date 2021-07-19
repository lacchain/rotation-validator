// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.8.0 <0.9.0;

import "./IRotation.sol";
import "../access/AccessControl.sol";

contract Rotation is AccessControl,IRotation{

    Validator[] removedValidators;  //list validators which will be removed
    address[] newValidators;  //list validators which will be voted 
    Validator[] validators;  //list of all validators
    mapping (address => uint256) private accountIndexOf; //1 based indexing. 0 means non-existent

    uint256 blockRotation;
   
    constructor(){
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function addInactiveValidator(Validator memory _validator) onlyAdmin external {
        require(_validator.id != address(0), "validator can't be 0x0");
        removedValidators.push(_validator);
        emit ValidatorToInactive(_validator.id);
    }

    function addActiveValidator(address _id) onlyAdmin external {
        require(_id != address(0), "validator can't be 0x0");
        Validator storage _validator = validators[accountIndexOf[_id]-1];
        require(_validator.status != State.INACTIVE_REVIEW, "Validator address is inactive review");
        newValidators.push(_id);
        emit ValidatorToActive(_id);
    }

    function executeRotation() onlyAdmin external {
        blockRotation = block.number;
        for (uint i=0; i<removedValidators.length ;i++){
            Validator memory _validatorRemoved = removedValidators[i];
            Validator storage _validator = validators[accountIndexOf[_validatorRemoved.id]-1];
            _validator.status = _validatorRemoved.status;
        }
        for (uint i=0; i<newValidators.length ;i++){
            Validator storage _validator = validators[accountIndexOf[newValidators[i]]-1];
            _validator.status = State.ACTIVE;
        }
        emit RotationStarted(block.number);
    }

    function getRemovedValidators() override external view returns(Validator[] memory){
        return removedValidators;
    }

    function getNewValidators() override external view returns(address[] memory){
        return newValidators;
    }

    function getValidators() external view returns(Validator[] memory){
        return validators;
    }

    function setStatus(address _id, State _status) onlyAdmin external returns (bool){
        if (accountIndexOf[_id] != 0) {
            Validator storage _validator = validators[accountIndexOf[_id]-1];
            _validator.status = _status;
            emit ChangeValidatorStatus(_id,_status);
            return true;
        }
        return false;
    }

    function isInactive(address _id) external view returns (bool){
        Validator memory _validator = validators[accountIndexOf[_id]-1];
        return (_validator.status==State.ACTIVE);
    }

    function removeValidatorsRound() onlyAdmin external returns (bool){
        delete removedValidators;
        delete newValidators;
        emit ValidatorsRemoved(msg.sender);
        return true;
    }

    function getBlockRotation() external view returns (uint256){
        return blockRotation;
    }

    function addValidator(address _id) onlyAdmin external returns (bool) {
        if (accountIndexOf[_id] == 0) {
            Validator memory _validator = Validator(_id,State.INACTIVE_NEW);
            validators.push(_validator);
            accountIndexOf[_id] = validators.length; 
            emit NewValidatorAdded(_id);
            return true;
        }
        return false;
    }

    function removeValidator(address _id) onlyAdmin external returns (bool) {
        uint256 index = accountIndexOf[_id];
        require(index>0, "Node doesn't exist");
        //move last address into index being vacated (unless we are dealing with last index)
        if (index > validators.length) return false; 
        Validator storage lastValidator = validators[validators.length - 1];
        validators[index - 1].id = lastValidator.id;

        accountIndexOf[lastValidator.id] = index;
        accountIndexOf[_id] = 0;
            
        delete validators[validators.length-1];
        validators.pop();
        emit ValidatorRemoved(_id);
        return true;
    }

    modifier onlyAdmin(){
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "Caller is not Admin");
        _;
    }

    event ValidatorToInactive(address validator);
    event ValidatorToActive(address validator);
    event ChangeValidatorStatus(address,State);
    event NewValidatorAdded(address);
    event ValidatorRemoved(address);
}
