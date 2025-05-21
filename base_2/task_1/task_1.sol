// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
 solidity：创建一个名为Voting的合约，包含以下功能：
一个mapping来存储候选人的得票数,
一个vote函数，允许用户投票给某个候选人,
一个getVotes函数，返回某个候选人的得票数,
一个resetVotes函数，重置所有候选人的得票数
*/

contract Voting {
    mapping(string => uint256) private votes;
    string[] private candidates;
    address private owner;

    constructor(string[] memory candidateList) {
        owner = msg.sender;
        candidates = candidateList;
        for (uint i = 0; i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
    }

// 投票给某个候选人
    function vote(string memory candidate) public {
        require(_isCandidate(candidate), "candidate does not exist");
        votes[candidate] += 1;
    }

// 查询某个候选人的得票数
    function getVotes(string memory candidate) public view returns (uint256) {
        require(_isCandidate(candidate), "candidate does not exist");
        return votes[candidate];
    }

// 重置所有候选人的票数（仅合约拥有者可调用）
    function resetVotes() public {
        require(msg.sender == owner, "only owner can reset votes");
        for (uint i = 0; i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
    }

// 判断某个名字是否是候选人
    function _isCandidate(string memory candidate) internal view returns (bool) {
        for (uint i = 0; i < candidates.length; i++) {
            if (keccak256(bytes(candidates[i])) == keccak256(bytes(candidate))) {
                return true;
            }
        }
        return false;
    }

// 查看所有候选人名字
    function getAllCandidates() public view returns (string[] memory) {
        return candidates;
    }
}
