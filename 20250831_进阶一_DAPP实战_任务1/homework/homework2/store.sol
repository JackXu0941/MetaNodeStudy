// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;
contract Store {
  event setCountSet(int256 count );
  int256 public count;
  
  constructor() {
    
  }
  function setCount() external {
    count++;
    emit setCountSet(count);
  }
}