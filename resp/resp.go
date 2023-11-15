package resp

import (
	"net/http"

	"github.com/rolluper-dev/brc-rpc/resource/db"
	"github.com/rolluper-dev/brc-rpc/utils"

	"github.com/gin-gonic/gin"
)

var EmptyStruct = struct{}{}

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Block string      `json:"block"`
	Data  interface{} `json:"data"`
}

type ListData struct {
	List  interface{} `json:"list"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int64       `json:"total"`
}

func SuccessNil(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:  CodeSuccess,
		Msg:   MsgSuccess,
		Block: db.GetHandlerBlockHeight(),
		Data:  EmptyStruct,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:  CodeSuccess,
		Msg:   MsgSuccess,
		Block: db.GetHandlerBlockHeight(),
		Data:  data,
	})
}

func SuccessList(c *gin.Context, list interface{}, page, amount int, total int64) {
	c.JSON(http.StatusOK, Response{
		Code:  CodeSuccess,
		Msg:   MsgSuccess,
		Block: db.GetHandlerBlockHeight(),
		Data: ListData{
			List:  list,
			Page:  page,
			Size:  amount,
			Total: total,
		},
	})
}

func SuccessCustom(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:  CodeSuccess,
		Msg:   MsgSuccess,
		Block: db.GetHandlerBlockHeight(),
		Data:  data,
	})
}

func Error(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  code2msg[code],
		Data: EmptyStruct,
	})
}

func ParameterErr(c *gin.Context, msg string) {
	if utils.IsEmpty(msg) {
		msg = code2msg[CodeParameterErr]
	}
	c.JSON(http.StatusOK, Response{
		Code: CodeParameterErr,
		Msg:  msg,
		Data: EmptyStruct,
	})
}
