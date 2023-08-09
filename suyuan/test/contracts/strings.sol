pragma solidity^0.8.1;

library strings {
    function isEqual(string memory a, string memory b) internal pure returns(bool) {
        //借助hash函数防碰撞特性
        bytes32 hashA = keccak256(abi.encode(a));
        bytes32 hashB = keccak256(abi.encode(b));

        return hashA == hashB;
    }
}