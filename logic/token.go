package logic

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/rolluper-dev/brc-rpc/constant"
	"github.com/rolluper-dev/brc-rpc/utils"
	"gorm.io/gorm"

	"github.com/rolluper-dev/brc-rpc/dao"
	"github.com/rolluper-dev/brc-rpc/entity"
	"github.com/rolluper-dev/brc-rpc/resource/log"
	"github.com/rolluper-dev/brc-rpc/resp"
)

func GetToken(tick string) (*entity.GetTokenResponse, int) {
	tk, err := dao.NewBRC20Token(tick).Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodeTokenNotExist
		}
		log.Logger().WithField("tick", tick).WithField("error", err).Error("failed to get token")
		return nil, resp.CodeInternalServerError
	}

	ub := dao.BRC20UserBalance{Tick: tick}
	count, err := ub.Count()
	if err != nil {
		log.Logger().WithField("tick", tick).WithField("error", err).Error("failed to count user balance")
		return nil, resp.CodeInternalServerError
	}
	return &entity.GetTokenResponse{
		TokenType:              constant.TokenTypeBRC20,
		Token:                  tk.Tick,
		TotalSupply:            tk.Max,
		Limit:                  tk.Limit,
		Decimals:               tk.Decimal,
		Minted:                 tk.TotalMinted,
		Creator:                tk.Deployer,
		DeployTime:             tk.DeployTime,
		DeployHeight:           tk.DeployHeight,
		Holders:                count,
		TxID:                   tk.TxID,
		InscriptionNumberStart: tk.InscriptionNumberStart,
		InscriptionNumberEnd:   tk.InscriptionNumberEnd,
	}, resp.CodeSuccess
}

func TokenList(status uint8, page, size int) (ret []*entity.TokenListResponse, count int64, code int) {
	ret = make([]*entity.TokenListResponse, 0)

	token := dao.BRC20Token{}
	if status == constant.TokenStatusMinting {
		token.IsCompleted = false
	} else if status == constant.TokenStatusComplete {
		token.IsCompleted = true
	}

	tokens, count, err := token.Find(nil, dao.Paginate(page, size))
	if err != nil {
		log.Logger().WithField("status", status).
			WithField("page", page).
			WithField("size", size).
			WithField("error", err).
			Error("failed to get token list")
		return nil, 0, resp.CodeInternalServerError
	}

	length := len(tokens)
	if length == 0 {
		return ret, count, resp.CodeSuccess
	}

	ticks := make([]string, 0, length)
	for _, t := range tokens {
		ticks = append(ticks, t.Tick)
	}

	ub := dao.BRC20UserBalance{}
	batchCount, err := ub.BatchCount(ticks)
	if err != nil {
		log.Logger().WithField("ticks", utils.JSON(ticks)).
			WithField("error", err).
			Error("failed to batch count user balance")
		return nil, 0, resp.CodeInternalServerError
	}

	ret = make([]*entity.TokenListResponse, 0)
	for _, t := range tokens {
		key := fmt.Sprintf("token_%s", t.Tick)
		if _, ok := batchCount[key]; !ok {
			log.Logger().Info("token is nil, please check", "key", key)
			continue
		}
		s := batchCount[key].(string)
		cnt, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Logger().WithField("count_string", s).
				WithField("error", err).
				Error("failed to parse string")
			return nil, 0, resp.CodeInternalServerError
		}
		ret = append(ret, &entity.TokenListResponse{
			TokenType:    constant.TokenTypeBRC20,
			Token:        t.Tick,
			TotalSupply:  t.Max,
			Limit:        t.Limit,
			Decimals:     t.Decimal,
			Minted:       t.TotalMinted,
			Creator:      t.Deployer,
			DeployTime:   t.DeployTime,
			DeployHeight: t.DeployHeight,
			Holders:      cnt,
		})
	}
	return ret, count, resp.CodeSuccess
}

func TokenHolder(tick string, page, size int) (ret []*entity.TokenHoldersResponse, count int64, code int) {
	ub := dao.BRC20UserBalance{Tick: tick}
	ubs, count, err := ub.Find(nil, dao.Paginate(page, size))
	if err != nil {
		log.Logger().WithField("token", tick).
			WithField("page", page).
			WithField("size", size).
			WithField("error", err).
			Error("failed to get userBalance list")
		return nil, 0, resp.CodeInternalServerError
	}

	for _, u := range ubs {
		ret = append(ret, &entity.TokenHoldersResponse{
			Address: u.Address,
			Amount:  strconv.FormatUint(u.TransferableBalance, 10),
		})
	}

	return ret, count, resp.CodeSuccess
}

func TokenHistory(tick, tp, from, to, brc201Chain, brc201Extend string, page, size int) (ret []*entity.TokenHistoryResponse, count int64, code int) {
	ub := dao.BRC20Tx{
		TokenTick:    tick,
		HistoryType:  constant.HistoryType(tp),
		From:         from,
		To:           to,
		BRC201Chain:  brc201Chain,
		BRC201Extend: constant.StrBrc201Extend(brc201Extend),
	}
	txs, count, err := ub.Find(nil, dao.Paginate(page, size))
	if err != nil {
		log.Logger().WithField("token", tick).
			WithField("page", page).
			WithField("size", size).
			WithField("error", err).
			Error("failed to get userBalance list")
		return nil, 0, resp.CodeInternalServerError
	}

	for _, t := range txs {
		ret = append(ret, &entity.TokenHistoryResponse{
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
