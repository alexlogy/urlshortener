package models

type Users struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	UserName string `gorm:"column:name"`
	Password string `gorm:"column:email"`
	Email    string `gorm:"column:email"`
	Updated  int64  `gorm:"autoUpdateTime"`
	Created  int64  `gorm:"autoCreateTime"`
}
