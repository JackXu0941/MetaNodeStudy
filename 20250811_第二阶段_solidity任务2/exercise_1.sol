// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract MyToken is IERC20 {
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;
     // 合约所有者
    address private _owner1;

    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;

    constructor(string memory tokenName,string memory tokenSymbol,uint8 tokenDecimals ,uint256 initialSupply ) {
        name = tokenName;
        symbol = tokenSymbol;
        decimals = tokenDecimals;
        totalSupply = initialSupply * 10**uint256(decimals);
        _balances[msg.sender] = totalSupply;
        _owner1 = msg.sender;
    }

    // 查询余额
    function balanceOf(address account) public view override returns (uint256) {
        return _balances[account];
    }

    //转账
    function transfer(address recipient, uint256 amount) public override returns (bool) {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    //退回授权
    function allowance(address owner, address spender) public view override returns (uint256) {
        return _allowances[owner][spender];
    }

    //授权
    function approve(address spender, uint256 amount) public override returns (bool) {
        _approve(msg.sender, spender, amount);
        return true;
    }

    //授权转账 - 代扣
    function transferFrom(address sender,address recipient, uint256 amount) public override returns (bool) {
        _transfer(sender, recipient, amount);

        uint256 currentAllowance = _allowances[sender][msg.sender];
        require(currentAllowance >= amount, "ERC20: transfer amount exceeds allowance");
        _approve(sender, msg.sender, currentAllowance - amount);

        return true;
    }

    //转账的方法
    function _transfer( address sender, address recipient,uint256 amount) internal {
        require(sender != address(0), "ERC20: transfer from the zero address");
        require(recipient != address(0), "ERC20: transfer to the zero address");

        uint256 senderBalance = _balances[sender];
        require(senderBalance >= amount, "ERC20: transfer amount exceeds balance");

        _balances[sender] = senderBalance - amount;
        _balances[recipient] += amount;

        emit Transfer(sender, recipient, amount);
    }

    //授权的方法
    function _approve(address owner,address spender,uint256 amount) internal {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    
    //增发代币
    function mint(address to, uint256 amount) public {
        require(msg.sender == _owner1, "ERC20: only owner can mint");
        totalSupply += amount;
        _balances[to] += amount;
        emit Transfer(address(0), to, amount);
    }



    //记录日志, 在区块链上可以通过合约查到
    event Received(address sender, uint amount);

    receive() external payable {
        _balances[address(this)] += msg.value;
        emit Received(msg.sender, msg.value);
    }
    //记录日志, 在区块链上可以通过合约查到
    event FallbackCalled(address sender, uint amount, bytes data);

    fallback() external payable {
        _balances[address(this)] += msg.value;
        emit FallbackCalled(msg.sender, msg.value, msg.data);
    }



}

