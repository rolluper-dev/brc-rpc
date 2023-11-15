package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/roullper-dev/brc-rpc/entity"
	"github.com/roullper-dev/brc-rpc/logic"
	"github.com/roullper-dev/brc-rpc/resp"
	"github.com/roullper-dev/brc-rpc/utils"
)

func InscriptionHistory(c *gin.Context) {
	req := &entity.InscriptionHistoryRequest{}
	if err := c.ShouldBind(req); err != nil {
		resp.ParameterErr(c, "")
		return
	}

	page, size := utils.ValidatePage(req.Page, req.Size)
	// id -> number
	ins, _ := logic.Inscription(req.InscriptionId, "")
	if ins == nil {
		resp.ParameterErr(c, "id is error")
		return
	}

	list, count, code := logic.InscriptionTxList(ins.InscriptionNumber, page, size)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessCustom(c, map[string]interface{}{
		"inscriptionId":     req.InscriptionId,
		"inscriptionNumber": ins.InscriptionNumber,
		"page":              page,
		"size":              size,
		"total":             count,
		"transactionList":   list,
	})
}
