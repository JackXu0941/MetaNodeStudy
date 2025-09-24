//SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";


contract MyToken_1  {

    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;

   
    address public owner;
    
    mapping (address => uint256) private _balances;
    
    mapping (address => mapping (address => uint256)) private _allowances;

    
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Mint(address indexed to, uint256 value);
    event FallbackCalled(address indexed sender,  uint256 value ,bytes data);
    event Received(address indexed sender,  uint256 value);

  


    constructor(string memory tokenName,string memory tokenSymbol, uint8 tokenDecimals, uint256  TokenTotalSupply)  {
       
        name = tokenName;
        symbol = tokenSymbol;
        decimals = tokenDecimals;
        totalSupply = TokenTotalSupply * (10 ** uint256(decimals));
       
        // _mint(owner, totalSupply);
        //发动初级转账事件
        emit Transfer(address(0), owner, totalSupply);

        _balances[msg.sender] = totalSupply;
        owner = msg.sender;
       
       
      
    }

    
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

   
    function transfer(address recipient, uint256 amount) public returns (bool) {
        require(recipient != address(0), "ERC20: transfer to the zero address");
        require(_balances[msg.sender]>amount, "Insufficient balance");

        _balances[msg.sender] -= amount;
        _balances[recipient] += amount;
        
        emit Transfer(msg.sender, recipient, amount);

        return true;
    }

  
     function approve(address from,address to, uint256 value) public returns (bool) {
        require(from != address(0), "ERC20: approve to the zero address");
        require(to != address(0), "ERC20: approve to the zero address");
        require(_balances[from]>value, " approve from Insufficient balance");


        _allowances[from][to] = value;
       
        emit Approval(from, to, value);
        return true;
    }

  
     function transferFrom(address from, address to, uint256 value) public returns (bool) {
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");
        require(_allowances[from][msg.sender] >= value, "ERC20: transfer amount exceeds allowance");
        require(_balances[from] >= value, "ERC20: transfer amount exceeds balance");

        _balances[from] -= value;
        _balances[to] += value;
        _allowances[from][msg.sender] -= value;
        
        emit Transfer(from, to, value);
        return true;
     }

  
     function mint(address to, uint256 value) public {
        require(msg.sender == owner, "Only the owner can mint tokens");
        _balances[to] += value;
        totalSupply += value;
        //发送增发事件
        emit Mint(to, value);
        emit Transfer(address(0), to, value);
    }

     function allowance(address tokenOwner, address spender) public view returns (uint256) {
        return _allowances[tokenOwner][spender];
    }


    receive() external payable {
        _balances[address(this)] += msg.value;
        emit Received(msg.sender, msg.value);
    }

    fallback() external payable {
        _balances[address(this)] += msg.value;
        emit FallbackCalled(msg.sender, msg.value, msg.data);
    }


}
