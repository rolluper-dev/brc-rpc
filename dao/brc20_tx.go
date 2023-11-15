package dao

import (
	"time"

	"github.com/rolluper-dev/brc-rpc/resource/db"
	"github.com/rolluper-dev/brc-rpc/utils"
)

const (
	TypeCreate   = iota + 1
	TypeTransfer = iota + 1
)

type BRC20Tx struct {
	ID                uint64    `gorm:"primarykey"`
	TokenID           uint64    `gorm:"column:token_id" json:"token_id" sql:"bigint(20)"`
	TokenTick         string    `gorm:"column:token_tick" json:"token_tick" sql:"char(4)"`
	UserBalanceID     uint64    `gorm:"column:user_balance_id" json:"user_balance_id" sql:"bigint(20)"`
	UserBalanceAddr   string    `gorm:"column:user_balance_addr" json:"user_balance_addr" sql:"varchar(256)"`
	InscriptionID     string    `gorm:"column:inscription_id" json:"inscription_id" sql:"varchar()"`
	InscriptionNumber uint64    `gorm:"column:inscription_number" json:"inscription_number" sql:"bigint(20)"`
	BlockHeight       uint64    `gorm:"column:block_height" json:"block_height" sql:"bigint(20)"`
	TxID              string    `gorm:"column:tx_id" json:"tx_id" sql:"varchar(255)"`
	TxIndex           int64     `gorm:"tx_index" json:"tx_index" sql:"int(11)"`
	From              string    `gorm:"column:from" json:"from" sql:"varchar(256)"`
	To                string    `gorm:"column:to" json:"to" sql:"varchar(256)"`
	Amount            uint64    `gorm:"column:amount" json:"amount" sql:"bigint(20)"`
	Type              uint8     `gorm:"column:type" json:"type" sql:"tinyint(4)"`
	HistoryType       uint8     `gorm:"column:history_type" json:"history_type" sql:"tinyint(4)"`
	IsValid           bool      `gorm:"column:is_valid" json:"is_valid" sql:"tinyint(1)"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	BRC201Extend      uint8     `gorm:"column:brc201_extend" json:"brc201_extend" sql:"tinyint(4)"`
	BRC201Chain       string    `gorm:"column:brc201_chain" json:"brc201_chain" sql:"varchar(128)"`
	BRC201Reference   string    `gorm:"column:brc201_reference" json:"brc201_reference" sql:"varchar(255)"`
}

func (t *BRC20Tx) IsTransfer() bool {
	return t.Type == TypeTransfer
}

func (t *BRC20Tx) TableName() string {
	return "brc20_tx"
}

func (t *BRC20Tx) Find(ext *QueryExtra, pager Pager) (ts []*BRC20Tx, count int64, err error) {
	tx := db.GetDB().Where(t)
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
		if err := tx.Model(t).Count(&count).Error; err != nil {
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
