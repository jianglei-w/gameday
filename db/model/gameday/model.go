package gameday

import "time"

type GameModel struct {
	ID       uint      `gorm:primarykey` // 主键
	CreateAt time.Time // 创建时间
	UpdateAt time.Time // 更新时间
	DeleteAt time.Time `gorm:"index" json:"-"` // 删除时间
}
