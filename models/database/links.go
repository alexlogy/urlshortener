package models

type Links struct {
	ID         uint   `gorm:"column:id;primaryKey"`
	ShortUrl   string `gorm:"column:short_url"`
	LongUrl    string `gorm:"column:long_url"`
	UserID     uint   `gorm:"column:user_id"`
	Clicks     int64  `gorm:"column:clicks;default:0"`
	IsDisabled bool   `gorm:"column:is_disabled;default:0"`
	Updated    int64  `gorm:"autoUpdateTime"`
	Created    int64  `gorm:"autoCreateTime"`
}
