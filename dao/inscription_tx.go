package dao

import (
	"github.com/roullper-dev/brc-rpc/resource/db"
	"github.com/roullper-dev/brc-rpc/utils"
)

type InscriptionTxsInfo struct {
	InsNumber   int64  `gorm:"ins_number"`
	BlockHeight int64  `gorm:"block_height"`
	TxIndex     int64  `gorm:"tx_index"`
	TxHash      string `gorm:"tx_hash"`
	OutputHash  string `gorm:"output_hash"`
	From        string `gorm:"from"`
	To          string `gorm:"to"`
	Event       int64  `gorm:"event"` // create 1 、 transfer 2
}

func (it *InscriptionTxsInfo) TableName() string {
	return "inscription_txs"
}

func (it *InscriptionTxsInfo) Find(ext *QueryExtra, pager Pager) (is []*InscriptionTxsInfo, count int64, err error) {
	tx := db.GetOrdDB().Where(it)
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
			return is, count, err
		}
		if count == 0 {
			return is, 0, nil
		}
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&is).Error
	return is, count, err
}
