package model

// 用户等级定义
type Grade struct {
	GradeId   uint32  `gorm:"column:grade_id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'等级ID'" json:"grade_id"`
	Name      string  `gorm:"column:name;default:;NOT NULL;comment:'等级名称'" json:"name"`
	MinBuy    float64 `gorm:"column:min_buy;default:0;NOT NULL;comment:'累计采购额最小值'" json:"min_buy"`
	MaxBuy    float64 `gorm:"column:max_buy;default:0;NOT NULL;comment:'累计采购额最小值'" json:"max_buy"`
	IsDefault int8    `gorm:"column:is_default;default:0;comment:'是否默认，1是，0否'" json:"is_default"`
	Remark    string  `gorm:"column:remark;default:;comment:'备注'" json:"remark"`
	IsDelete  uint8   `gorm:"column:is_delete;default:0;NOT NULL;comment:'是否删除'" json:"is_delete"`
	//是否享有推广收益
	EnjoyInviteProfit int8 `json:"enjoy_invite_profit"`
	//是否自动升级
	AutoUp int8 `json:"auto_up"`
}

func (u *Grade) TableName() string {
	return "grade"
}
