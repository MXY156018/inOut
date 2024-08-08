package model

const (
	// 默认类型(字符串)
	Parameter_Type_Default int8 = iota
	// 图片类型(选择图片类配置)
	Parameter_Type_ImageUrl
)

// 可选值说明(主要方便前端编辑)
type ParameteLimit struct {
	//标签(展示)
	Label string `json:"l"`
	//实际值
	Value string `json:"v"`
}

// 通用参数
type Parameter struct {
	// 参数
	Parameter string `gorm:"primaryKey" json:"parameter"`
	// 值
	Value string `json:"value"`
	// 说明
	Remark string `json:"remark,optional"`
	// 类型 Parameter_Type_
	Type int8 `json:"t,optional"`
	// 可选值(JSON格式化后的 []ParameteLimit )
	Limit string `json:"limit,omitempty,optional"`
}

func (u *Parameter) TableName() string {
	return "parameter"
}
