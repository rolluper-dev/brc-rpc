package dao

import (
	"time"

	"github.com/rolluper-dev/brc-rpc/resource/db"
	"github.com/rolluper-dev/brc-rpc/utils"
)

type Inscription struct {
	Id                 int64     `gorm:"id"`
	InsNumber          int64     `gorm:"ins_number"`
	InsId              string    `gorm:"ins_id"`
	Address            string    `gorm:"address"`
	OutPutValue        string    `gorm:"out_put_value"`
	Sat                int64     `gorm:"sat"`
	Preview            string    `gorm:"preview"`
	Content            string    `gorm:"content"`
	ContentLength      string    `gorm:"content_length"`
	ContentType        string    `gorm:"content_type"`
	Timestamp          time.Time `gorm:"timestamp"`
	GenesisHeight      int64     `gorm:"genesis_height"`
	GenesisFee         string    `gorm:"genesis_fee"`
	GenesisTransaction string    `gorm:"genesis_transaction"`
	Location           string    `gorm:"location"`
	Output             string    `gorm:"output"`
	Offset             int64     `gorm:"offset"`
	GenesisAddress     string    `gorm:"genesis_address"`
	FinalAddress       string    `gorm:"final_address"`
}

func (i *Inscription) TableName() string {
	return "inscription"
}

func (i *Inscription) Find(ext *QueryExtra, pager Pager) (is []*Inscription, count int64, err error) {
	tx := db.GetOrdDB().Where(i)
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
		if err := tx.Model(i).Count(&count).Error; err != nil {
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

func (i *Inscription) First(where map[string]interface{}) error {
	return db.GetOrdDB().Where(where).First(i).Error
}
