package logic

import (
	"strconv"

	"github.com/rolluper-dev/brc-rpc/dao"
	"github.com/rolluper-dev/brc-rpc/entity"
	"github.com/rolluper-dev/brc-rpc/resource/log"
	"github.com/rolluper-dev/brc-rpc/resp"
)

func AddressTransferList(address, tick, brc201Chain string, page, size int) (ret []*entity.AddressTransferListResponse, count int64, code int) {
	ub := dao.BRC20InscribeTransfer{
		UserBalanceAddr: address,
		TokenTick:       tick,
		BRC201Chain:     brc201Chain,
		IsUsed:          false,
	}
	txs, count, err := ub.Find(nil, dao.Paginate(page, size))
	if err != nil {
		log.Logger().WithField("token", tick).
			WithField("page", page).
			WithField("size", size).
			WithField("error", err).
			Error("addressTokenBalance failed to get userBalance list")
		return nil, 0, resp.CodeInternalServerError
	}

	for _, t := range txs {
		ret = append(ret, &entity.AddressTransferListResponse{
			TokenType:         "brc20",
			TokenTick:         t.TokenTick,
			Amount:            strconv.FormatUint(t.Amount, 10),
			InscriptionId:     t.InscriptionID,
			InscriptionNumber: strconv.FormatUint(t.InscriptionNumber, 10),
			Brc201Extend:      int(t.BRC201Extend),
			Brc201Chain:       t.BRC201Chain,
			Brc201Ref:         t.BRC201Reference,
		})
	}

	return ret, count, resp.CodeSuccess
}
