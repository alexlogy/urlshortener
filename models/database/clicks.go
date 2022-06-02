package models

type Clicks struct {
	ID           uint   `gorm:"column:id;primaryKey"`
	IPAddr       string `gorm:"column:ip_address"`
	Country      string `gorm:"column:country"`
	Referrer     string `gorm:"column:referrer"`
	ReferrerHost string `gorm:"column:referrer_host"`
	UserAgent    string `gorm:"column:user_agent"`
	LinksID      uint   `gorm:"column:link_id"`
	Updated      int64  `gorm:"autoUpdateTime"`
	Created      int64  `gorm:"autoCreateTime"`
}
