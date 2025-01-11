package constants

import "time"

type Ledger struct {
	ID          uint      `json:"id" gorm:"not null;primaryKey;autoIncrement"`
	ProjectType uint      `json:"projectType" gorm:"not null"`     // 项目类型，0：接单
	TradeType   uint      `json:"tradeType" gorm:"not null"`       // 交易类型，0：闲鱼，1：微信
	Price       float64   `json:"price" gorm:"not null"`           // 价格
	ExtraPrice  float64   `json:"extraPrice"`                      // 额外价格
	FinalPrice  float64   `json:"finalPrice"`                      // 最终价格
	Remark      string    `json:"remark" gorm:"type:varchar(255)"` // 备注信息
	TradeStart  time.Time `json:"tradeStart" gorm:"not null"`      // 交易开始时间
	TradeEnd    time.Time `json:"tradeEnd"`                        // 交易结束时间
	CreatedAt   time.Time `json:"createdAt" gorm:"AutoCreateTime"` // 创建时间
	UpdatedAt   time.Time `json:"updatedAt" gorm:"AutoUpdateTime"` // 更新时间
}

type LedgerListOpt struct {
	ProjectType uint
}

func (Ledger) TableName() string {
	return "ledger"
}
