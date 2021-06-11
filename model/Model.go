package model

type Model struct {
	ID        uint `gorm:"primarykey;AUTO_INCREMENT" json:"id"` // 主键ID
	CreatedAt int  `json:"created_at"`           // 创建时间
	UpdatedAt int  `json:"updated_at"`           // 更新时间
	DeletedAt int  `gorm:"index" json:"-"`       // 删除时间
}
