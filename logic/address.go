package logic

import (
	"fmt"
	"github.com/roullper-dev/brc-rpc/constant"
	"github.com/roullper-dev/brc-rpc/dao"
	"github.com/roullper-dev/brc-rpc/entity"
	"github.com/roullper-dev/brc-rpc/resource/log"
	"github.com/roullper-dev/brc-rpc/resp"
	"strconv"
)

func AddressTokenBalance(address, tick string, page, size int) (ret []*entity.AddressTokenBalanceResponse, count int64, code int) {
	ub := dao.BRC20UserBalance{
		Address: address,
		Tick:    tick,
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
		ret = append(ret, &entity.AddressTokenBalanceResponse{
			TokenType:           "brc20",
			TokenTick:           t.Tick,
			Balance:             strconv.FormatUint(t.Balance, 10),
			TransferableBalance: strconv.FormatUint(t.TransferableBalance, 10),
			AvaliableBalance:    strconv.FormatUint(t.BalanceSafe, 10),
		})
	}

	return ret, count, resp.CodeSuccess
}

func AddressHistory(address, token, tp, brc201Chain, brc201Extend string, page, size int) (ret []*entity.AddressHistoryResponse, count int64, code int) {
	ub := dao.BRC20Tx{
		TokenTick:       token,
		HistoryType:     constant.HistoryType(tp),
		UserBalanceAddr: address,
		BRC201Chain:     brc201Chain,
		BRC201Extend:    constant.StrBrc201Extend(brc201Extend),
	}
	txs, count, err := ub.Find(nil, dao.Paginate(page, size))
	if err != nil {
		log.Logger().WithField("token", token).
			WithField("page", page).
			WithField("size", size).
			WithField("error", err).
			Error("addressTokenBalance failed to get userBalance list")
		return nil, 0, resp.CodeInternalServerError
	}

	fmt.Println("-------------- txs", len(txs), txs, "count", count)
	for _, t := range txs {
		ret = append(ret, &entity.AddressHistoryResponse{
			TxId:              t.TxID,
			Block:             strconv.FormatUint(t.BlockHeight, 10),
			TokenType:         "brc20",
			Token:             t.TokenTick,
			TxType:            constant.StrHistoryType(t.HistoryType),
			From:              t.From,
			To:                t.To,
			InscriptionNumber: strconv.FormatUint(t.InscriptionNumber, 10),
			InscriptionId:     t.InscriptionID,
			Time:              t.CreatedAt.Format(constant.TimeFormat),
			Amount:            strconv.FormatUint(t.Amount, 10),
			Brc201Extend:      int(t.BRC201Extend),
			Brc201Chain:       t.BRC201Chain,
		})
	}

	return ret, count, resp.CodeSuccess
}
