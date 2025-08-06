// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


//  1 .创建一个名为Voting的合约，包含以下功能：
//  一个mapping来存储候选人的得票数
//  一个vote函数，允许用户投票给某个候选人
//  一个getVotes函数，返回某个候选人的得票数
//  一个resetVotes函数，重置所有候选人的得票数

contract Voting{


    //定义投票者映射,记录是否已经投票
    mapping(address => bool) public voters;

    // 用于存放所有 被投票人的名字, 用户后续循环 , mapping 不能在 solidiy中遍历
    string[] public voltingPersongs;

    // 用于存放所有 投票人及票数
    mapping (string name => int num ) public voltingNum ;

    //根据投票人的 名字, 来为其投票
    function vote(string memory name ) public {    

        //先验证,投票人是否已经投过票
        // require(voters[msg.sender],"You have already voted!");

        //将投票人加入到已经投票的映射中
        voters[msg.sender] = true;

        uint256 i = 0 ;
        for (i = 0 ; i < voltingPersongs.length ; i++){
            if (keccak256(abi.encodePacked(voltingPersongs[i])) == keccak256(abi.encodePacked(name))){
            //    break ;
            }
        }
    
        if(i == voltingPersongs.length){
                voltingPersongs.push(name);
        }

         voltingNum[name] += 1;

    }

   //根据投票人的 名字, 获得其票数
    function getVotes(string memory name) public view returns(int){
      
        return voltingNum[name];

    }
  // 重置所有投票人的 票数
    function resetVotes() public {
        for ( uint256 i = 0 ; i < voltingPersongs.length ; i++){

            delete voltingNum[voltingPersongs[i]];
        }

    }


}