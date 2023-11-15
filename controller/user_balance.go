package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rolluper-dev/brc-rpc/entity"
	"github.com/rolluper-dev/brc-rpc/logic"
	"github.com/rolluper-dev/brc-rpc/resp"
	"github.com/rolluper-dev/brc-rpc/utils"
)

func AddressTokenBalance(c *gin.Context) {
	req := &entity.AddressTokenBalanceRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}

	page, size := utils.ValidatePage(req.Page, req.Size)
	list, count, code := logic.AddressTokenBalance(req.Address, req.Token, page, size)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessCustom(c, map[string]interface{}{
		"address":    req.Address,
		"page":       page,
		"size":       size,
		"total":      count,
		"tokenTotal": count,
		"tokenList":  list,
	})
}

func AddressTransferList(c *gin.Context) {
	req := &entity.AddressTransferListRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}

	page, size := utils.ValidatePage(req.Page, req.Size)
	list, count, code := logic.AddressTransferList(req.Address, req.Token, req.Brc201Chain, page, size)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessCustom(c, map[string]interface{}{
		"address": req.Address,
		"page":    page,
		"size":    size,
		"total":   count,
		"list":    list,
	})
}

func AddressHistory(c *gin.Context) {
	req := &entity.AddressHistoryRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}

	page, size := utils.ValidatePage(req.Page, req.Size)
	list, count, code := logic.AddressHistory(req.Address, req.Token, req.Type, req.Brc201Chain, req.Brc201Extend, page, size)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessCustom(c, map[string]interface{}{
		"page":       page,
		"size":       size,
		"total":      count,
		"tokenTotal": count,
		"list":       list,
	})
}
