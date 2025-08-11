// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract BeggingContract is Ownable {
    mapping(address => uint256) public donations;
    address[] public topDonors;
    uint256 public startTime;
    uint256 public endTime;

    event Donation(address indexed donor, uint256 amount);

    constructor(uint256 _startTime, uint256 _endTime) Ownable(msg.sender) {
        startTime = _startTime;
        endTime = _endTime;
    }

    modifier onlyDuringDonationPeriod() {
        require(block.timestamp >= startTime && block.timestamp <= endTime, "Donations are not allowed at this time");
        _;
    }

    function donate() external payable onlyDuringDonationPeriod {
        require(msg.value > 0, "Donation amount must be greater than zero");
        donations[msg.sender] += msg.value;
        emit Donation(msg.sender, msg.value);
        updateTopDonors(msg.sender);
    }

    function withdraw() external onlyOwner {
        payable(owner()).transfer(address(this).balance);
    }

    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }

    function getTopDonors() external view returns (address[3] memory) {
        address[3] memory topThree;
        for (uint i = 0; i < 3 && i < topDonors.length; i++) {
            topThree[i] = topDonors[i];
        }
        return topThree;
    }

    function updateTopDonors(address donor) internal {
        uint256 donorIndex = findDonorIndex(donor);
        if (donorIndex != topDonors.length) {
            // Donor is already in the list, update their position
            for (uint256 i = donorIndex; i > 0; i--) {
                if (donations[topDonors[i - 1]] < donations[donor]) {
                    topDonors[i] = topDonors[i - 1];
                } else {
                    break;
                }
            }
        } else {
            // Donor is not in the list, add them
            topDonors.push(donor);
            donorIndex = topDonors.length - 1;
        }
        topDonors[donorIndex] = donor;
    }

    function findDonorIndex(address donor) internal view returns (uint256) {
        for (uint256 i = 0; i < topDonors.length; i++) {
            if (topDonors[i] == donor) {
                return i;
            }
        }
        return topDonors.length;
    }
}