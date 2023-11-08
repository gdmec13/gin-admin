package model

import (
	"errors"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt int  `gorm:"type:int(11); comment:创建时间"`
	UpdatedAt int  `gorm:"type:int(11); comment:更新时间"`
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("id") {
		return errors.New("ID 作为唯一值不能更新")
	}
	return nil
}
