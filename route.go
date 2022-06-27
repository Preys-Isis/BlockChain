package routes

import (
	"fmt"
	"net/http"

	"github.com/Preys-Isis/Packages"
	"github.com/gin-gonic/gin"
)

//通用响应消息格式
type RespData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//自定义一套错误码
const (
	RESP_OK         = "0"
	RESP_PARAMERR   = "1"
	RESP_BLOCKERR   = "2"
	RESP_UNKNOWNERR = "3"
)

// code 对应的 msg 值
var respMap map[string]string = map[string]string{
	RESP_OK:         "正常",
	RESP_PARAMERR:   "参数异常",
	RESP_BLOCKERR:   "区块链操作异常",
	RESP_UNKNOWNERR: "未知异常",
}

// 返回对应 code 的 msg 信息
func getMsg(code string) string {
	return respMap[code]
}

//响应消息
func RespJsonMsg(c *gin.Context, r *RespData) {
	r.Msg = getMsg(r.Code)
	c.JSON(http.StatusOK, r)
}

// get 功能函数
func get(c *gin.Context) {
	//组织响应消息
	var resp RespData
	var err error
	resp.Code = "0"

	defer RespJsonMsg(c, &resp)

	//解析请求消息
	// 无

	//业务规则处理
	resp.Data, err = Hblocks.getStr()
	if err != nil {
		resp.Code = RESP_BLOCKERR
		fmt.Println("Failed to getStr ", err)
		return
	}
}

// 接收消息
type Str_Set struct {
	str string `json:"string"`
}

// set 功能函数
func set(c *gin.Context) {
	//组织响应消息
	var resp RespData
	resp.Code = "0"

	defer RespJsonMsg(c, &resp)

	//解析请求消息
	var _string Str_Set
	err := c.ShouldBindJSON(&_string)
	if err != nil {
		resp.Code = RESP_PARAMERR
		fmt.Println("Failed to ShouldBindJSON ", err)
		return
	}

	//业务规则处理
	err = Hblocks.setStr(_string.str)
	if err != nil {
		resp.Code = RESP_BLOCKERR
		fmt.Println("Failed to setStr ", err)
		return
	}
}
