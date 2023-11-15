package dao

import (
	"fmt"
	"time"

	"github.com/rolluper-dev/brc-rpc/utils"

	"github.com/rolluper-dev/brc-rpc/resource/db"
)

type BRC20UserBalance struct {
	ID                  uint64    `gorm:"primarykey"`
	Address             string    `gorm:"column:address" json:"address" sql:"varchar(255)"`
	Tick                string    `gorm:"column:tick" json:"tick" sql:"char(4)"`
	TokenID             uint64    `gorm:"column:token_id" json:"token_id" sql:"bigint(20)"`
	Balance             uint64    `gorm:"column:balance" json:"balance" sql:"bigint(20)"`
	BalanceSafe         uint64    `gorm:"column:balance_safe" json:"balance_safe" sql:"bigint(20)"`
	TransferableBalance uint64    `gorm:"column:transferable_balance" json:"transferable_balance" sql:"bigint(20)"`
	CreatedAt           time.Time `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
}

func (ub *BRC20UserBalance) TableName() string {
	return "brc20_user_balance"
}

func (ub *BRC20UserBalance) Count() (n int64, err error) {
	err = db.GetDB().Model(ub).Where(ub).Count(&n).Error
	return n, err
}

// BatchCount
// SELECT id, sum(tick=ordi) AS token_ordi, sum(tick=abcd) as token_abcd, sum(tick=meme) as token_meme from brc20_user_balance;
func (ub *BRC20UserBalance) BatchCount(ticks []string) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	s := selectQuery(ticks)
	err := db.GetDB().Model(ub).Select(s).Scan(&m).Error
	return m, err
}

func (ub *BRC20UserBalance) Find(ext *QueryExtra, pager Pager) (ts []*BRC20UserBalance, count int64, err error) {
	tx := db.GetDB().Where(ub)
	if ext != nil {
		if ext.Conditions != nil {
			for k, v := range ext.Conditions {
				tx = tx.Where(k, v)
			}
		}
		if !utils.IsEmpty(ext.OrderStr) {
			tx = tx.Order(ext.OrderStr)
		}
	}

	if pager != nil {
		if err := tx.Model(ub).Count(&count).Error; err != nil {
			return ts, count, err
		}
		if count == 0 {
			return ts, 0, nil
		}
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&ts).Error
	return ts, count, err
}

func selectQuery(ticks []string) string {
	selectStr := ""
	length := len(ticks)
	for i, t := range ticks {
		if i == length-1 {
			selectStr += fmt.Sprintf("sum(tick='%s') as 'token_%s'", t, t)
			return selectStr
		}
		selectStr += fmt.Sprintf("sum(tick='%s') as 'token_%s', ", t, t)
	}
	return selectStr
}
