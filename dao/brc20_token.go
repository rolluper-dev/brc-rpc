package dao

import (
	"github.com/roullper-dev/brc-rpc/resource/db"
	"github.com/roullper-dev/brc-rpc/utils"
	"time"
)

type BRC20Token struct {
	ID                     uint64    `gorm:"primarykey"`
	Tick                   string    `gorm:"column:tick" json:"tick" sql:"char(4)"`
	Max                    uint64    `gorm:"column:max" json:"max" sql:"bigint(20)"`
	Limit                  uint64    `gorm:"column:limit" json:"limit" sql:"bigint(20)"`
	Decimal                uint8     `gorm:"column:decimal" json:"decimal" sql:"tinyint(6)"`
	MintedTimes            uint64    `gorm:"column:minted_times" json:"minted_times" sql:"bigint(20)"`
	TotalMinted            uint64    `gorm:"column:total_minted" json:"total_minted" sql:"bigint(20) default=0"`
	Deployer               string    `gorm:"column:deployer" json:"deployer" sql:"char(4)"`
	DeployTime             int64     `gorm:"column:deploy_time" json:"deploy_time" sql:"bigint(20)"`
	DeployHeight           int64     `gorm:"column:deploy_height" json:"deploy_height" sql:"bigint(20)"`
	IsCompleted            bool      `gorm:"column:is_completed" json:"is_completed" sql:"tinyint(4)"`
	CompletedTime          int64     `gorm:"column:completed_time" json:"completed_time" sql:"bigint(20)"`
	CompletedHeight        int64     `gorm:"column:completed_height" json:"completed_height" sql:"bigint(20)"`
	TxID                   string    `gorm:"column:tx_id" json:"tx_id" sql:"varchar(255)"`
	InscriptionID          string    `gorm:"column:inscription_id" json:"inscription_id" sql:"char(66)"`
	InscriptionNumberStart uint64    `gorm:"column:inscription_number_start" json:"inscription_number_start" sql:"bigint(20)"`
	InscriptionNumberEnd   uint64    `gorm:"column:inscription_number_end" json:"inscription_number_end" sql:"bigint(20)"`
	CreatedAt              time.Time `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt              time.Time `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
}

func (t *BRC20Token) TableName() string {
	return "brc20_token"
}

func NewBRC20Token(tick string) *BRC20Token {
	return &BRC20Token{
		Tick: tick,
	}
}

func (t *BRC20Token) Get() (token *BRC20Token, err error) {
	err = db.GetDB().Model(t).Where(t).First(&token).Error
	return token, err
}

func (t *BRC20Token) Find(ext *QueryExtra, pager Pager) (ts []*BRC20Token, count int64, err error) {
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
