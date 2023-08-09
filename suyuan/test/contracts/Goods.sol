pragma solidity ^0.8.1;
pragma experimental ABIEncoderV2;

contract Goods {
    struct TraceData {
        string addr;     //操作员地址
        uint16 status;     //货物状态
        uint timestamp;   //操作时间
        string data;    //产品数据
    }
    string owner_addr;
    string _goodsId;
    uint16 _status;   //当前状态
    TraceData[] _traceData;
        
    constructor(string memory _addr, string memory _data, string memory goodsid) public {
        _goodsId = goodsid;
        owner_addr = _addr;
        _traceData.push(TraceData({addr:_addr, status:0, timestamp:block.timestamp, data:_data}));
    }
    
    function changeStatus(uint16 goodsStatus, string memory _addr, string memory _data) public {
        _status = goodsStatus;
        owner_addr = _addr;
        _traceData.push(TraceData({addr:_addr, status:_status, timestamp:block.timestamp, data:_data}));
    }
      
    function getStatus() public view returns(uint16) {
        return _status;
    }
    
    function getTraceInfo() public view returns(TraceData[] memory) {
        return _traceData;
    }

    function getOwneraddr() public view returns(string memory) {
        return owner_addr;
    }

}