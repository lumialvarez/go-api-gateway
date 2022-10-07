package dao

type Route struct {
	Id           int64  `gorm:"primaryKey,autoIncrement"`
	RelativePath string `gorm:"unique"`
	UrlTarget    string
	TypeTarget   string
	Secure       bool
	Enable       bool
}
