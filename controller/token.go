package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rolluper-dev/brc-rpc/constant"
	"github.com/rolluper-dev/brc-rpc/entity"
	"github.com/rolluper-dev/brc-rpc/logic"
	"github.com/rolluper-dev/brc-rpc/resp"
	"github.com/rolluper-dev/brc-rpc/utils"
)

func GetToken(c *gin.Context) {
	req := &entity.GetTokenRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}
	if utils.IsEmpty(req.Token) {
		resp.ParameterErr(c, "token missing")
		return
	}

	response, code := logic.GetToken(req.Token)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.Success(c, response)
}

func TokenList(c *gin.Context) {
	req := &entity.TokenListRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}
	if req.Status != constant.TokenStatusNon && req.Status != constant.TokenStatusMinting && req.Status != constant.TokenStatusComplete {
		resp.ParameterErr(c, "invalid status")
		return
	}

	page, size := utils.ValidatePage(req.Page, req.Size)

	list, count, code := logic.TokenList(req.Status, page, size)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, page, len(list), count)
}

func TokenHolders(c *gin.Context) {
	req := &entity.TokenHoldersRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}

	page, size := utils.ValidatePage(req.Page, req.Size)

	list, count, code := logic.TokenHolder(req.Token, page, size)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, page, size, count)
}

func TokenHistory(c *gin.Context) {
	req := &entity.TokenHistoryRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}

	page, size := utils.ValidatePage(req.Page, req.Size)
	list, count, code := logic.TokenHistory(req.Token, req.Type, req.From, req.To, req.Brc201Chain, req.Brc201Extend, page, size)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, page, size, count)
}
