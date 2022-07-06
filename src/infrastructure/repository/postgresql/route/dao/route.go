package dao

type Route struct {
	Id           int64  `gorm:"primaryKey"`
	RelativePath string `gorm:"unique"`
	UrlTarget    string
	TypeTarget   string
	Enable       bool
}
