package dao

import (
	"time"

	"github.com/rolluper-dev/brc-rpc/resource/db"
	"github.com/rolluper-dev/brc-rpc/utils"
)

type BRC20InscribeTransfer struct {
	ID                uint64    `gorm:"primarykey"`
	InscriptionTxsID  uint64    `gorm:"column:inscription_txs_id" json:"inscription_txs_id" sql:"bigint(20)"`
	TokenID           uint64    `gorm:"column:token_id" json:"token_id" sql:"bigint(20)"`
	TokenTick         string    `gorm:"column:token_tick" json:"token_tick" sql:"char(4)"`
	UserBalanceID     uint64    `gorm:"column:user_balance_id" json:"user_balance_id" sql:"bigint(20)"`
	UserBalanceAddr   string    `gorm:"column:user_balance_addr" json:"user_balance_addr" sql:"varchar(256)"`
	InscriptionID     string    `gorm:"column:inscription_id" json:"inscription_id" sql:"varchar(255)"`
	InscriptionNumber uint64    `gorm:"column:inscription_number" json:"inscription_number" sql:"bigint(20)"`
	To                string    `gorm:"column:to" json:"to" sql:"varchar(255)"`
	Amount            uint64    `gorm:"column:amount" json:"amount" sql:"bigint(20)"`
	BRC201Extend      uint8     `gorm:"column:brc201_extend" json:"brc201_extend" sql:"tinyint(4)"`
	BRC201Chain       string    `gorm:"column:brc201_chain" json:"brc201_chain" sql:"varchar(128)"`
	BRC201Reference   string    `gorm:"column:brc201_reference" json:"brc201_reference" sql:"varchar(255)"`
	IsValid           bool      `gorm:"column:is_valid" json:"is_valid" sql:"tinyint(1)"`
	IsUsed            bool      `gorm:"column:is_used" json:"is_used" sql:"tinyint(1)"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
}

func (it *BRC20InscribeTransfer) TableName() string {
	return "brc20_inscribe_transfer"
}

func (it *BRC20InscribeTransfer) Find(ext *QueryExtra, pager Pager) (ts []*BRC20InscribeTransfer, count int64, err error) {
	tx := db.GetDB().Where(it)
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
		if err := tx.Model(it).Count(&count).Error; err != nil {
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
