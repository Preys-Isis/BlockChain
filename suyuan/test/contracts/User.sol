pragma solidity^0.8.1;

import "./IUSER.sol";
import "./strings.sol";

contract User is IUSER{
    using strings for string;//string类型的变量可以使用strings库的函数
    mapping(string=>string) _users;
    address admin;

    constructor() public {
        admin = msg.sender;
    }

    //注册
    function register(string memory username, string memory passwd) override external {
        //检测：用户没有注册过，用户不能为空，密码不能为空
        require(!username.isEqual(""), "username must not null");
        require(!passwd.isEqual(""), "passwd must not null");
        require(_users[username].isEqual(""), "user must not exists");
        _users[username] = passwd;
    }
    //登录
    function login(string memory username, string memory passwd) override external view returns(bool) {
        if(username.isEqual("") || passwd.isEqual("")) {
            return false;
        }
        return _users[username].isEqual(passwd);
    }
}