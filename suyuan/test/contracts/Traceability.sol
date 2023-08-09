pragma solidity ^0.8.1;

import "./Goods.sol";
import "./User.sol";
import "./strings.sol";

contract Traceability {
    using strings for string;//ring类型的变量可以使用strings库的函数
    struct GoodsData {
        Goods traceGoods;
        bool valid;
    }
    User users;
    string _category;  
    mapping(string => GoodsData) private goods;
    constructor(string memory goodsTp) public {
        _category = goodsTp;
        users = new User();
    }

    // 注册
    function register(string calldata username, string calldata passwd)  external {
        users.register(username, passwd);
    }

    // 登录
    function login(string calldata username, string calldata passwd)  external view returns(bool) {
        return users.login(username, passwd);
    }

    event newGoodsEvent(string  _addr, string _data, string goodsid);
    
    // 产品上链
    function createGoods(
        string memory _username,
        string memory _passwd,
        string memory data,
        string memory goodsID
    )
    public returns(Goods)
    {
        require(users.login(_username, _passwd), "You must have a legal account");
        require(!goods[goodsID].valid, "The Good exists");
        
        goods[goodsID].valid = true;
        Goods traceGoods = new Goods(_username, data, goodsID);
        emit newGoodsEvent(_username, data, goodsID);
        goods[goodsID].traceGoods = traceGoods;
        return traceGoods;
    }
    // 更改产品状态
    function changeGoodsStatus(
        string memory _username,
        string memory _passwd,
        uint16 goodsStatus,
        string memory data,
        string memory goodsID,
        string memory newaddr
    )
    public
    {
        require(users.login(_username, _passwd), "You must have a legal account");
        require(goods[goodsID].valid, "not exists");
        string memory addr = goods[goodsID].traceGoods.getOwneraddr();
        require(addr.isEqual(_username),"This product is not yours");
        goods[goodsID].traceGoods.changeStatus(goodsStatus, newaddr, data);
    }
      
    function getStatus(string memory goodsId) public view returns(uint16) {
         require(goods[goodsId].valid, "not exists");
         return goods[goodsId].traceGoods.getStatus();
    }

    function getGoods(string memory goodsId) public view returns(Goods.TraceData[] memory) {
         require(goods[goodsId].valid, "not exists");
         return goods[goodsId].traceGoods.getTraceInfo();
    }
}
