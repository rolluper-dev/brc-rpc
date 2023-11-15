package logic

import (
	"strconv"

	"github.com/rolluper-dev/brc-rpc/dao"
	"github.com/rolluper-dev/brc-rpc/entity"
	"github.com/rolluper-dev/brc-rpc/resource/log"
	"github.com/rolluper-dev/brc-rpc/resp"
)

func InscriptionTxList(insNumber int64, page, size int) (ret []*entity.InscriptionHistoryResponse, count int64, code int) {
	insTx := dao.InscriptionTxsInfo{InsNumber: insNumber}
	inscriptionTxs, count, err := insTx.Find(nil, dao.Paginate(page, size))
	if err != nil {
		log.Logger().WithField("insNumber", insNumber).
			WithField("page", page).
			WithField("size", size).
			WithField("error", err).
			Error("failed to get inscriptionTx list")
		return nil, 0, resp.CodeInternalServerError
	}

	ret = make([]*entity.InscriptionHistoryResponse, 0)
	for _, tx := range inscriptionTxs {
		ret = append(ret, &entity.InscriptionHistoryResponse{
			TxID:   tx.TxHash,
			Block:  strconv.FormatInt(tx.BlockHeight, 10),
			TxType: convertEvent(tx.Event),
			From:   tx.From,
			To:     tx.To,
			Time:   "", // todo 没有这个数据源
		})
	}
	return ret, count, resp.CodeSuccess
}

func convertEvent(t int64) string {
	switch t {
	case 1:
		return "inscribe"
	case 2:
		return "transfer"
	default:
		return ""
	}
}
