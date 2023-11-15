package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rolluper-dev/brc-rpc/entity"
	"github.com/rolluper-dev/brc-rpc/logic"
	"github.com/rolluper-dev/brc-rpc/resp"
	"github.com/rolluper-dev/brc-rpc/utils"
)

func InscriptionList(c *gin.Context) {
	req := &entity.InscriptionListRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}

	page, size := utils.ValidatePage(req.Page, req.Size)

	list, count, code := logic.InscriptionList(req.Address, page, size)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, page, len(list), count)
}

func Inscription(c *gin.Context) {
	req := &entity.InscriptionRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}
	if req.InscriptionId == "" && req.InscriptionNumber == "" {
		resp.ParameterErr(c, "")
		return
	}

	ret, code := logic.Inscription(req.InscriptionId, req.InscriptionNumber)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.Success(c, ret)
}
