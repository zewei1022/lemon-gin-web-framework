package model

type Model struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"` // 主键ID
	CreatedAt int  `json:"created_at"`           // 创建时间
	UpdatedAt int  `json:"updated_at"`           // 更新时间
}
