package logic

import (
	"strconv"

	"github.com/rolluper-dev/brc-rpc/constant"
	"github.com/rolluper-dev/brc-rpc/dao"
	"github.com/rolluper-dev/brc-rpc/entity"
	"github.com/rolluper-dev/brc-rpc/resource/log"
	"github.com/rolluper-dev/brc-rpc/resp"
)

func InscriptionList(address string, page, size int) (ret []*entity.InscriptionListResponse, count int64, code int) {
	ins := dao.Inscription{Address: address}
	inscriptions, count, err := ins.Find(nil, dao.Paginate(page, size))
	if err != nil {
		log.Logger().WithField("address", address).
			WithField("page", page).
			WithField("size", size).
			WithField("error", err).
			Error("failed to get inscription list")
		return nil, 0, resp.CodeInternalServerError
	}

	ret = make([]*entity.InscriptionListResponse, 0)
	for _, i := range inscriptions {
		ret = append(ret, &entity.InscriptionListResponse{
			Owner:             i.Address,
			Output:            i.Output,
			InscriptionID:     i.InsId,
			InscriptionNumber: i.InsNumber,
		})
	}
	return ret, count, resp.CodeSuccess
}

func Inscription(id, number string) (ret *entity.InscriptionResponse, code int) {
	var insNumber int64
	if number != "" {
		insNumber, _ = strconv.ParseInt(number, 10, 64)
	}
	where := map[string]interface{}{}
	if id != "" {
		where["ins_id"] = id
	}
	if insNumber != 0 {
		where["ins_number"] = insNumber
	}

	ins := dao.Inscription{}
	err := ins.First(where)
	if err != nil {
		log.Logger().WithField("id", id).
			WithField("number", number).
			WithField("error", err).
			Error("failed to get inscription")
		return nil, resp.CodeInternalServerError
	}
	ret = &entity.InscriptionResponse{
		InscriptionID:     ins.InsId,
		InscriptionNumber: ins.InsNumber,
		TxID:              ins.GenesisTransaction,
		Holder:            ins.GenesisAddress,
		Creator:           ins.Address,
		DeployHeight:      strconv.FormatInt(ins.GenesisHeight, 10),
		DeployTime:        ins.Timestamp.Format(constant.TimeFormat),
		Output:            ins.Output,
	}

	return ret, resp.CodeSuccess
}
