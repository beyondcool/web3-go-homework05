// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;
import "@openzeppelin/contracts/access/Ownable.sol";

contract Counter is Ownable {
  
  uint totalBalance;

  uint public num;

  event DonateEvent(address indexed donor, uint amount);
  event WithdrawEvent(address indexed owner, uint amount);
  event IncrementEvent(uint by);

  constructor() Ownable(msg.sender) {
    
  }

  function donate() public payable {
    require(msg.value > 0, "donate: amount should be positive");
    totalBalance += msg.value;
    emit DonateEvent(msg.sender, msg.value);
  }

  function withdraw() public onlyOwner {
    require(totalBalance > 0, "withdraw: totalBalance should be positive");
    payable(msg.sender).transfer(totalBalance);
    emit WithdrawEvent(msg.sender, totalBalance);
  }

  function inc() public {
    num++;
    emit IncrementEvent(1);
  }

  function incBy(uint by) public {
    require(by > 0, "incBy: increment should be positive");
    num += by;
    emit IncrementEvent(by);
  }
}
