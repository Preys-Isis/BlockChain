package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"test/blocks"

	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
)

//通用响应消息格式
type RespData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//自定义一套错误编码
const (
	RESP_OK         = "0"
	RESP_PARAMERR   = "1"
	RESP_LOGINERR   = "2"
	RESP_BLOCKERR   = "3"
	RESP_UNKNOWNERR = "4"
)

//错误码与错误消息的 key => value 映射
var respMap map[string]string = map[string]string{
	RESP_OK:         "正常",
	RESP_PARAMERR:   "参数异常",
	RESP_LOGINERR:   "登录失败",
	RESP_BLOCKERR:   "区块链操作失败",
	RESP_UNKNOWNERR: "未知异常",
}

//返回对应 code 的错误消息
func getCodeMsg(code string) string {
	return respMap[code]
}

//响应消息
func respJsonMsg(c *gin.Context, r *RespData) {
	r.Msg = getCodeMsg(r.Code)
	c.JSON(http.StatusOK, r)
}

//用来接收登录请求消息的消息结构体
type UserMsg struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

//注册功能函数
func Register(c *gin.Context) {
	//3. 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)

	//1. 解析请求消息
	var um UserMsg
	um.UserName = c.Query("username")
	um.PassWord = c.Query("password")
	fmt.Printf("Register:%+v\n", um)
	//2. 业务规则处理
	err := blocks.Register(um.UserName, um.PassWord)
	if err != nil {
		resp.Code = RESP_BLOCKERR
		resp.Data = err
		return
	}
}

//登录功能函数
func Login(c *gin.Context) {
	//3. 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	//1. 解析请求消息
	var ui UserMsg
	ui.UserName = c.Query("username")
	ui.PassWord = c.Query("password")
	fmt.Printf("Login:%+v\n", ui)
	//2. 业务规则处理
	ok, err := blocks.Login(ui.UserName, ui.PassWord)
	if err != nil {
		resp.Code = RESP_BLOCKERR
		resp.Data = err
		fmt.Println("failed to Login ", err)
		return
	}
	if !ok {
		fmt.Println("uername not exits or passwoed err")
		resp.Code = RESP_LOGINERR
		resp.Data = err
		return
	}
	// 设置session
	store := ginsession.FromContext(c)
	store.Set("username", ui.UserName)
	store.Set("password", ui.PassWord)
	err = store.Save()
	if err != nil {
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("save session err ", err)
		return
	}
}

// 产品信息
type GoodsMsg struct {
	DaTa    string `json:"data"`
	GoodsId string `json:"goodsid"`
}

// 产品上链
func CreateGoods(c *gin.Context) {
	//4. 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	//1. 解析请求消息
	var gm GoodsMsg
	gm.DaTa = c.Query("data")
	gm.GoodsId = c.Query("goodsid")
	fmt.Printf("Goods:%+v\n", gm)
	//2. 读取session
	store := ginsession.FromContext(c)
	username, ok := store.Get("username")
	if !ok {
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("get username err ", username)
		return
	}
	password, ok := store.Get("password")
	if !ok {
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("get password err ", password)
		return
	}
	//3. 业务规则处理
	err := blocks.Create_Goods(username.(string), password.(string), gm.DaTa, gm.GoodsId)
	if err != nil {
		resp.Code = RESP_BLOCKERR
		resp.Data = err
		fmt.Println("failed to Create_Goods ", err)
		return
	}
}

type ChangeMsg struct {
	DaTa    string `json:"data"`
	GoodsId string `json:"goodsid"`
	NewAddr string `json:"newaddr"`
	Status  string `json:"status"`
}

func Changestatus(c *gin.Context) {
	//4. 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	// 解析请求消息
	var cm ChangeMsg
	cm.DaTa = c.Query("data")
	cm.GoodsId = c.Query("goodsid")
	cm.NewAddr = c.Query("newaddr")
	cm.Status = c.Query("status")
	fmt.Printf("Goods:%+v\n", cm)
	// 读取session
	store := ginsession.FromContext(c)
	username, ok := store.Get("username")
	if !ok {
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("get username err ", username)
		return
	}
	password, ok := store.Get("password")
	if !ok {
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("get password err ", password)
		return
	}
	// 业务规则处理
	// strconv.Atoi(v)
	status, _ := strconv.Atoi(cm.Status)
	err := blocks.ChangeStatus(username.(string), password.(string), cm.DaTa, cm.GoodsId, cm.NewAddr, uint16(status))
	if err != nil {
		resp.Code = RESP_BLOCKERR
		resp.Data = err
		fmt.Println("failed to ChangeStatus ", err)
		return
	}
}

type GetMsg struct {
	UserName string `json:"username"`
	GoodsId  string `json:"goodsid"`
}

func Getstatus(c *gin.Context) {
	// 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	// 解析请求消息
	var gm GetMsg
	gm.UserName = c.Query("username")
	gm.GoodsId = c.Query("goodsid")
	fmt.Println("GoodsID: " + gm.GoodsId)
	// 业务规则处理
	status, err := blocks.GetGoodsStatus(gm.GoodsId)
	if err != nil {
		resp.Code = RESP_BLOCKERR
		resp.Data = err
		fmt.Println("failed to GetGoodsStatus ", err)
		return
	}
	resp.Data = status
}

func GetgoodsMsg(c *gin.Context) {
	// 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	// 解析请求消息
	var err error
	var gm GetMsg
	gm.UserName = c.Query("username")
	gm.GoodsId = c.Query("goodsid")
	fmt.Println("GoodsID: " + gm.GoodsId)
	// 业务规则处理
	resp.Data, err = blocks.GetGoodsMsg(gm.GoodsId)
	if err != nil {
		resp.Code = RESP_BLOCKERR
		resp.Data = err
		fmt.Println("failed to GetGoodsMsg ", err)
		return
	}
}
