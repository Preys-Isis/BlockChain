pragma solidity^0.8.1;

interface IUSER {
    //注册接口设计
    function register(string memory username, string memory passwd) external;
    //登录接口设计
    function login(string memory username, string memory passwd) external view returns(bool);
}